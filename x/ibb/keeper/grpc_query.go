package keeper

import (
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

var _ types.QueryServer = Keeper{}
