package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDepositEarned{}

func NewMsgCreateDepositEarned(creator string, blockHeight int32, asset string, amount int32, denom string) *MsgCreateDepositEarned {
	return &MsgCreateDepositEarned{
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgCreateDepositEarned) Route() string {
	return RouterKey
}

func (msg *MsgCreateDepositEarned) Type() string {
	return "CreateDepositEarned"
}

func (msg *MsgCreateDepositEarned) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDepositEarned) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDepositEarned) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDepositEarned{}

func NewMsgUpdateDepositEarned(creator string, id uint64, blockHeight int32, asset string, amount int32, denom string) *MsgUpdateDepositEarned {
	return &MsgUpdateDepositEarned{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgUpdateDepositEarned) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDepositEarned) Type() string {
	return "UpdateDepositEarned"
}

func (msg *MsgUpdateDepositEarned) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDepositEarned) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDepositEarned) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateDepositEarned{}

func NewMsgDeleteDepositEarned(creator string, id uint64) *MsgDeleteDepositEarned {
	return &MsgDeleteDepositEarned{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteDepositEarned) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDepositEarned) Type() string {
	return "DeleteDepositEarned"
}

func (msg *MsgDeleteDepositEarned) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDepositEarned) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDepositEarned) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
