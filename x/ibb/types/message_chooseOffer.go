package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgChooseOffer{}

func NewMsgChooseOffer(creator string, nftId int32, offerId int32) *MsgChooseOffer {
	return &MsgChooseOffer{
		Creator: creator,
		NftId:   nftId,
		OfferId: offerId,
	}
}

func (msg *MsgChooseOffer) Route() string {
	return RouterKey
}

func (msg *MsgChooseOffer) Type() string {
	return "ChooseOffer"
}

func (msg *MsgChooseOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChooseOffer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChooseOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
