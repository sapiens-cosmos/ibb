package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateTxHistory(goCtx context.Context, msg *types.MsgCreateTxHistory) (*types.MsgCreateTxHistoryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendTxHistory(
		ctx,
		msg.Creator,
		msg.BlockHeight,
		msg.Tx,
		msg.Asset,
		msg.Amount,
		msg.Denom,
	)

	return &types.MsgCreateTxHistoryResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTxHistory(goCtx context.Context, msg *types.MsgUpdateTxHistory) (*types.MsgUpdateTxHistoryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var txHistory = types.TxHistory{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Tx:          msg.Tx,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasTxHistory(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetTxHistoryOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTxHistory(ctx, txHistory)

	return &types.MsgUpdateTxHistoryResponse{}, nil
}

func (k msgServer) DeleteTxHistory(goCtx context.Context, msg *types.MsgDeleteTxHistory) (*types.MsgDeleteTxHistoryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasTxHistory(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetTxHistoryOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTxHistory(ctx, msg.Id)

	return &types.MsgDeleteTxHistoryResponse{}, nil
}
