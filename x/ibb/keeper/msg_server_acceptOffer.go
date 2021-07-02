package keeper

import (
	"context"
	"time"

	// "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func (k msgServer) AcceptOffer(goCtx context.Context, msg *types.MsgAcceptOffer) (*types.MsgAcceptOfferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	queryNft := k.GetNft(ctx, uint64(msg.NftId))
	queryNft.Offers[msg.OfferId].OfferStartAt = time.Now().Unix()
	selectedOffer := queryNft.Offers[msg.OfferId]
	queryNft.SelectedOffer = selectedOffer
	k.SetNft(ctx, queryNft)

	//send coins
	// feeCoins, err := sdk.ParseCoinsNormalized(fmt.Sprint(selectedOffer.Amount, selectedOffer.Denom))
	// if err != nil {
	// 	return nil, err
	// }
	// if err := k.bankKeeper.SendCoins(ctx, sdk.AccAddress(msg.Creator), sdk.AccAddress(queryNft.OwnerAddress), feeCoins); err != nil {
	// 	return nil, err
	// }

	//append Coins to user Borrow
	userList := k.GetAllUser(ctx)
	var queryUser types.User

	for _, user := range userList {
		if user.Creator == msg.Creator {
			// if user.Creator == queryNft.OwnerAddress {
			queryUser = user
		}
	}

	// for i, _ := range queryUser.Borrow {
	// 	// for i, eachBorrow := range queryUser.Borrow {
	// 	queryUser.Borrow[i].Amount = queryUser.Borrow[i].Amount + selectedOffer.Amount
	// 	// if eachBorrow.Denom == selectedOffer.Denom {
	// 	// }
	// }
	for i, eachBorrow := range queryUser.Borrow {
		if eachBorrow.Denom == "uatom" {
			queryUser.Borrow[i].Amount = queryUser.Borrow[i].Amount + 50
		}
	}

	k.SetUser(ctx, queryUser)

	//TODO: Logic that changes owner of the nft to the module address
	//also check if module address is fixed
	return &types.MsgAcceptOfferResponse{}, nil
}
