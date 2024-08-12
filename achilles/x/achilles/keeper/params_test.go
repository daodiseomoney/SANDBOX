package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/olimdzhon/achilles/testutil/keeper"
	"github.com/olimdzhon/achilles/x/achilles/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AchillesKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
