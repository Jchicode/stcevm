package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/Jchicode/stcevm/testutil/keeper"
	"github.com/Jchicode/stcevm/testutil/nullify"
	"github.com/Jchicode/stcevm/x/stcevm/keeper"
	"github.com/Jchicode/stcevm/x/stcevm/types"
)

func createTestKv(keeper *keeper.Keeper, ctx sdk.Context) types.Kv {
	item := types.Kv{}
	keeper.SetKv(ctx, item)
	return item
}

func TestKvGet(t *testing.T) {
	keeper, ctx := keepertest.StcevmKeeper(t)
	item := createTestKv(keeper, ctx)
	rst, found := keeper.GetKv(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestKvRemove(t *testing.T) {
	keeper, ctx := keepertest.StcevmKeeper(t)
	createTestKv(keeper, ctx)
	keeper.RemoveKv(ctx)
	_, found := keeper.GetKv(ctx)
	require.False(t, found)
}
