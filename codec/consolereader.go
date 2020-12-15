// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"

	pbcodec "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/codec/v1"
	"go.uber.org/zap"
)

var supportedVersions = []string{"12", "13"}

type conversionOption interface{}
type ConsoleReaderOption interface {
	apply(reader *ConsoleReader)
}

// ConsoleReader is what reads the `nodeos` output directly. It builds
// up some LogEntry objects. See `LogReader to read those entries .
type ConsoleReader struct {
	src        io.Reader
	scanner    *bufio.Scanner
	close      func()
	readBuffer chan string
	done       chan interface{}

	ctx *parseCtx
}

func NewConsoleReader(reader io.Reader, opts ...ConsoleReaderOption) (*ConsoleReader, error) {
	l := &ConsoleReader{
		src:   reader,
		close: func() {},
		ctx:   newParseCtx(),
		done:  make(chan interface{}),
	}

	for _, opt := range opts {
		opt.apply(l)
	}

	l.setupScanner()
	return l, nil
}

func (l *ConsoleReader) setupScanner() {
	maxTokenSize := uint64(50 * 1024 * 1024)
	if maxBufferSize := os.Getenv("MINDREADER_MAX_TOKEN_SIZE"); maxBufferSize != "" {
		bs, err := strconv.ParseUint(maxBufferSize, 10, 64)
		if err != nil {
			zlog.Error("environment variable 'MINDREADER_MAX_TOKEN_SIZE' is set but invalid parse uint", zap.Error(err))
		} else {
			zlog.Info("setting max_token_size from environment variable MINDREADER_MAX_TOKEN_SIZE", zap.Uint64("max_token_size", bs))
			maxTokenSize = bs
		}
	}
	buf := make([]byte, maxTokenSize)
	scanner := bufio.NewScanner(l.src)
	scanner.Buffer(buf, len(buf))
	l.scanner = scanner
	l.readBuffer = make(chan string, 2000)

	go func() {
		for l.scanner.Scan() {
			line := l.scanner.Text()
			if !strings.HasPrefix(line, "DMLOG ") {
				continue
			}
			l.readBuffer <- line
		}

		err := l.scanner.Err()
		if err != nil && err != io.EOF {
			zlog.Error("console read line scanner encountered an error", zap.Error(err))
		}

		close(l.readBuffer)
		close(l.done)
	}()
}

func (l *ConsoleReader) Done() <-chan interface{} {
	return l.done
}

func (l *ConsoleReader) Close() {
	l.close()
}

type parseCtx struct {
	slot          *pbcodec.Slot
	activeSlotNum uint64
	trxIndex      uint64
	trxTraceMap   map[string]*pbcodec.TransactionTrace

	conversionOptions []conversionOption
}

func newParseCtx() *parseCtx {
	return &parseCtx{
		slot:        &pbcodec.Slot{},
		trxTraceMap: map[string]*pbcodec.TransactionTrace{},
	}
}

func (l *ConsoleReader) Read() (out interface{}, err error) {
	ctx := l.ctx

	for line := range l.readBuffer {
		line = line[6:]

		if traceEnabled {
			zlog.Debug("extracing deep mind data from line", zap.String("line", line))
		}

		// Order of conditions is based (approximately) on those that will appear more often
		switch {
		case strings.HasPrefix(line, "TRANSACTION START"):
			err = ctx.readTransactionStart(line)

		case strings.HasPrefix(line, "INSTRUCTION START"):
			err = ctx.readInstructionTraceStart(line)

		case strings.HasPrefix(line, "ACCOUNT_CHANGE"):
			err = ctx.readAccountChange(line)

		case strings.HasPrefix(line, "LAMPORT_CHANGE"):
			err = ctx.readLamportsChange(line)

		default:
			zlog.Info("unknown log line", zap.String("line", line))
		}

		if err != nil {
			return nil, l.formatError(line, err)
		}
	}

	if l.scanner.Err() == nil {
		return nil, io.EOF
	}

	return nil, l.scanner.Err()
}

func (l *ConsoleReader) formatError(line string, err error) error {
	chunks := strings.SplitN(line, " ", 2)
	return fmt.Errorf("%s: %s (line %q)", chunks[0], err, line)
}

type creationOp struct {
	kind        string // ROOT, NOTIFY, CFA_INLINE, INLINE
	actionIndex int
}

func (ctx *parseCtx) resetBlock() {
	if ctx.activeSlotNum != 0 {
		ctx.resetTrx()
	}

	ctx.slot = &pbcodec.Slot{}
}

func (ctx *parseCtx) resetTrx() {
	ctx.trxTraceMap = map[string]*pbcodec.TransactionTrace{}

}

func (ctx *parseCtx) readSlotStart(line string) error {
	ctx.resetTrx()
	ctx.activeSlotNum = 0 //todo: get slot from line ...
	return nil
}

func (ctx *parseCtx) readTransactionStart(line string) error {
	chunks := strings.SplitN(line, " ", -1)
	if len(chunks) != 4 {
		return fmt.Errorf("read transaction start: expected 4 fields, got %d", len(chunks))
	}

	ctx.resetTrx()

	id := chunks[2]
	signatures := []string{id}

	var solMessage *solana.Message
	messageData, err := hex.DecodeString(chunks[3])
	if err != nil {
		return fmt.Errorf("read transaction start: hex decode message: %w", err)
	}

	err = bin.NewDecoder(messageData).Decode(&solMessage)
	if err != nil {
		return fmt.Errorf("read transaction start: binary decode message: %w", err)
	}

	var accountKeys [][]byte
	for _, k := range solMessage.AccountKeys {
		accountKeys = append(accountKeys, k[:])
	}

	var instructions []*pbcodec.CompiledInstruction
	for _, i := range solMessage.Instructions {

		var accountIdIndexes []uint32
		for _, i := range i.Accounts {
			accountIdIndexes = append(accountIdIndexes, uint32(i))
		}

		instruction := &pbcodec.CompiledInstruction{
			ProgramIdIndex: uint32(i.ProgramIDIndex),
			Accounts:       accountIdIndexes,
			Data:           i.Data,
		}
		instructions = append(instructions, instruction)
	}

	message := &pbcodec.Message{
		Header: &pbcodec.MessageHeader{
			NumRequiredSignatures:       uint32(solMessage.Header.NumRequiredSignatures),
			NumReadonlySignedAccounts:   uint32(solMessage.Header.NumReadonlySignedAccounts),
			NumReadonlyUnsignedAccounts: uint32(solMessage.Header.NumReadonlyUnsignedAccounts),
		},
		AccountKeys:     accountKeys,
		RecentBlockhash: solMessage.RecentBlockhash[:],
		Instructions:    instructions,
	}

	transaction := &pbcodec.Transaction{
		Id:         id,
		Index:      ctx.trxIndex,
		Signatures: signatures,
		Msg:        message,
	}
	ctx.recordTransaction(transaction)

	transactionTrace := &pbcodec.TransactionTrace{
		Id:       id,
		Index:    ctx.trxIndex,
		SlotNum:  uint64(ctx.slot.Number),
		SlotHash: ctx.slot.Id,
	}

	ctx.recordTransactionTrace(transactionTrace)
	return nil
}

func (ctx *parseCtx) recordTransaction(trx *pbcodec.Transaction) {
	ctx.slot.Transactions = append(ctx.slot.Transactions, trx)
}

func (ctx *parseCtx) recordTransactionTrace(trxTrace *pbcodec.TransactionTrace) {
	ctx.trxTraceMap[trxTrace.Id] = trxTrace
	ctx.trxIndex++

	return
}

func (ctx *parseCtx) readInstructionTraceStart(line string) error {
	chunks := strings.SplitN(line, " ", -1)
	if len(chunks) != 7 {
		return fmt.Errorf("read instructionTrace start: expected 7 fields, got %d", len(chunks))
	}
	trxID := chunks[2]
	ordinal, err := strconv.Atoi(chunks[3])
	if err != nil {
		return fmt.Errorf("read instructionTrace start: ordinal to int: %w", err)
	}

	parentOrdinal, err := strconv.Atoi(chunks[4])
	if err != nil {
		return fmt.Errorf("read instructionTrace start: parent ordinal to int: %w", err)
	}

	program := chunks[5]
	data := chunks[6]
	hexData, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("read instructionTrace start: hex decode data: %w", err)
	}

	instructionTrace := &pbcodec.InstructionTrace{
		ProgramId:     program,
		Data:          hexData,
		Ordinal:       uint32(ordinal),
		ParentOrdinal: uint32(parentOrdinal),
	}

	err = ctx.recordInstructionTrace(trxID, instructionTrace)
	if err != nil {
		return fmt.Errorf("read instructionTrace start: %w", err)
	}

	return nil
}

func (ctx *parseCtx) recordInstructionTrace(trxID string, instruction *pbcodec.InstructionTrace) error {
	trxTrace := ctx.trxTraceMap[trxID]
	if trxTrace == nil {
		return fmt.Errorf("record instruction: transaction trace not found in context: %s", trxID)
	}

	trxTrace.InstructionTraces = append(trxTrace.InstructionTraces, instruction)

	return nil
}

func (ctx *parseCtx) readAccountChange(line string) error {
	chunks := strings.SplitN(line, " ", -1)
	if len(chunks) != 7 {
		return fmt.Errorf("read account change: expected 7 fields, got %d", len(chunks))
	}
	trxID := chunks[1]
	ordinal, err := strconv.Atoi(chunks[2])
	if err != nil {
		return fmt.Errorf("read account change: ordinal to int: %w", err)
	}

	parentOrdinal, err := strconv.Atoi(chunks[3])
	if err != nil {
		return fmt.Errorf("read account change: parent ordinal to int: %w", err)
	}
	_ = parentOrdinal

	pubKey := chunks[4]

	prevData, err := hex.DecodeString(chunks[5])
	if err != nil {
		return fmt.Errorf("read account change: hex decode prev data: %w", err)
	}

	newData, err := hex.DecodeString(chunks[6])
	if err != nil {
		return fmt.Errorf("read account change: hex decode new data: %w", err)
	}

	accountChange := &pbcodec.AccountChange{
		Pubkey:   pubKey,
		PrevData: prevData,
		NewData:  newData,
	}

	err = ctx.recordAccountChange(trxID, ordinal, accountChange)
	if err != nil {
		return fmt.Errorf("read account change: %w", err)
	}

	return nil
}

func (ctx *parseCtx) recordAccountChange(trxID string, ordinal int, accountChange *pbcodec.AccountChange) error {
	trxTrace := ctx.trxTraceMap[trxID]
	if trxTrace == nil {
		return fmt.Errorf("record account change: transaction trace not found in context: %s", trxID)
	}

	trxTrace.InstructionTraces[ordinal-1].AccountChanges = append(trxTrace.InstructionTraces[ordinal-1].AccountChanges, accountChange)

	return nil
}

func (ctx *parseCtx) readLamportsChange(line string) error {
	chunks := strings.SplitN(line, " ", -1)
	if len(chunks) != 6 {
		return fmt.Errorf("read lamport change: expected 6 fields, got %d", len(chunks))
	}
	trxID := chunks[1]
	ordinal, err := strconv.Atoi(chunks[2])
	if err != nil {
		return fmt.Errorf("read lamport change: ordinal to int: %w", err)
	}

	owner := chunks[3]

	prevLamports, err := strconv.Atoi(chunks[4])
	if err != nil {
		return fmt.Errorf("read lamport change: hex decode prev lamports data: %w", err)
	}

	newLamports, err := strconv.Atoi(chunks[5])
	if err != nil {
		return fmt.Errorf("read lamport change: hex decode new lamports data: %w", err)
	}

	balanceChange := &pbcodec.BalanceChange{
		Pubkey:       owner,
		PrevLamports: uint64(prevLamports),
		NewLamports:  uint64(newLamports),
	}

	err = ctx.recordLamportsChange(trxID, ordinal, balanceChange)
	if err != nil {
		return fmt.Errorf("read lamports change: %w", err)
	}

	return nil
}

func (ctx *parseCtx) recordLamportsChange(trxID string, ordinal int, balanceChange *pbcodec.BalanceChange) error {
	trxTrace := ctx.trxTraceMap[trxID]
	if trxTrace == nil {
		return fmt.Errorf("record balanace change: transaction trace not found in context: %s", trxID)
	}

	trxTrace.InstructionTraces[ordinal-1].BalanceChanges = append(trxTrace.InstructionTraces[ordinal-1].BalanceChanges, balanceChange)

	return nil
}

func splitNToM(line string, min, max int) ([]string, error) {
	chunks := strings.SplitN(line, " ", -1)
	if len(chunks) < min || len(chunks) > max {
		return nil, fmt.Errorf("expected between %d to %d fields (inclusively), got %d", min, max, len(chunks))
	}

	return chunks, nil
}
