package stcevm

import (
	"github.com/Jchicode/stcevm/x/stcevm/keeper"
	"github.com/Jchicode/stcevm/x/stcevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.Kv != nil {
		k.SetKv(ctx, *genState.Kv)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all kv
	kv, found := k.GetKv(ctx)
	if found {
		genesis.Kv = &kv
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
