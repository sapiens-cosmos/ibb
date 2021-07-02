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

func (k Keeper) NftAll(c context.Context, req *types.QueryAllNftRequest) (*types.QueryAllNftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nfts []*types.Nft
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nftStore := prefix.NewStore(store, types.KeyPrefix(types.NftKey))

	pageRes, err := query.Paginate(nftStore, req.Pagination, func(key []byte, value []byte) error {
		var nft types.Nft
		if err := k.cdc.UnmarshalBinaryBare(value, &nft); err != nil {
			return err
		}

		nfts = append(nfts, &nft)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNftResponse{Nft: nfts, Pagination: pageRes}, nil
}

func (k Keeper) Nft(c context.Context, req *types.QueryGetNftRequest) (*types.QueryGetNftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nft types.Nft
	var queryNft types.NftResponse

	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasNft(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetNftIDBytes(req.Id)), &nft)

	queryNft.Collection = nft.Collection
	queryNft.Id = int32(nft.Id)
	queryNft.ImageUrl = nft.ImageUrl
	queryNft.Name = nft.Name
	queryNft.NftCreatorAddress = nft.NftCreatorAddress
	queryNft.OwnerAddress = nft.OwnerAddress

	return &types.QueryGetNftResponse{Nft: &queryNft}, nil
}

func (k Keeper) NftLoad(c context.Context, req *types.QueryLoadNftRequest) (*types.QueryLoadNftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	nftList := k.GetAllNft(ctx)

	var userNftList []*types.NftResponse
	var dashboardNftList []*types.NftResponse

	for _, nft := range nftList {
		var nftResponse types.NftResponse
		nftResponse.Collection = nft.Collection
		nftResponse.OwnerAddress = nft.OwnerAddress
		nftResponse.ImageUrl = nft.ImageUrl
		nftResponse.Name = nft.Name
		nftResponse.NftCreatorAddress = nft.NftCreatorAddress
		nftResponse.Id = int32(nft.Id)
		if nft.OwnerAddress == req.Id {
			userNftList = append(userNftList, &nftResponse)
		} else {
			dashboardNftList = append(dashboardNftList, &nftResponse)
		}
	}

	return &types.QueryLoadNftResponse{UserNft: userNftList, DashboardNft: dashboardNftList}, nil
}

func (k Keeper) NftOfferList(c context.Context, req *types.QueryNftOfferListRequest) (*types.QueryNftOfferListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	var nft types.Nft
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetNftIDBytes(req.Id)), &nft)

	var offerList []*types.Offer

	for _, offer := range nft.Offers {
		var pinnedOffer = types.Offer{
			Denom:           offer.Denom,
			Amount:          offer.Amount,
			PaybackAmount:   offer.PaybackAmount,
			PaybackDuration: offer.PaybackDuration,
			OfferStartAt:    offer.OfferStartAt,
			NftId:           offer.NftId,
			OfferCreator:    offer.OfferCreator,
			Id:              offer.Id,
		}
		offerList = append(offerList, &pinnedOffer)
	}

	return &types.QueryNftOfferListResponse{OfferList: offerList}, nil
}
