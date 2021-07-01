package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func CmdCreateDepositEarned() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-depositEarned [blockHeight] [asset] [amount] [denom]",
		Short: "Creates a new depositEarned",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBlockHeight, _ := strconv.ParseInt(args[0], 10, 64)
			argsAsset := string(args[1])
			argsAmount, _ := strconv.ParseInt(args[2], 10, 64)
			argsDenom := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDepositEarned(clientCtx.GetFromAddress().String(), int32(argsBlockHeight), string(argsAsset), int32(argsAmount), string(argsDenom))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDepositEarned() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-depositEarned [id] [blockHeight] [asset] [amount] [denom]",
		Short: "Update a depositEarned",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBlockHeight, _ := strconv.ParseInt(args[1], 10, 64)
			argsAsset := string(args[2])
			argsAmount, _ := strconv.ParseInt(args[3], 10, 64)
			argsDenom := string(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDepositEarned(clientCtx.GetFromAddress().String(), id, int32(argsBlockHeight), string(argsAsset), int32(argsAmount), string(argsDenom))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteDepositEarned() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-depositEarned [id] [blockHeight] [asset] [amount] [denom]",
		Short: "Delete a depositEarned by id",
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

			msg := types.NewMsgDeleteDepositEarned(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
