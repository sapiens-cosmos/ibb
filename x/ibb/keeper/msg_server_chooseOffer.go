package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) ChooseOffer(goCtx context.Context, msg *types.MsgChooseOffer) (*types.MsgChooseOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgChooseOfferResponse{}, nil
}
