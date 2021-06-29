package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateWithdraw(goCtx context.Context, msg *types.MsgCreateWithdraw) (*types.MsgCreateWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	feeCoins, err := sdk.ParseCoinsNormalized(fmt.Sprint(msg.Amount, msg.Denom))
	if err != nil {
		return nil, err
	}
	userList := k.GetAllUser(ctx)

	var queryUser types.User
	for _, user := range userList {
		if user.Creator == msg.Creator {
			queryUser = user
		}
	}

	for i, eachDeposit := range queryUser.Deposit {
		if eachDeposit.Denom == msg.Denom {
			queryUser.Deposit[i].Amount = queryUser.Deposit[i].Amount - msg.Amount
		}
	}
	k.SetUser(ctx, queryUser)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.bankKeeper.AddCoins(ctx, creatorAddress, feeCoins); err != nil {
		return nil, err
	}

	poolList := k.GetAllPool(ctx)
	var queryPool types.Pool
	for _, pool := range poolList {
		if pool.Asset == msg.Asset {
			queryPool = pool
		}
	}
	queryPool.DepositBalance = queryPool.DepositBalance - msg.Amount
	k.SetPool(ctx, queryPool)

	//TODO: add logic for paying back with interest
	id := k.AppendWithdraw(
		ctx,
		msg.Creator,
		int32(ctx.BlockHeight()),
		msg.Asset,
		msg.Amount,
		msg.Denom,
	)

	return &types.MsgCreateWithdrawResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateWithdraw(goCtx context.Context, msg *types.MsgUpdateWithdraw) (*types.MsgUpdateWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var withdraw = types.Withdraw{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasWithdraw(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetWithdrawOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWithdraw(ctx, withdraw)

	return &types.MsgUpdateWithdrawResponse{}, nil
}

func (k msgServer) DeleteWithdraw(goCtx context.Context, msg *types.MsgDeleteWithdraw) (*types.MsgDeleteWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasWithdraw(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetWithdrawOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWithdraw(ctx, msg.Id)

	return &types.MsgDeleteWithdrawResponse{}, nil
}
