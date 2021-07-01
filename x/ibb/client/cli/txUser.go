package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func CmdCreateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "Create a new user",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			initialAktDeposit := types.Deposit{Asset: "AKT", Denom: "uakt"}
			initialAtomDeposit := types.Deposit{Asset: "ATOM", Denom: "uatom"}
			initialCroDeposit := types.Deposit{Asset: "CRO", Denom: "ucro"}
			initialDvpnDeposit := types.Deposit{Asset: "DVPN", Denom: "udvpn"}
			initialIrisDeposit := types.Deposit{Asset: "IRIS", Denom: "uiris"}
			initialXprtDeposit := types.Deposit{Asset: "XPRT", Denom: "uxprt"}
			initialAktBorrow := types.Borrow{Asset: "AKT", Denom: "uakt"}
			initialAtomBorrow := types.Borrow{Asset: "ATOM", Denom: "uatom"}
			initialCroBorrow := types.Borrow{Asset: "CRO", Denom: "ucro"}
			initialDvpnBorrow := types.Borrow{Asset: "DVPN", Denom: "udvpn"}
			initialIrisBorrow := types.Borrow{Asset: "IRIS", Denom: "uiris"}
			initialXprtBorrow := types.Borrow{Asset: "XPRT", Denom: "uxprt"}
			initialAktDepositEarned := types.DepositEarned{Asset: "AKT", Denom: "uakt"}
			initialAtomDepositEarned := types.DepositEarned{Asset: "ATOM", Denom: "uatom"}
			initialCroDepositEarned := types.DepositEarned{Asset: "CRO", Denom: "ucro"}
			initialDvpnDepositEarned := types.DepositEarned{Asset: "IRIS", Denom: "uiris"}
			initialIrisDepositEarned := types.DepositEarned{Asset: "IRIS", Denom: "uiris"}
			initialXprtDepositEarned := types.DepositEarned{Asset: "XPRT", Denom: "uxprt"}
			initialAktBorrowEarned := types.BorrowAccrued{Asset: "AKT", Denom: "uakt"}
			initialAtomBorrowEarned := types.BorrowAccrued{Asset: "ATOM", Denom: "uatom"}
			initialCroBorrowEarned := types.BorrowAccrued{Asset: "CRO", Denom: "ucro"}
			initialDvpnBorrowEarned := types.BorrowAccrued{Asset: "IRIS", Denom: "uiris"}
			initialIrisBorrowEarned := types.BorrowAccrued{Asset: "IRIS", Denom: "uiris"}
			initialXprtBorrowEarned := types.BorrowAccrued{Asset: "XPRT", Denom: "uxprt"}

			argsCollateral := []bool{false, false, false, false, false, false}
			argsDeposit := []*types.Deposit{&initialAktDeposit, &initialAtomDeposit, &initialCroDeposit, &initialDvpnDeposit, &initialIrisDeposit, &initialXprtDeposit}
			argsBorrow := []*types.Borrow{&initialAktBorrow, &initialAtomBorrow, &initialCroBorrow, &initialDvpnBorrow, &initialIrisBorrow, &initialXprtBorrow}
			argsAssetBalances := []int32{2000, 1000, 1000, 1000, 1000, 1000, 1000, 1000}
			argsDepositEarneds := []*types.DepositEarned{&initialAktDepositEarned, &initialAtomDepositEarned, &initialCroDepositEarned, &initialDvpnDepositEarned, &initialIrisDepositEarned, &initialXprtDepositEarned}
			argsBorrowAccrueds := []*types.BorrowAccrued{&initialAktBorrowEarned, &initialAtomBorrowEarned, &initialCroBorrowEarned, &initialDvpnBorrowEarned, &initialIrisBorrowEarned, &initialXprtBorrowEarned}
			argsTxHistories := []*types.TxHistory{}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUser(clientCtx.GetFromAddress().String(), argsCollateral, argsDeposit, argsBorrow, argsAssetBalances, argsDepositEarneds, argsBorrowAccrueds, argsTxHistories)
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
