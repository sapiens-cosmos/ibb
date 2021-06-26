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
			argsAPR := []*types.InterfaceApr{}
			argsUsers := []*types.User{}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), argsAsset, argsDenom, argsCollatoralFactor, argsDepositBalance, argsBorrowBalance, argsAPR, argsUsers)
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
			msgAtom := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "ATOM", "uatom", 75, 812012, 20000, []*types.InterfaceApr{}, []*types.User{})
			msgIris := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "IRIS", "uatom", 80, 20000, 10000, []*types.InterfaceApr{}, []*types.User{})
			msgOsmo := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "OSMO", "uosmo", 80, 20000, 10000, []*types.InterfaceApr{}, []*types.User{})
			msgRegen := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "REGEN", "uregen", 80, 20000, 10000, []*types.InterfaceApr{}, []*types.User{})
			msgDvpn := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "DVPN", "udvpn", 70, 20000, 10000, []*types.InterfaceApr{}, []*types.User{})
			msgXprt := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "XPRT", "uxprt", 80, 20000, 10000, []*types.InterfaceApr{}, []*types.User{})
			msgCro := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "CRO", "ucro", 80, 20000, 10000, []*types.InterfaceApr{}, []*types.User{})
			msgAkt := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), "AKT", "uakt", 80, 20000, 10000, []*types.InterfaceApr{}, []*types.User{})

			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgAtom)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgIris)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgOsmo)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgRegen)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgDvpn)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgXprt)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCro)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgAkt)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
