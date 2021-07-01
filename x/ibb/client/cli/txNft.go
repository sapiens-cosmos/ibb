package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func CmdCreateNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-nft [collection] [ownerAddress] [imageUrl] [name] [nftCreatorAddress]",
		Short: "Create a new nft",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsCollection, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsOwnerAddress, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsImageUrl, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsName, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsNftCreatorAddress, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), argsCollection, argsOwnerAddress, argsImageUrl, argsName, argsNftCreatorAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-nft [id] [collection] [ownerAddress] [imageUrl] [name]",
		Short: "Update a nft",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsCollection, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsOwnerAddress, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsImageUrl, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsName, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNft(clientCtx.GetFromAddress().String(), id, argsCollection, argsOwnerAddress, argsImageUrl, argsName)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-nft [id]",
		Short: "Delete a nft by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteNft(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreateMockNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-all-nft",
		Short: "Create a set of mock nfts for basic project setup",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msgNft := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Collection Name", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "google.com", "Nft Name", "nftCreatorAddress")
			msgNft1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Collection Name", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "google.com", "Nft Name", "nftCreatorAddress")

			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgNft)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgNft1)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
