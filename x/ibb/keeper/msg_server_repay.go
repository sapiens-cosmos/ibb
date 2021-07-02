package keeper

import (
	"context"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/interest"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateRepay(goCtx context.Context, msg *types.MsgCreateRepay) (*types.MsgCreateRepayResponse, error) {
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

	poolList := k.GetAllPool(ctx)
	var queryPool types.Pool
	for _, pool := range poolList {
		if pool.Asset == msg.Asset {
			queryPool = pool
			break
		}
	}

	var userDeposit int32 = 0
	for _, deposit := range queryUser.Deposit {
		if deposit.Asset == msg.Asset {
			userDeposit = deposit.Amount
			break
		}
	}

	_, borrowAccrued := interest.GetInterests(msg.Asset, int32(ctx.BlockHeight()), queryUser.TxHistories, queryPool.Aprs)
	var repayBorrow int32 = 0
	var repayInterest int32 = 0
	if userDeposit < msg.Amount {
		repayBorrow = userDeposit
		repayInterest = int32(math.Min(float64(msg.Amount-userDeposit), float64(borrowAccrued)))
	} else {
		repayBorrow = msg.Amount
		repayInterest = 0
	}

	for i, eachBorrow := range queryUser.Borrow {
		if eachBorrow.Denom == msg.Denom {
			queryUser.Borrow[i].Amount = queryUser.Borrow[i].Amount - repayBorrow
		}
	}

	var txHistory types.TxHistory
	txHistory.BlockHeight = int32(ctx.BlockHeight())
	txHistory.Tx = "repay"
	txHistory.Asset = msg.Asset
	txHistory.Amount = repayBorrow + repayInterest
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
	queryPool.BorrowBalance = queryPool.BorrowBalance - repayBorrow

	currentTargetBorrowRatio := float64(queryPool.BorrowBalance) / float64(queryPool.DepositBalance)
	currentDepositApy := types.DepositInterest + types.DepositInterest*(currentTargetBorrowRatio-float64(types.TargetBorrowRatio)*0.01)*types.InterestFactor
	currentDepositApy = math.Max(currentDepositApy, types.MinimumDepositInterest)
	var apr types.Apr
	apr.BlockHeight = int32(ctx.BlockHeight())
	apr.DepositApy = int32(currentDepositApy * 1000000)
	apr.BorrowApy = int32(currentDepositApy / currentTargetBorrowRatio * 1000000)

	queryPool.Aprs = append(queryPool.Aprs, &apr)

	k.SetPool(ctx, queryPool)

	id := k.AppendRepay(
		ctx,
		msg.Creator,
		int32(ctx.BlockHeight()),
		msg.Asset,
		repayBorrow+repayInterest,
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
