package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nft",
		Short: "list all nft",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNftRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NftAll(context.Background(), params)
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

func CmdShowNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft [id]",
		Short: "shows a nft",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetNftRequest{
				Id: id,
			}

			res, err := queryClient.Nft(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdLoadNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load-nft [id]",
		Short: "Load Nfts in Endpoint format",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			walletAddress, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryLoadNftRequest{
				Id: walletAddress,
			}

			res, err := queryClient.NftLoad(context.Background(), params)
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

func CmdNftOfferList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-offer-list [id]",
		Short: "shows the list of offers for a specific nft",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryNftOfferListRequest{
				Id: id,
			}

			res, err := queryClient.NftOfferList(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
