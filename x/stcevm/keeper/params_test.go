package keeper_test

import (
	"testing"

	testkeeper "github.com/Jchicode/stcevm/testutil/keeper"
	"github.com/Jchicode/stcevm/x/stcevm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.StcevmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
