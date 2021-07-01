package keeper

import (
	"context"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

//TODO: takes asset out of User, put it in the pool
func (k msgServer) CreateDeposit(goCtx context.Context, msg *types.MsgCreateDeposit) (*types.MsgCreateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var deposit = types.Deposit{
		Creator:     msg.Creator,
		BlockHeight: int32(ctx.BlockHeight()),
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

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
			queryUser.Deposit[i].Amount = queryUser.Deposit[i].Amount + msg.Amount
		}
	}

	var txHistory types.TxHistory
	txHistory.BlockHeight = int32(ctx.BlockHeight())
	txHistory.Tx = "deposit"
	txHistory.Asset = msg.Asset
	txHistory.Amount = msg.Amount
	txHistory.Denom = msg.Denom
	queryUser.TxHistories = append(queryUser.TxHistories, &txHistory)

	k.SetUser(ctx, queryUser)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SubtractCoins(ctx, creatorAddress, feeCoins); err != nil {
		return nil, err
	}

	poolList := k.GetAllPool(ctx)
	var queryPool types.Pool
	for _, pool := range poolList {
		if pool.Asset == msg.Asset {
			queryPool = pool
		}
	}
	queryPool.DepositBalance = queryPool.DepositBalance + msg.Amount

	currentTargetBorrowRatio := float64(queryPool.BorrowBalance) / float64(queryPool.DepositBalance)
	currentDepositApy := types.DepositInterest + types.DepositInterest*(currentTargetBorrowRatio-float64(types.TargetBorrowRatio)*0.01)*types.InterestFactor
	currentDepositApy = math.Max(currentDepositApy, types.MinimumDepositInterest)
	var apr types.Apr
	apr.BlockHeight = int32(ctx.BlockHeight())
	apr.DepositApy = int32(currentDepositApy * 1000000)
	apr.BorrowApy = int32(currentDepositApy / currentTargetBorrowRatio * 1000000)

	queryPool.Aprs = append(queryPool.Aprs, &apr)

	k.SetPool(ctx, queryPool)

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
