package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/Jchicode/stcevm/testutil/keeper"
	"github.com/Jchicode/stcevm/x/stcevm/keeper"
	"github.com/Jchicode/stcevm/x/stcevm/types"
)

func TestKvMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.StcevmKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateKv{Creator: creator}
	_, err := srv.CreateKv(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetKv(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestKvMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateKv
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateKv{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateKv{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StcevmKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateKv{Creator: creator}
			_, err := srv.CreateKv(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateKv(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetKv(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestKvMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteKv
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteKv{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteKv{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StcevmKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateKv(wctx, &types.MsgCreateKv{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteKv(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetKv(ctx)
				require.False(t, found)
			}
		})
	}
}
