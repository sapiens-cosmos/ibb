package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateOffer(goCtx context.Context, msg *types.MsgCreateOffer) (*types.MsgCreateOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var newOffer = types.Offer{
		Amount:          msg.Amount,
		Denom:           msg.Denom,
		PaybackAmount:   msg.PaybackAmount,
		PaybackDuration: msg.PaybackDuration,
		OfferStartAt:    time.Now().Unix(),
		NftId:           int32(msg.NftId),
	}
	queryNft := k.GetNft(ctx, msg.NftId)

	queryNft.Offers = append(queryNft.Offers, &newOffer)
	k.SetNft(ctx, queryNft)

	return &types.MsgCreateOfferResponse{}, nil
}
