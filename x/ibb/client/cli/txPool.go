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
		Use:   "create-pool [asset] [denom] [collatoralFactor] [depositBalance] [borrowBalance]",
		Short: "Create a new pool",
		Args:  cobra.ExactArgs(5),
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
			argsDepositBalance, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argsBorrowBalance, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			argsUsers := []*types.User{}

			argsAprs := []*types.Apr{}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), argsAsset, argsDenom, argsCollatoralFactor, argsDepositBalance, argsBorrowBalance, argsUsers, argsAprs)
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

func CmdCreateAllPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-all-pool",
		Short: "Create a new pool",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msgAtom := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "ATOM", "uatom", 75, 1000000000, 80000000, []*types.User{}, []*types.Apr{})
			msgIris := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "IRIS", "uatom", 80, 1000000000, 80000000, []*types.User{}, []*types.Apr{})
			msgDvpn := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "DVPN", "udvpn", 70, 1000000000, 80000000, []*types.User{}, []*types.Apr{})
			msgXprt := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "XPRT", "uxprt", 80, 1000000000, 80000000, []*types.User{}, []*types.Apr{})
			msgCro := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "CRO", "ucro", 80, 1000000000, 80000000, []*types.User{}, []*types.Apr{})
			msgAkt := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "AKT", "uakt", 80, 1000000000, 80000000, []*types.User{}, []*types.Apr{})

			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgAkt)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgAtom)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCro)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgDvpn)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgIris)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgXprt)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
