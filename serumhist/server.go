package serumhist

import (
	"context"
	"fmt"

	"github.com/dfuse-io/solana-go"

	"github.com/dfuse-io/logging"

	pbserumhist "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/serumhist/v1"
	pbhealth "github.com/dfuse-io/pbgo/grpc/health/v1"
	"go.uber.org/zap"
)

func (s *Injector) TrackOrder(r *pbserumhist.TrackOrderRequest, stream pbserumhist.SerumOrderTracker_TrackOrderServer) error {
	ctx := stream.Context()
	logger := logging.Logger(ctx, zlog)
	logger.Debug("tracking order", zap.Reflect("request", r))

	market, err := solana.PublicKeyFromBase58(r.Market)
	if err != nil {
		return fmt.Errorf("unable to decode market key")
	}

	subscription, err := s.manager.subscribe(r.OrderId, market, logger)
	if err != nil {
		return fmt.Errorf("unable to create subscription: %w", err)
	}
	defer s.manager.unsubscribe(ctx, subscription)

	transition, err := s.reader.GetInitializeOrder(ctx, market, r.OrderId)
	if err != nil {
		return fmt.Errorf("unable to get initialized orders: %w", err)
	}

	if transition != nil {
		err := stream.Send(transition)
		if err != nil {
			logger.Info("failed writing to socket, shutting down subscription", zap.Error(err))
			return err
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case resp, opened := <-subscription.conn:
			if !opened {
				// we've been shutdown somehow, simply close the current connection.
				// we'll have logged at the source\
				return nil
			}
			logger.Debug("sending order transition",
				zap.Stringer("current_state", resp.CurrentState),
				zap.Stringer("previous_state", resp.PreviousState),
				zap.Stringer("transition", resp.Transition),
			)

			err := stream.Send(resp)
			if err != nil {
				logger.Info("failed writing to socket, shutting down subscription", zap.Error(err))
				return err
			}
		}
	}
}

func (s *Injector) Check(ctx context.Context, request *pbhealth.HealthCheckRequest) (*pbhealth.HealthCheckResponse, error) {
	panic("implement me")
}
