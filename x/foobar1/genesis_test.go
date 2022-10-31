package foobar1_test

import (
	"testing"

	keepertest "github.com/fadeev/foo-bar1/testutil/keeper"
	"github.com/fadeev/foo-bar1/testutil/nullify"
	"github.com/fadeev/foo-bar1/x/foobar1"
	"github.com/fadeev/foo-bar1/x/foobar1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Foobar1Keeper(t)
	foobar1.InitGenesis(ctx, *k, genesisState)
	got := foobar1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
