package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func CmdCreateApr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-apr [blockHeight] [depositApy] [borrowApy]",
		Short: "Creates a new apr",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBlockHeight, _ := strconv.ParseInt(args[0], 10, 64)
			argsDepositApy, _ := strconv.ParseInt(args[1], 10, 64)
			argsBorrowApy, _ := strconv.ParseInt(args[2], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateApr(clientCtx.GetFromAddress().String(), int32(argsBlockHeight), int32(argsDepositApy), int32(argsBorrowApy))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateApr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-apr [id] [blockHeight] [depositApy] [borrowApy]",
		Short: "Update a apr",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBlockHeight, _ := strconv.ParseInt(args[1], 10, 64)
			argsDepositApy, _ := strconv.ParseInt(args[2], 10, 64)
			argsBorrowApy, _ := strconv.ParseInt(args[3], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateApr(clientCtx.GetFromAddress().String(), id, int32(argsBlockHeight), int32(argsDepositApy), int32(argsBorrowApy))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteApr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-apr [id] [blockHeight] [depositApy] [borrowApy]",
		Short: "Delete a apr by id",
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

			msg := types.NewMsgDeleteApr(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
