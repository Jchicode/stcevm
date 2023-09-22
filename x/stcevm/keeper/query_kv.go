package keeper

import (
	"context"

	"github.com/Jchicode/stcevm/x/stcevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Kv(goCtx context.Context, req *types.QueryGetKvRequest) (*types.QueryGetKvResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetKv(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKvResponse{Kv: val}, nil
}
