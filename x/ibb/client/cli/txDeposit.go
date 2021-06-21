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

func CmdCreateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deposit [blockHeight] [asset] [amount] [denom]",
		Short: "Create a new deposit",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBlockHeight, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argsAsset, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsAmount, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argsDenom, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDeposit(clientCtx.GetFromAddress().String(), argsBlockHeight, argsAsset, argsAmount, argsDenom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-deposit [id] [blockHeight] [asset] [amount] [denom]",
		Short: "Update a deposit",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBlockHeight, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			argsAsset, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsAmount, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			argsDenom, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDeposit(clientCtx.GetFromAddress().String(), id, argsBlockHeight, argsAsset, argsAmount, argsDenom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-deposit [id]",
		Short: "Delete a deposit by id",
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

			msg := types.NewMsgDeleteDeposit(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
