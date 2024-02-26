package items_test

import (
	"testing"

	keepertest "switcheo/testutil/keeper"
	"switcheo/testutil/nullify"
	items "switcheo/x/items/module"
	"switcheo/x/items/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ItemList: []types.Item{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ItemCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ItemsKeeper(t)
	items.InitGenesis(ctx, k, genesisState)
	got := items.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ItemList, got.ItemList)
	require.Equal(t, genesisState.ItemCount, got.ItemCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
