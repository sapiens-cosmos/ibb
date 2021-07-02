package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/interest"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateClaim(goCtx context.Context, msg *types.MsgCreateClaim) (*types.MsgCreateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userList := k.GetAllUser(ctx)
	var queryUser types.User
	for _, user := range userList {
		if user.Creator == msg.Creator {
			queryUser = user
			break
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

	depositEarned, _ := interest.GetInterests(msg.Asset, int32(ctx.BlockHeight()), queryUser.TxHistories, queryPool.Aprs)

	var txHistory types.TxHistory
	txHistory.BlockHeight = int32(ctx.BlockHeight())
	txHistory.Tx = "claim"
	txHistory.Asset = msg.Asset
	txHistory.Amount = depositEarned
	txHistory.Denom = msg.Denom
	queryUser.TxHistories = append(queryUser.TxHistories, &txHistory)

	k.SetUser(ctx, queryUser)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	earnedCoins, err := sdk.ParseCoinsNormalized(fmt.Sprint(depositEarned, msg.Denom))
	if err != nil {
		return nil, err
	}
	if err := k.bankKeeper.AddCoins(ctx, creatorAddress, earnedCoins); err != nil {
		return nil, err
	}

	id := k.AppendClaim(
		ctx,
		msg.Creator,
		msg.BlockHeight,
		msg.Asset,
		depositEarned,
		msg.Denom,
	)

	return &types.MsgCreateClaimResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateClaim(goCtx context.Context, msg *types.MsgUpdateClaim) (*types.MsgUpdateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var claim = types.Claim{
		Creator:     msg.Creator,
		Id:          msg.Id,
		BlockHeight: msg.BlockHeight,
		Asset:       msg.Asset,
		Amount:      msg.Amount,
		Denom:       msg.Denom,
	}

	// Checks that the element exists
	if !k.HasClaim(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetClaimOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetClaim(ctx, claim)

	return &types.MsgUpdateClaimResponse{}, nil
}

func (k msgServer) DeleteClaim(goCtx context.Context, msg *types.MsgDeleteClaim) (*types.MsgDeleteClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasClaim(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetClaimOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveClaim(ctx, msg.Id)

	return &types.MsgDeleteClaimResponse{}, nil
}
