package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/fadeev/foo-bar1/testutil/keeper"
	"github.com/fadeev/foo-bar1/x/foobar1/keeper"
	"github.com/fadeev/foo-bar1/x/foobar1/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Foobar1Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
