package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateInterfaceApr(goCtx context.Context, msg *types.MsgCreateInterfaceApr) (*types.MsgCreateInterfaceAprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var interfaceApr = types.InterfaceApr{
		Creator:   msg.Creator,
		BlockTime: msg.BlockTime,
		TimeApr:   msg.TimeApr,
	}

	id := k.AppendInterfaceApr(
		ctx,
		interfaceApr,
	)

	return &types.MsgCreateInterfaceAprResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateInterfaceApr(goCtx context.Context, msg *types.MsgUpdateInterfaceApr) (*types.MsgUpdateInterfaceAprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var interfaceApr = types.InterfaceApr{
		Creator:   msg.Creator,
		Id:        msg.Id,
		BlockTime: msg.BlockTime,
		TimeApr:   msg.TimeApr,
	}

	// Checks that the element exists
	if !k.HasInterfaceApr(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetInterfaceAprOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetInterfaceApr(ctx, interfaceApr)

	return &types.MsgUpdateInterfaceAprResponse{}, nil
}

func (k msgServer) DeleteInterfaceApr(goCtx context.Context, msg *types.MsgDeleteInterfaceApr) (*types.MsgDeleteInterfaceAprResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasInterfaceApr(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetInterfaceAprOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveInterfaceApr(ctx, msg.Id)

	return &types.MsgDeleteInterfaceAprResponse{}, nil
}
