package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateApr{}, "ibb/CreateApr", nil)
	cdc.RegisterConcrete(&MsgUpdateApr{}, "ibb/UpdateApr", nil)
	cdc.RegisterConcrete(&MsgDeleteApr{}, "ibb/DeleteApr", nil)

	cdc.RegisterConcrete(&MsgCreateRepay{}, "ibb/CreateRepay", nil)
	cdc.RegisterConcrete(&MsgUpdateRepay{}, "ibb/UpdateRepay", nil)
	cdc.RegisterConcrete(&MsgDeleteRepay{}, "ibb/DeleteRepay", nil)

	cdc.RegisterConcrete(&MsgCreateWithdraw{}, "ibb/CreateWithdraw", nil)
	cdc.RegisterConcrete(&MsgUpdateWithdraw{}, "ibb/UpdateWithdraw", nil)
	cdc.RegisterConcrete(&MsgDeleteWithdraw{}, "ibb/DeleteWithdraw", nil)

	cdc.RegisterConcrete(&MsgCreateUser{}, "ibb/CreateUser", nil)
	cdc.RegisterConcrete(&MsgUpdateUser{}, "ibb/UpdateUser", nil)
	cdc.RegisterConcrete(&MsgDeleteUser{}, "ibb/DeleteUser", nil)

	cdc.RegisterConcrete(&MsgCreateBorrow{}, "ibb/CreateBorrow", nil)
	cdc.RegisterConcrete(&MsgUpdateBorrow{}, "ibb/UpdateBorrow", nil)
	cdc.RegisterConcrete(&MsgDeleteBorrow{}, "ibb/DeleteBorrow", nil)

	cdc.RegisterConcrete(&MsgCreateDeposit{}, "ibb/CreateDeposit", nil)
	cdc.RegisterConcrete(&MsgUpdateDeposit{}, "ibb/UpdateDeposit", nil)
	cdc.RegisterConcrete(&MsgDeleteDeposit{}, "ibb/DeleteDeposit", nil)

	cdc.RegisterConcrete(&MsgCreatePool{}, "ibb/CreatePool", nil)
	cdc.RegisterConcrete(&MsgUpdatePool{}, "ibb/UpdatePool", nil)
	cdc.RegisterConcrete(&MsgDeletePool{}, "ibb/DeletePool", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateApr{},
		&MsgUpdateApr{},
		&MsgDeleteApr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRepay{},
		&MsgUpdateRepay{},
		&MsgDeleteRepay{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWithdraw{},
		&MsgUpdateWithdraw{},
		&MsgDeleteWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUser{},
		&MsgUpdateUser{},
		&MsgDeleteUser{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBorrow{},
		&MsgUpdateBorrow{},
		&MsgDeleteBorrow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDeposit{},
		&MsgUpdateDeposit{},
		&MsgDeleteDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
		&MsgUpdatePool{},
		&MsgDeletePool{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
