package resolvers

import (
	"github.com/dfuse-io/dauth/ratelimiter"
	pbserumhist "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/serumhist/v1"
	"github.com/dfuse-io/dfuse-solana/registry"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
)

// Root is the root resolvers of the schema
type Root struct {
	rpcClient          *rpc.Client
	wsURL              string
	registryServer     *registry.Server
	requestRateLimiter ratelimiter.RateLimiter
	serumHistoryClient pbserumhist.SerumHistoryClient

	// Interfaces we use internally for testing purposes
	marketGetter  func(address *solana.PublicKey) *registry.Market
	marketsGetter func() []*registry.Market
	tokenGetter   func(in *solana.PublicKey) *registry.Token
}

func NewRoot(
	rpcClient *rpc.Client,
	wsURL string,
	mdServer *registry.Server,
	requestRateLimiter ratelimiter.RateLimiter,
	serumHistoryClient pbserumhist.SerumHistoryClient,
) (*Root, error) {
	return &Root{
		rpcClient:          rpcClient,
		wsURL:              wsURL,
		registryServer:     mdServer,
		requestRateLimiter: requestRateLimiter,
		serumHistoryClient: serumHistoryClient,

		marketGetter:  mdServer.GetMarket,
		marketsGetter: mdServer.GetMarkets,
		tokenGetter:   mdServer.GetToken,
	}, nil
}
