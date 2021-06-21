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

func CmdCreateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user [collateral] [totalDeposit] [totalBorrow] [deposit] [borrow]",
		Short: "Create a new user",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsCollateral := []bool{}

			argsTotalDeposit, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argsTotalBorrow, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argsDeposit := []*types.Deposit{}

			argsBorrow := []*types.Borrow{}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUser(clientCtx.GetFromAddress().String(), argsCollateral, argsTotalDeposit, argsTotalBorrow, argsDeposit, argsBorrow)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-user [id] [collateral] [totalDeposit] [totalBorrow] [deposit] [borrow]",
		Short: "Update a user",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsCollateral := []bool{}

			argsTotalDeposit, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			argsTotalBorrow, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			argsDeposit := []*types.Deposit{}

			argsBorrow := []*types.Borrow{}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUser(clientCtx.GetFromAddress().String(), id, argsCollateral, argsTotalDeposit, argsTotalBorrow, argsDeposit, argsBorrow)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user [id]",
		Short: "Delete a user by id",
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

			msg := types.NewMsgDeleteUser(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
