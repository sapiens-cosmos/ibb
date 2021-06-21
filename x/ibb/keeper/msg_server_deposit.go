package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateDeposit(goCtx context.Context, msg *types.MsgCreateDeposit) (*types.MsgCreateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var deposit = types.Deposit{
		Creator:     msg.Creator,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	id := k.AppendDeposit(
		ctx,
		deposit,
	)

	return &types.MsgCreateDepositResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateDeposit(goCtx context.Context, msg *types.MsgUpdateDeposit) (*types.MsgUpdateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var deposit = types.Deposit{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasDeposit(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetDepositOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDeposit(ctx, deposit)

	return &types.MsgUpdateDepositResponse{}, nil
}

func (k msgServer) DeleteDeposit(goCtx context.Context, msg *types.MsgDeleteDeposit) (*types.MsgDeleteDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasDeposit(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetDepositOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDeposit(ctx, msg.Id)

	return &types.MsgDeleteDepositResponse{}, nil
}
