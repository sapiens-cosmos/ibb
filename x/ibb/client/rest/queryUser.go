package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

type createUserRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
}

func loadUserHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createUserRequest
		walletAddress := mux.Vars(r)["id"]

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/load-user/%s", types.QuerierRoute, walletAddress), nil)

		if err == sdkerrors.ErrKeyNotFound {
			initialAktDeposit := types.Deposit{Asset: "akt", Denom: "uakt"}
			initialAtomDeposit := types.Deposit{Asset: "atom", Denom: "uatom"}
			initialCroDeposit := types.Deposit{Asset: "cro", Denom: "ucro"}
			initialDvpnDeposit := types.Deposit{Asset: "dvpn", Denom: "udvpn"}
			initialIrisDeposit := types.Deposit{Asset: "iris", Denom: "uiris"}
			initialXprtDeposit := types.Deposit{Asset: "xprt", Denom: "uxprt"}
			initialAktBorrow := types.Borrow{Asset: "akt", Denom: "uakt"}
			initialAtomBorrow := types.Borrow{Asset: "atom", Denom: "uatom"}
			initialCroBorrow := types.Borrow{Asset: "cro", Denom: "ucro"}
			initialDvpnBorrow := types.Borrow{Asset: "dvpn", Denom: "udvpn"}
			initialIrisBorrow := types.Borrow{Asset: "iris", Denom: "uiris"}
			initialXprtBorrow := types.Borrow{Asset: "xprt", Denom: "uxprt"}
			initialAktDepositEarned := types.DepositEarned{Amount: 0, BlockHeight: 0}
			initialAtomDepositEarned := types.DepositEarned{Amount: 0, BlockHeight: 0}
			initialCroDepositEarned := types.DepositEarned{Amount: 0, BlockHeight: 0}
			initialDvpnDepositEarned := types.DepositEarned{Amount: 0, BlockHeight: 0}
			initialIrisDepositEarned := types.DepositEarned{Amount: 0, BlockHeight: 0}
			initialXprtDepositEarned := types.DepositEarned{Amount: 0, BlockHeight: 0}
			initialAktBorrowEarned := types.BorrowAccrued{Amount: 0, BlockHeight: 0}
			initialAtomBorrowEarned := types.BorrowAccrued{Amount: 0, BlockHeight: 0}
			initialCroBorrowEarned := types.BorrowAccrued{Amount: 0, BlockHeight: 0}
			initialDvpnBorrowEarned := types.BorrowAccrued{Amount: 0, BlockHeight: 0}
			initialIrisBorrowEarned := types.BorrowAccrued{Amount: 0, BlockHeight: 0}
			initialXprtBorrowEarned := types.BorrowAccrued{Amount: 0, BlockHeight: 0}

			initialCollateral := []bool{false, false, false, false, false, false}
			initialDepositList := []*types.Deposit{&initialAktDeposit, &initialAtomDeposit, &initialCroDeposit, &initialDvpnDeposit, &initialIrisDeposit, &initialXprtDeposit}
			initialBorrowList := []*types.Borrow{&initialAktBorrow, &initialAtomBorrow, &initialCroBorrow, &initialDvpnBorrow, &initialIrisBorrow, &initialXprtBorrow}
			initialAssetBalances := []int32{0, 0, 0, 0, 0, 0}
			initialDepositEarneds := []*types.DepositEarned{&initialAktDepositEarned, &initialAtomDepositEarned, &initialCroDepositEarned, &initialDvpnDepositEarned, &initialIrisDepositEarned, &initialXprtDepositEarned}
			initialBorrowAccrueds := []*types.BorrowAccrued{&initialAktBorrowEarned, &initialAtomBorrowEarned, &initialCroBorrowEarned, &initialDvpnBorrowEarned, &initialIrisBorrowEarned, &initialXprtBorrowEarned}
			initialTxHistories := []*types.TxHistory{}

			msg := types.NewMsgCreateUser(walletAddress, initialCollateral, initialDepositList, initialBorrowList, initialAssetBalances, initialDepositEarneds, initialBorrowAccrueds, initialTxHistories)
			tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
		}

		// if err != nil {
		// 	rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		// 	return
		// }

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
