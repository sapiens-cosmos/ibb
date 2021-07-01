package keeper

import (
	"context"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateBorrow(goCtx context.Context, msg *types.MsgCreateBorrow) (*types.MsgCreateBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var borrow = types.Borrow{
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

	for i, eachBorrow := range queryUser.Borrow {
		if eachBorrow.Denom == msg.Denom {
			queryUser.Borrow[i].Amount = queryUser.Borrow[i].Amount + msg.Amount
		}
	}
	k.SetUser(ctx, queryUser)

	// queryUser.Borrow = append(queryUser.Borrow, &borrow)
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
	queryPool.BorrowBalance = queryPool.BorrowBalance + msg.Amount

	currentTargetBorrowRatio := float64(queryPool.BorrowBalance) / float64(queryPool.DepositBalance)
	currentDepositApy := types.DepositInterest + types.DepositInterest*(currentTargetBorrowRatio-float64(types.TargetBorrowRatio)*0.01)*types.InterestFactor
	currentDepositApy = math.Max(currentDepositApy, types.MinimumDepositInterest)
	var apr types.Apr
	apr.BlockHeight = int32(ctx.BlockHeight())
	apr.DepositApy = int32(currentDepositApy * 1000000)
	apr.BorrowApy = int32(currentDepositApy / currentTargetBorrowRatio * 1000000)

	queryPool.Aprs = append(queryPool.Aprs, &apr)

	k.SetPool(ctx, queryPool)

	//TODO : add collateral logic to borrowing

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
