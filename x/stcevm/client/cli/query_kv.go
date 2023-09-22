package cli

import (
	"context"

	"github.com/Jchicode/stcevm/x/stcevm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowKv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-kv",
		Short: "shows kv",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetKvRequest{}

			res, err := queryClient.Kv(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
