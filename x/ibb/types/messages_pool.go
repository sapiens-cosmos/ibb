package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, asset string, denom string, collatoralFactor int32, depositBalance int32, borrowBalance int32, users []*User, aprs []*Apr) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:          creator,
		Asset:            asset,
		Denom:            denom,
		CollatoralFactor: collatoralFactor,
		DepositBalance:   depositBalance,
		BorrowBalance:    borrowBalance,
		Users:            users,
		Aprs:             aprs,
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return "CreatePool"
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePool{}

func NewMsgUpdatePool(creator string, id uint64, asset string, denom string, collatoralFactor int32, borrowBalance int32, depositBalance int32, users []*User, aprs []*Apr) *MsgUpdatePool {
	return &MsgUpdatePool{
		Id:               id,
		Creator:          creator,
		Asset:            asset,
		Denom:            denom,
		CollatoralFactor: collatoralFactor,
		BorrowBalance:    borrowBalance,
		DepositBalance:   depositBalance,
		Users:            users,
		Aprs:             aprs,
	}
}

func (msg *MsgUpdatePool) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePool) Type() string {
	return "UpdatePool"
}

func (msg *MsgUpdatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePool{}

func NewMsgDeletePool(creator string, id uint64) *MsgDeletePool {
	return &MsgDeletePool{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePool) Route() string {
	return RouterKey
}

func (msg *MsgDeletePool) Type() string {
	return "DeletePool"
}

func (msg *MsgDeletePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
