package stcevm_test

import (
	"testing"

	keepertest "github.com/Jchicode/stcevm/testutil/keeper"
	"github.com/Jchicode/stcevm/testutil/nullify"
	"github.com/Jchicode/stcevm/x/stcevm"
	"github.com/Jchicode/stcevm/x/stcevm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Kv: &types.Kv{},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StcevmKeeper(t)
	stcevm.InitGenesis(ctx, *k, genesisState)
	got := stcevm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Kv, got.Kv)
	// this line is used by starport scaffolding # genesis/test/assert
}
