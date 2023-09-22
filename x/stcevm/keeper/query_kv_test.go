package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/Jchicode/stcevm/testutil/keeper"
	"github.com/Jchicode/stcevm/testutil/nullify"
	"github.com/Jchicode/stcevm/x/stcevm/types"
)

func TestKvQuery(t *testing.T) {
	keeper, ctx := keepertest.StcevmKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestKv(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetKvRequest
		response *types.QueryGetKvResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetKvRequest{},
			response: &types.QueryGetKvResponse{Kv: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Kv(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
