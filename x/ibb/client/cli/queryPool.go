package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/spf13/cobra"
)

func CmdListPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-pool",
		Short: "list all pool",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPoolRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PoolAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-pool [id]",
		Short: "shows a pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetPoolRequest{
				Id: id,
			}

			res, err := queryClient.Pool(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdLoadPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load-pool",
		Short: "load pools",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPoolRequest{
				Pagination: pageReq,
			}

			ex_res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/load-pool", types.QuerierRoute), nil)
			if err != nil {
				return err
			}
			fmt.Println("==========Response from QueryWithData==============")
			fmt.Println(ex_res)

			res, err := queryClient.PoolAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdTestLoadPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load-pool-test",
		Short: "test list all pool",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLoadPoolRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PoolLoad(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
