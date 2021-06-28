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

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	// r.HandleFunc("/voter/votes/{id}", getVoteHandler(clientCtx)).Methods("GET")
	// r.HandleFunc("/voter/votes", listVoteHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/pools/{id}", getPoolHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/pools", listPoolHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibb/loadPools", loadPoolHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/ibb/loadUser/{id}", loadUserHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	// r.HandleFunc("/voter/votes", createVoteHandler(clientCtx)).Methods("POST")
	// r.HandleFunc("/voter/votes/{id}", updateVoteHandler(clientCtx)).Methods("POST")
	// r.HandleFunc("/voter/votes/{id}", deleteVoteHandler(clientCtx)).Methods("POST")

	// r.HandleFunc("/voter/polls", createPollHandler(clientCtx)).Methods("POST")
	// r.HandleFunc("/voter/polls/{id}", updatePollHandler(clientCtx)).Methods("POST")
	// r.HandleFunc("/voter/polls/{id}", deletePollHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/createDeposit", createDepositHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibb/createBorrow", createBorrowHandler(clientCtx)).Methods("POST")
}
