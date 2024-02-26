package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "switcheo/testutil/keeper"
	"switcheo/x/items/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ItemsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
