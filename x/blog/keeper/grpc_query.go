package keeper

import (
	"github.com/cgaribay/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
