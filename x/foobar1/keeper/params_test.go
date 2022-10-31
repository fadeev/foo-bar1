package keeper_test

import (
	"testing"

	testkeeper "github.com/fadeev/foo-bar1/testutil/keeper"
	"github.com/fadeev/foo-bar1/x/foobar1/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Foobar1Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
