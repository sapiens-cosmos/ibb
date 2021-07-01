package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateDepositEarned(goCtx context.Context, msg *types.MsgCreateDepositEarned) (*types.MsgCreateDepositEarnedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendDepositEarned(
		ctx,
		msg.Creator,
		msg.BlockHeight,
		msg.Asset,
		msg.Amount,
		msg.Denom,
	)

	return &types.MsgCreateDepositEarnedResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateDepositEarned(goCtx context.Context, msg *types.MsgUpdateDepositEarned) (*types.MsgUpdateDepositEarnedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var depositEarned = types.DepositEarned{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasDepositEarned(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetDepositEarnedOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDepositEarned(ctx, depositEarned)

	return &types.MsgUpdateDepositEarnedResponse{}, nil
}

func (k msgServer) DeleteDepositEarned(goCtx context.Context, msg *types.MsgDeleteDepositEarned) (*types.MsgDeleteDepositEarnedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasDepositEarned(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetDepositEarnedOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDepositEarned(ctx, msg.Id)

	return &types.MsgDeleteDepositEarnedResponse{}, nil
}
