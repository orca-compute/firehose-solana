package resolvers

import (
	"context"
	"time"

	pbserumhist "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/serumhist/v1"
	"github.com/streamingfast/solana-go"
	gqerrs "github.com/graph-gophers/graphql-go/errors"
)

type SerumFillHistoryRequest struct {
	Trader *string
	Market *string
}

func (r *SerumFillHistoryRequest) validate() error {
	if r.Trader == nil && r.Market == nil {
		return gqerrs.Errorf("trader and market cannot both be empty")
	}
	return nil
}

func (r *Root) QuerySerumFillHistory(ctx context.Context, in *SerumFillHistoryRequest) (out *SerumFillConnection, err error) {
	if err := in.validate(); err != nil {
		return nil, err
	}

	request := &pbserumhist.GetFillsRequest{}

	if in.Trader != nil {
		traderKey, err := solana.PublicKeyFromBase58(*in.Trader)
		if err != nil {
			return nil, gqerrs.Errorf(`invalid "trader" argument %q: %s`, *in.Trader, err)
		}
		request.Trader = traderKey.String()
	}

	if in.Market != nil {
		marketKey, err := solana.PublicKeyFromBase58(*in.Market)
		if err != nil {
			return nil, gqerrs.Errorf(`invalid "market" argument %q: %s`, *in.Market, err)
		}
		request.Market = marketKey.String()
	}

	getCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	fillsResponse, err := r.serumHistoryClient.GetFills(getCtx, request)
	if err != nil {
		return nil, graphqlErrorFromGRPC(getCtx, err)
	}

	edges := make([]*SerumFillEdge, len(fillsResponse.Fill))
	for i, fill := range fillsResponse.Fill {
		edges[i] = &SerumFillEdge{cursor: "", node: r.newSerumFill(fill)}
	}

	return &SerumFillConnection{
		Edges:    edges,
		PageInfo: NewPageInfoFromEdges(edges, fillsResponse.HasMore),
	}, nil
}

type SerumFillEdge struct {
	cursor string
	node   SerumFill
	err    error
}

func NewSerumFillEdge(node SerumFill, cursor string) *SerumFillEdge {
	return &SerumFillEdge{
		cursor: cursor,
		node:   node,
	}
}

func (e *SerumFillEdge) Node() SerumFill          { return e.node }
func (e *SerumFillEdge) Cursor() string           { return e.cursor }
func (e *SerumFillEdge) SubscriptionError() error { return e.err }

type SerumFillConnection struct {
	Edges    []*SerumFillEdge
	PageInfo PageInfo
}

func NewSerumFillConnection(edges []*SerumFillEdge, pageInfo PageInfo) *SerumFillConnection {
	return &SerumFillConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}
}

var emptyPageInfo = PageInfo{}

func NewPageInfoFromEdges(edges []*SerumFillEdge, hasMore bool) PageInfo {
	if len(edges) == 0 {
		return emptyPageInfo
	}

	return NewPageInfo(edges[0].cursor, edges[len(edges)-1].cursor, hasMore)
}
