package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/liquidity/x/liquidity/simulation"
	"github.com/tendermint/liquidity/x/liquidity/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var liquidityGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &liquidityGenesis)

	dec1, _ := sdk.NewIntFromString("1000000")
	dec2, _ := sdk.NewIntFromString("1000000")
	dec3, _ := sdk.NewDecFromStr("0.003000000000000000")
	dec4, _ := sdk.NewDecFromStr("0.003000000000000000")

	require.Equal(t, dec1, liquidityGenesis.Params.LiquidityPoolTypes)
	require.Equal(t, dec2, liquidityGenesis.Params.MinInitDepositToPool)
	require.Equal(t, dec3, liquidityGenesis.Params.SwapFeeRate)
	require.Equal(t, dec4, liquidityGenesis.Params.WithdrawFeeRate)
}
