package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateBorrowAccrued(goCtx context.Context, msg *types.MsgCreateBorrowAccrued) (*types.MsgCreateBorrowAccruedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendBorrowAccrued(
		ctx,
		msg.Creator,
		msg.BlockHeight,
		msg.Asset,
		msg.Amount,
		msg.Denom,
	)

	return &types.MsgCreateBorrowAccruedResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBorrowAccrued(goCtx context.Context, msg *types.MsgUpdateBorrowAccrued) (*types.MsgUpdateBorrowAccruedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var borrowAccrued = types.BorrowAccrued{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasBorrowAccrued(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBorrowAccruedOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBorrowAccrued(ctx, borrowAccrued)

	return &types.MsgUpdateBorrowAccruedResponse{}, nil
}

func (k msgServer) DeleteBorrowAccrued(goCtx context.Context, msg *types.MsgDeleteBorrowAccrued) (*types.MsgDeleteBorrowAccruedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBorrowAccrued(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBorrowAccruedOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBorrowAccrued(ctx, msg.Id)

	return &types.MsgDeleteBorrowAccruedResponse{}, nil
}
