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

func (k Keeper) RepayAll(c context.Context, req *types.QueryAllRepayRequest) (*types.QueryAllRepayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var repays []*types.Repay
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	repayStore := prefix.NewStore(store, types.KeyPrefix(types.RepayKey))

	pageRes, err := query.Paginate(repayStore, req.Pagination, func(key []byte, value []byte) error {
		var repay types.Repay
		if err := k.cdc.UnmarshalBinaryBare(value, &repay); err != nil {
			return err
		}

		repays = append(repays, &repay)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRepayResponse{Repay: repays, Pagination: pageRes}, nil
}

func (k Keeper) Repay(c context.Context, req *types.QueryGetRepayRequest) (*types.QueryGetRepayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var repay types.Repay
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasRepay(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetRepayIDBytes(req.Id)), &repay)

	return &types.QueryGetRepayResponse{Repay: &repay}, nil
}
