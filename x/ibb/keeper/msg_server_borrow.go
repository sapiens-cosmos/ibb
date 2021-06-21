package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateBorrow(goCtx context.Context, msg *types.MsgCreateBorrow) (*types.MsgCreateBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var borrow = types.Borrow{
		Creator:     msg.Creator,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	id := k.AppendBorrow(
		ctx,
		borrow,
	)

	return &types.MsgCreateBorrowResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBorrow(goCtx context.Context, msg *types.MsgUpdateBorrow) (*types.MsgUpdateBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var borrow = types.Borrow{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasBorrow(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBorrowOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBorrow(ctx, borrow)

	return &types.MsgUpdateBorrowResponse{}, nil
}

func (k msgServer) DeleteBorrow(goCtx context.Context, msg *types.MsgDeleteBorrow) (*types.MsgDeleteBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBorrow(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBorrowOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBorrow(ctx, msg.Id)

	return &types.MsgDeleteBorrowResponse{}, nil
}
