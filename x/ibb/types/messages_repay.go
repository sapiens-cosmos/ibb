package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRepay{}

func NewMsgCreateRepay(creator string, blockHeight int32, asset string, amount int32, denom string) *MsgCreateRepay {
	return &MsgCreateRepay{
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgCreateRepay) Route() string {
	return RouterKey
}

func (msg *MsgCreateRepay) Type() string {
	return "CreateRepay"
}

func (msg *MsgCreateRepay) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRepay) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRepay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRepay{}

func NewMsgUpdateRepay(creator string, id uint64, blockHeight int32, asset string, amount int32, denom string) *MsgUpdateRepay {
	return &MsgUpdateRepay{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgUpdateRepay) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRepay) Type() string {
	return "UpdateRepay"
}

func (msg *MsgUpdateRepay) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRepay) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRepay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateRepay{}

func NewMsgDeleteRepay(creator string, id uint64) *MsgDeleteRepay {
	return &MsgDeleteRepay{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteRepay) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRepay) Type() string {
	return "DeleteRepay"
}

func (msg *MsgDeleteRepay) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRepay) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRepay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
