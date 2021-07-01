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
		Use:   "create-nft [collection] [ownerAddress] [imageUrl] [name]",
		Short: "Create a new nft",
		Args:  cobra.ExactArgs(4),
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

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), argsCollection, argsOwnerAddress, argsImageUrl, argsName)
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

func CmdMintNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-nft [denomID] [tokenID] [tokenNm] [tokenURI] [tokenData] [recipient]",
		Short: "Mint a new nft",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDenomID, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsTokenID, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsTokenNm, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsTokenURI, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			// argsTokenData, err := cast.ToStringE(args[4])
			// if err != nil {
			// 	return err
			// }
			// argsRecipient, err := cast.ToStringE(args[5])
			// if err != nil {
			// 	return err
			// }

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), argsDenomID, argsTokenID, argsTokenNm, argsTokenURI)
			// msg := types.NewMsgMint()
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
