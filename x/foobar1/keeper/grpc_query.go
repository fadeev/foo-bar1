package keeper

import (
	"github.com/fadeev/foo-bar1/x/foobar1/types"
)

var _ types.QueryServer = Keeper{}
