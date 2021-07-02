package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func CmdCreateTxHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-txHistory [blockHeight] [tx] [asset] [amount] [denom]",
		Short: "Creates a new txHistory",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBlockHeight, _ := strconv.ParseInt(args[0], 10, 64)
			argsTx := string(args[1])
			argsAsset := string(args[2])
			argsAmount, _ := strconv.ParseInt(args[3], 10, 64)
			argsDenom := string(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTxHistory(clientCtx.GetFromAddress().String(), int32(argsBlockHeight), string(argsTx), string(argsAsset), int32(argsAmount), string(argsDenom))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateTxHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-txHistory [id] [blockHeight] [tx] [asset] [amount] [denom]",
		Short: "Update a txHistory",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBlockHeight, _ := strconv.ParseInt(args[1], 10, 64)
			argsTx := string(args[2])
			argsAsset := string(args[3])
			argsAmount, _ := strconv.ParseInt(args[4], 10, 64)
			argsDenom := string(args[5])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTxHistory(clientCtx.GetFromAddress().String(), id, int32(argsBlockHeight), string(argsTx), string(argsAsset), int32(argsAmount), string(argsDenom))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteTxHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-txHistory [id] [blockHeight] [tx] [asset] [amount] [denom]",
		Short: "Delete a txHistory by id",
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

			msg := types.NewMsgDeleteTxHistory(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
