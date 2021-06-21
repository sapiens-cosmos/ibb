package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pool = types.Pool{
		Creator:          msg.Creator,
		Asset:            msg.Asset,
		Denom:            msg.Denom,
		CollatoralFactor: msg.CollatoralFactor,
		Depth:            msg.Depth,
		APR:              msg.APR,
		Users:            msg.Users,
	}

	id := k.AppendPool(
		ctx,
		pool,
	)

	return &types.MsgCreatePoolResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pool = types.Pool{
		Creator:          msg.Creator,
		Id:               msg.Id,
		Asset:            msg.Asset,
		Denom:            msg.Denom,
		CollatoralFactor: msg.CollatoralFactor,
		Depth:            msg.Depth,
		APR:              msg.APR,
		Users:            msg.Users,
	}

	// Checks that the element exists
	if !k.HasPool(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPoolOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPool(ctx, pool)

	return &types.MsgUpdatePoolResponse{}, nil
}

func (k msgServer) DeletePool(goCtx context.Context, msg *types.MsgDeletePool) (*types.MsgDeletePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasPool(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetPoolOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePool(ctx, msg.Id)

	return &types.MsgDeletePoolResponse{}, nil
}
