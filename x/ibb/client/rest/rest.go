package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers voter-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/ibb/claims/{id}", getClaimHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/claims", listClaimHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/txHistories/{id}", getTxHistoryHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/txHistories", listTxHistoryHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/borrowAccrueds/{id}", getBorrowAccruedHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/borrowAccrueds", listBorrowAccruedHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/depositEarneds/{id}", getDepositEarnedHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/depositEarneds", listDepositEarnedHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/aprs/{id}", getAprHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/aprs", listAprHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/repays/{id}", getRepayHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/repays", listRepayHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/withdraws/{id}", getWithdrawHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/withdraws", listWithdrawHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/pools/{id}", getPoolHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/pools", listPoolHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/loadPools", loadPoolHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/loadUser/{id}", loadUserHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/collections", listCollectionHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/ibb/claims", createClaimHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/claims/{id}", updateClaimHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/claims/{id}", deleteClaimHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/ibb/txHistories", createTxHistoryHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/txHistories/{id}", updateTxHistoryHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/txHistories/{id}", deleteTxHistoryHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/ibb/borrowAccrueds", createBorrowAccruedHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/borrowAccrueds/{id}", updateBorrowAccruedHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/borrowAccrueds/{id}", deleteBorrowAccruedHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/ibb/depositEarneds", createDepositEarnedHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/depositEarneds/{id}", updateDepositEarnedHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/depositEarneds/{id}", deleteDepositEarnedHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/ibb/aprs", createAprHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/aprs/{id}", updateAprHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/aprs/{id}", deleteAprHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/ibb/repays", createRepayHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/repays/{id}", updateRepayHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/repays/{id}", deleteRepayHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/ibb/withdraws", createWithdrawHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/withdraws/{id}", updateWithdrawHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/withdraws/{id}", deleteWithdrawHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/ibb/createDeposit", createDepositHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/createBorrow", createBorrowHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/createWithdraw", createWithdrawHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/createRepay", createRepayHandler(clientCtx)).Methods("POST")
}
