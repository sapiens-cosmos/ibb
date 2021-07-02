package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/spf13/cobra"
)

func CmdListDepositEarned() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-depositEarned",
		Short: "list all depositEarned",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDepositEarnedRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DepositEarnedAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowDepositEarned() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-depositEarned [id]",
		Short: "shows a depositEarned",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetDepositEarnedRequest{
				Id: id,
			}

			res, err := queryClient.DepositEarned(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
