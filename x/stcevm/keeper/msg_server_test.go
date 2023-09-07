package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Jchicode/stcevm/testutil/keeper"
	"github.com/Jchicode/stcevm/x/stcevm/keeper"
	"github.com/Jchicode/stcevm/x/stcevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.StcevmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
