package achilles_test

import (
	"testing"

	keepertest "github.com/olimdzhon/achilles/testutil/keeper"
	"github.com/olimdzhon/achilles/testutil/nullify"
	achilles "github.com/olimdzhon/achilles/x/achilles/module"
	"github.com/olimdzhon/achilles/x/achilles/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AchillesKeeper(t)
	achilles.InitGenesis(ctx, k, genesisState)
	got := achilles.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
