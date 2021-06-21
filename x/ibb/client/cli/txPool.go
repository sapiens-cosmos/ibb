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

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [asset] [denom] [collatoralFactor] [depth] [APR] [users]",
		Short: "Create a new pool",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAsset, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsDenom, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsCollatoralFactor, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argsDepth, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argsAPR, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsUsers, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), argsAsset, argsDenom, argsCollatoralFactor, argsDepth, argsAPR, argsUsers)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-pool [id] [asset] [denom] [collatoralFactor] [depth] [APR] [users]",
		Short: "Update a pool",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsAsset, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsDenom, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsCollatoralFactor, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			argsDepth, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			argsAPR, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}

			argsUsers, err := cast.ToStringE(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePool(clientCtx.GetFromAddress().String(), id, argsAsset, argsDenom, argsCollatoralFactor, argsDepth, argsAPR, argsUsers)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-pool [id]",
		Short: "Delete a pool by id",
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

			msg := types.NewMsgDeletePool(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
