package keeper

import (
	"context"

	"github.com/Jchicode/stcevm/x/stcevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKv(goCtx context.Context, msg *types.MsgCreateKv) (*types.MsgCreateKvResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetKv(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var kv = types.Kv{
		Creator: msg.Creator,
	}

	k.SetKv(
		ctx,
		kv,
	)
	return &types.MsgCreateKvResponse{}, nil
}

func (k msgServer) UpdateKv(goCtx context.Context, msg *types.MsgUpdateKv) (*types.MsgUpdateKvResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKv(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var kv = types.Kv{
		Creator: msg.Creator,
	}

	k.SetKv(ctx, kv)

	return &types.MsgUpdateKvResponse{}, nil
}

func (k msgServer) DeleteKv(goCtx context.Context, msg *types.MsgDeleteKv) (*types.MsgDeleteKvResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKv(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveKv(ctx)

	return &types.MsgDeleteKvResponse{}, nil
}
