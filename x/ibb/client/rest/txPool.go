package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

type createPoolRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
	Title   string       `json:"title"`
	Options []string     `json:"options"`
}

func createPoolHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPoolRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// parsedTitle := req.Title

		// parsedOptions := req.Options

		// msg := types.NewMsgCreatePool(
		// 	req.Creator,
		// 	parsedTitle,
		// 	parsedOptions,
		// )

		// tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type createDepositRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
	Asset   string       `json:"asset"`
	Amount  int32        `json:"amount"`
	Denom   string       `json:"denom"`
}

func createDepositHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createDepositRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgCreateDeposit(
			req.Creator,
			int32(clientCtx.Height),
			req.Asset,
			req.Amount,
			req.Denom,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type createBorrowRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
	Asset   string       `json:"asset"`
	Amount  int32        `json:"amount"`
	Denom   string       `json:"denom"`
}

func createBorrowHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createDepositRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgCreateBorrow(
			req.Creator,
			int32(clientCtx.Height),
			req.Asset,
			req.Amount,
			req.Denom,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

//TODO: implement update Pool Request, delete Pool request logic
// type updatePollRequest struct {
// 	BaseReq rest.BaseReq `json:"base_req"`
// 	Creator string       `json:"creator"`
// 	Title   string       `json:"title"`
// 	Options []string     `json:"options"`
// }

// func updatePollHandler(clientCtx client.Context) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
// 		if err != nil {
// 			return
// 		}

// 		var req updatePollRequest
// 		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
// 			return
// 		}

// 		baseReq := req.BaseReq.Sanitize()
// 		if !baseReq.ValidateBasic(w) {
// 			return
// 		}

// 		_, err = sdk.AccAddressFromBech32(req.Creator)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		parsedTitle := req.Title

// 		parsedOptions := req.Options

// 		msg := types.NewMsgUpdatePoll(
// 			req.Creator,
// 			id,
// 			parsedTitle,
// 			parsedOptions,
// 		)

// 		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
// 	}
// }

// type deletePollRequest struct {
// 	BaseReq rest.BaseReq `json:"base_req"`
// 	Creator string       `json:"creator"`
// }

// func deletePollHandler(clientCtx client.Context) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
// 		if err != nil {
// 			return
// 		}

// 		var req deletePollRequest
// 		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
// 			return
// 		}

// 		baseReq := req.BaseReq.Sanitize()
// 		if !baseReq.ValidateBasic(w) {
// 			return
// 		}

// 		_, err = sdk.AccAddressFromBech32(req.Creator)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		msg := types.NewMsgDeletePoll(
// 			req.Creator,
// 			id,
// 		)

// 		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
// 	}
// }
