package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

type createClaimRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	BlockHeight string       `json:"blockHeight"`
	Asset       string       `json:"asset"`
	Amount      string       `json:"amount"`
	Denom       string       `json:"denom"`
}

func createClaimHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createClaimRequest
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

		parsedBlockHeight64, err := strconv.ParseInt(req.BlockHeight, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedBlockHeight := int32(parsedBlockHeight64)

		parsedAsset := req.Asset

		parsedAmount64, err := strconv.ParseInt(req.Amount, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedAmount := int32(parsedAmount64)

		parsedDenom := req.Denom

		msg := types.NewMsgCreateClaim(
			req.Creator,
			parsedBlockHeight,
			parsedAsset,
			parsedAmount,
			parsedDenom,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateClaimRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	BlockHeight string       `json:"blockHeight"`
	Asset       string       `json:"asset"`
	Amount      string       `json:"amount"`
	Denom       string       `json:"denom"`
}

func updateClaimHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateClaimRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedBlockHeight64, err := strconv.ParseInt(req.BlockHeight, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedBlockHeight := int32(parsedBlockHeight64)

		parsedAsset := req.Asset

		parsedAmount64, err := strconv.ParseInt(req.Amount, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedAmount := int32(parsedAmount64)

		parsedDenom := req.Denom

		msg := types.NewMsgUpdateClaim(
			req.Creator,
			id,
			parsedBlockHeight,
			parsedAsset,
			parsedAmount,
			parsedDenom,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteClaimRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteClaimHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteClaimRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeleteClaim(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
