package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateWithdraw{}

func NewMsgCreateWithdraw(creator string, blockHeight int32, asset string, amount int32, denom string) *MsgCreateWithdraw {
	return &MsgCreateWithdraw{
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgCreateWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgCreateWithdraw) Type() string {
	return "CreateWithdraw"
}

func (msg *MsgCreateWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWithdraw{}

func NewMsgUpdateWithdraw(creator string, id uint64, blockHeight int32, asset string, amount int32, denom string) *MsgUpdateWithdraw {
	return &MsgUpdateWithdraw{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgUpdateWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWithdraw) Type() string {
	return "UpdateWithdraw"
}

func (msg *MsgUpdateWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateWithdraw{}

func NewMsgDeleteWithdraw(creator string, id uint64) *MsgDeleteWithdraw {
	return &MsgDeleteWithdraw{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWithdraw) Type() string {
	return "DeleteWithdraw"
}

func (msg *MsgDeleteWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
