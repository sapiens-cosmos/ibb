package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// "github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func listCollection(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	// msgs := keeper.GetAllPool(ctx)
	queryNft := keeper.GetNft(ctx, uint64(8))
	queryNft.Offers[0].OfferStartAt = time.Now().Unix()
	selectedOffer := queryNft.Offers[0]

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, selectedOffer)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
