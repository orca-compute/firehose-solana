package serumhist

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/dfuse-io/dfuse-solana/serumhist/keyer"

	bin "github.com/dfuse-io/binary"
	pbserumhist "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/serumhist/v1"
	"github.com/dfuse-io/solana-go/diff"
	"github.com/golang/protobuf/proto"

	pbcodec "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/codec/v1"
	kvdb "github.com/dfuse-io/kvdb/store"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/programs/serum"
	"go.uber.org/zap"
)

// TODO: this is very repetitive we need to optimze the account setting in solana go
// once that is done we can clean all this up!
func (i *Injector) processInstruction(ctx context.Context, slotNumber uint64, blkTime time.Time, inst *serumInstruction) error {
	var kvs []*kvdb.KV
	var err error

	trxID := inst.trxtID
	serumInstruction := inst.native
	// we only care about new order instruction that modify the request queue
	if newOrder, ok := serumInstruction.Impl.(*serum.InstructionNewOrder); ok {
		zlog.Debug("processing new order v1",
			zap.Uint64("slot_number", slotNumber),
			zap.String("trx_id", trxID),
		)
		trader := newOrder.Accounts.Owner.PublicKey
		tradingAccount := newOrder.Accounts.OpenOrders.PublicKey
		i.cache.setTradingAccount(ctx, tradingAccount, trader)
	} else if newOrderV2, ok := serumInstruction.Impl.(*serum.InstructionNewOrderV2); ok {
		zlog.Debug("processing new order v2",
			zap.Uint64("slot_number", slotNumber),
			zap.String("trx_id", trxID),
		)
		trader := newOrderV2.Accounts.Owner.PublicKey
		tradingAccount := newOrderV2.Accounts.OpenOrders.PublicKey
		i.cache.setTradingAccount(ctx, tradingAccount, trader)
	} else if mathOrder, ok := serumInstruction.Impl.(*serum.InstructionMatchOrder); ok {
		zlog.Debug("processing match order",
			zap.Uint64("slot_number", slotNumber),
			zap.String("trx_id", trxID),
		)

		if kvs, err = i.processMatchOrderInstruction(ctx, slotNumber, blkTime, inst.trxIdx, inst.instIdx, mathOrder, inst.accChanges); err != nil {
			return fmt.Errorf("generating serumhist keys: %w", err)
		}
	}

	if len(kvs) == 0 {
		return nil
	}

	zlog.Debug("inserting serumhist keys",
		zap.Int("key_count", len(kvs)),
	)

	for _, kv := range kvs {
		if err := i.kvdb.Put(ctx, kv.Key, kv.Value); err != nil {
			zlog.Warn("failed to write key-value", zap.Error(err))
		}
	}
	return nil
}

func (i *Injector) processMatchOrderInstruction(ctx context.Context, slotNumber uint64, blkTime time.Time, trxIdx, instIdx uint64, inst *serum.InstructionMatchOrder, accountChanges []*pbcodec.AccountChange) (out []*kvdb.KV, err error) {
	eventQueueAccountChange, err := filterAccountChange(accountChanges, func(flag *serum.AccountFlag) bool {
		return flag.Is(serum.AccountFlagInitialized) && flag.Is(serum.AccountFlagEventQueue)
	})

	if eventQueueAccountChange == nil {
		return nil, nil
	}

	old, new, err := decodeEventQueue(eventQueueAccountChange)
	if err != nil {
		return nil, fmt.Errorf("unable to decode event queue change: %w", err)
	}

	return i.getFillKeyValues(ctx, slotNumber, blkTime, trxIdx, instIdx, inst.Accounts.Market.PublicKey, old, new)
}

func (i *Injector) getFillKeyValues(ctx context.Context, slotNumber uint64, blkTime time.Time, trxIdx, instIdx uint64, market solana.PublicKey, old, new *serum.EventQueue) (out []*kvdb.KV, err error) {
	diff.Diff(old, new, diff.OnEvent(func(event diff.Event) {
		if match, _ := event.Match("Events[#]"); match {
			e := event.Element().Interface().(*serum.Event)
			switch event.Kind {
			case diff.KindAdded:
				zlog.Debug("found a serum event diff",
					zap.Stringer("diff_kind", event.Kind),
					zap.Stringer("event_flag", e.Flag),
					zap.Bool("is_fill", e.Flag.IsFill()),
				)

				if e.Flag.IsFill() {

					tradingAccount := e.Owner
					trader, err := i.cache.getTrader(ctx, tradingAccount)
					if err != nil {
						zlog.Error("error retrieving trader key",
							zap.Error(err),
							zap.Stringer("trading_account", tradingAccount),
						)
						return
					}

					number := make([]byte, 16)

					binary.BigEndian.PutUint64(number[:], e.OrderID.Lo)
					binary.BigEndian.PutUint64(number[8:], e.OrderID.Hi)

					fill := &pbserumhist.Fill{
						OrderId:           fmt.Sprintf("%s%s", hex.EncodeToString(number[:8]), hex.EncodeToString(number[8:])),
						Side:              pbserumhist.Side(e.Side()),
						Maker:             false,
						NativeQtyPaid:     e.NativeQtyPaid,
						NativeQtyReceived: e.NativeQtyReleased,
						NativeFeeOrRebate: e.NativeFeeOrRebate,
						FeeTier:           pbserumhist.FeeTier(e.FeeTier),
						Timestamp:         mustProtoTimestamp(blkTime),
					}

					cnt, err := proto.Marshal(fill)
					if err != nil {
						zlog.Error("unable to marshal to fill", zap.Error(err))
						return
					}

					orderSeqNum := extractOrderSeqNum(e.Side(), e.OrderID)

					zlog.Info("serum new fill",
						zap.Uint32("side", uint32(e.Side())),
						zap.Stringer("market", market),
						zap.Stringer("trader", trader),
						zap.Stringer("trading_Account", tradingAccount),
						zap.Uint64("order_seq_num", orderSeqNum),
						zap.Uint64("slot_num", slotNumber),
					)

					out = append(out,
						&kvdb.KV{
							Key:   keyer.EncodeFillByTrader(*trader, market, slotNumber, trxIdx, instIdx, orderSeqNum),
							Value: cnt,
						},
						&kvdb.KV{
							Key:   keyer.EncodeFillByMarketTrader(*trader, market, slotNumber, trxIdx, instIdx, orderSeqNum),
							Value: cnt,
						},
					)
				}
			default:
			}
		}
	}))
	return out, nil
}

func extractOrderSeqNum(side serum.Side, orderID bin.Uint128) uint64 {
	if side == serum.SideBid {
		return ^orderID.Lo
	}
	return orderID.Lo
}

func decodeEventQueue(accountChange *pbcodec.AccountChange) (old *serum.EventQueue, new *serum.EventQueue, err error) {
	if err := bin.NewDecoder(accountChange.PrevData).Decode(&old); err != nil {
		return nil, nil, fmt.Errorf("unable to decode 'event queue' old data: %w", err)
	}

	if err := bin.NewDecoder(accountChange.NewData).Decode(&new); err != nil {
		return nil, nil, fmt.Errorf("unable to decode 'event queue' new data: %w", err)
	}

	return
}

func mustProtoTimestamp(in time.Time) *timestamp.Timestamp {
	out, err := ptypes.TimestampProto(in)
	if err != nil {
		panic(fmt.Sprintf("invalid timestamp conversion %q: %s", in, err))
	}
	return out
}
