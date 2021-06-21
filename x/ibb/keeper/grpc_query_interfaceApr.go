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

func (k Keeper) InterfaceAprAll(c context.Context, req *types.QueryAllInterfaceAprRequest) (*types.QueryAllInterfaceAprResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var interfaceAprs []*types.InterfaceApr
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	interfaceAprStore := prefix.NewStore(store, types.KeyPrefix(types.InterfaceAprKey))

	pageRes, err := query.Paginate(interfaceAprStore, req.Pagination, func(key []byte, value []byte) error {
		var interfaceApr types.InterfaceApr
		if err := k.cdc.UnmarshalBinaryBare(value, &interfaceApr); err != nil {
			return err
		}

		interfaceAprs = append(interfaceAprs, &interfaceApr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInterfaceAprResponse{InterfaceApr: interfaceAprs, Pagination: pageRes}, nil
}

func (k Keeper) InterfaceApr(c context.Context, req *types.QueryGetInterfaceAprRequest) (*types.QueryGetInterfaceAprResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var interfaceApr types.InterfaceApr
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasInterfaceApr(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetInterfaceAprIDBytes(req.Id)), &interfaceApr)

	return &types.QueryGetInterfaceAprResponse{InterfaceApr: &interfaceApr}, nil
}
