package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) MintNft(goCtx context.Context, msg *types.MsgMintNft) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	// creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return nil, err
	// }
	// // if err := k.nftKeeper.MintNFT(
	// // 	ctx,
	// // 	msg.DenomID,
	// // 	msg.TokenID,
	// // 	msg.TokenNm,
	// // 	msg.TokenURI,
	// // 	msg.TokenData,
	// // 	creatorAddress,
	// // ); err != nil {
	// // 	return nil, nil
	// // }

	return &types.MsgMintNftResponse{}, nil
}
