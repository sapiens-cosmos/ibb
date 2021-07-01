package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/spf13/cobra"
)

func CmdListBorrowAccrued() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-borrowAccrued",
		Short: "list all borrowAccrued",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBorrowAccruedRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BorrowAccruedAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowBorrowAccrued() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-borrowAccrued [id]",
		Short: "shows a borrowAccrued",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetBorrowAccruedRequest{
				Id: id,
			}

			res, err := queryClient.BorrowAccrued(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
