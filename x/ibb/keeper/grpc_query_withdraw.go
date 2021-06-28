package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WithdrawAll(c context.Context, req *types.QueryAllWithdrawRequest) (*types.QueryAllWithdrawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var withdraws []*types.Withdraw
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	withdrawStore := prefix.NewStore(store, types.KeyPrefix(types.WithdrawKey))

	pageRes, err := query.Paginate(withdrawStore, req.Pagination, func(key []byte, value []byte) error {
		var withdraw types.Withdraw
		if err := k.cdc.UnmarshalBinaryBare(value, &withdraw); err != nil {
			return err
		}

		withdraws = append(withdraws, &withdraw)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWithdrawResponse{Withdraw: withdraws, Pagination: pageRes}, nil
}

func (k Keeper) Withdraw(c context.Context, req *types.QueryGetWithdrawRequest) (*types.QueryGetWithdrawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var withdraw types.Withdraw
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasWithdraw(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetWithdrawIDBytes(req.Id)), &withdraw)

	return &types.QueryGetWithdrawResponse{Withdraw: &withdraw}, nil
}
