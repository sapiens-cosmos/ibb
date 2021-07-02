package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) CreateOffer(goCtx context.Context, msg *types.MsgCreateOffer) (*types.MsgCreateOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateOfferResponse{}, nil
}
