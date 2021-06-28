package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateRepay(goCtx context.Context, msg *types.MsgCreateRepay) (*types.MsgCreateRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendRepay(
		ctx,
		msg.Creator,
		msg.BlockHeight,
		msg.Asset,
		msg.Amount,
		msg.Denom,
	)

	return &types.MsgCreateRepayResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateRepay(goCtx context.Context, msg *types.MsgUpdateRepay) (*types.MsgUpdateRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var repay = types.Repay{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasRepay(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetRepayOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetRepay(ctx, repay)

	return &types.MsgUpdateRepayResponse{}, nil
}

func (k msgServer) DeleteRepay(goCtx context.Context, msg *types.MsgDeleteRepay) (*types.MsgDeleteRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasRepay(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetRepayOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRepay(ctx, msg.Id)

	return &types.MsgDeleteRepayResponse{}, nil
}
