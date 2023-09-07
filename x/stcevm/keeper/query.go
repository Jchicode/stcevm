package keeper

import (
	"github.com/Jchicode/stcevm/x/stcevm/types"
)

var _ types.QueryServer = Keeper{}
