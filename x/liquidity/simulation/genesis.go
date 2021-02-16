package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/liquidity/x/liquidity/types"
)

// Simulation parameter constants
const (
	LiquidityPoolTypes       = "liquidity_pool_types"
	MinInitDepositToPool     = "min_init_deposit_to_pool"
	InitPoolCoinMintAmount   = "init_pool_coin_mint_amount"
	LiquidityPoolCreationFee = "liquidity_pool_creation_fee"
	SwapFeeRate              = "swap_fee_rate"
	WithdrawFeeRate          = "withdraw_fee_rate"
	MaxOrderAmountRatio      = "max_order_amount_ratio"
	UnitBatchSize            = "unit_batch_size"
	LiquidityPool            = "liquidity_pool"
	LiquidityPoolMetadata    = "liquidity_pool_metadata"
	LiquidityPoolBatch       = "liquidity_pool_batch"
	BatchPoolDepositMsgs     = "batch_pool_deposit_msgs"
	BatchPoolWithdrawMsgs    = "batch_pool_withdraw_msgs"
	BatchPoolSwapMsgs        = "batch_pool_swap_msgs"
)

// GenLiquidityPoolTypes randomized LiquidityPoolTypes
func GenLiquidityPoolTypes(r *rand.Rand) []types.LiquidityPoolType {
	liquidityPoolTypes := []types.LiquidityPoolType{}

	liquidityPoolType := types.LiquidityPoolType{
		PoolTypeIndex:     1,
		Name:              "simulation",
		MinReserveCoinNum: 2,
		MaxReserveCoinNum: 2,
		Description:       "simulation",
	}

	liquidityPoolTypes = append(liquidityPoolTypes, liquidityPoolType)

	return liquidityPoolTypes
}

// GenMinInitDepositToPool randomized MinInitDepositToPool
// [Qs] min_init_deposit_to_pool param is "1000000". Should MinInitDepositToPool be 10^6?
func GenMinInitDepositToPool(r *rand.Rand) sdk.Int {
	return sdk.NewInt(1e6)
}

// GenInitPoolCoinMintAmount randomized InitPoolCoinMintAmount
// [Qs] init_pool_coin_mint_amount param is "1000000". Should InitPoolCoinMintAmount be 10^6?
func GenInitPoolCoinMintAmount(r *rand.Rand) sdk.Int {
	return sdk.NewInt(1e6)
}

// GenLiquidityPoolCreationFee randomized LiquidityPoolCreationFee 100000000
// [Qs] liquidity_pool_creation_fee amount is "100000000".
func GenLiquidityPoolCreationFee(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simulation.RandIntBetween(r, 1e6, 1e7))))
}

// GenSwapFeeRate randomized SwapFeeRate
// [Qs] swap_fee_rate param is 0.003000000000000000
func GenSwapFeeRate(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(3, 18)
}

// GenWithdrawFeeRate randomized WithdrawFeeRate
// [Qs] withdraw_fee_rate param is 0.003000000000000000
func GenWithdrawFeeRate(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(3, 18)
}

// GenMaxOrderAmountRatio randomized MaxOrderAmountRatio
// [Qs] max_order_amount_ratio param is 0.100000000000000000
func GenMaxOrderAmountRatio(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(1, 18)
}

// GenUnitBatchSize randomized UnitBatchSize
// [Qs] unit batch size can be randomized?
func GenUnitBatchSize(r *rand.Rand) uint32 {
	return r.Uint32()
}

// GenLiquidityPool randomized LiquidityPool
func GenLiquidityPool(r *rand.Rand, simState *module.SimulationState) types.LiquidityPool {
	liquidityPool := types.LiquidityPool{
		PoolId:                r.Uint64(),
		PoolTypeIndex:         uint32(1),
		ReserveCoinDenoms:     []string{sdk.DefaultBondDenom, "uatom"},
		ReserveAccountAddress: simState.Accounts[0].Address.String(),
		PoolCoinDenom:         sdk.DefaultBondDenom,
	}

	return liquidityPool
}

// GenLiquidityPoolMetadata randomized LiquidityPoolMetadata
func GenLiquidityPoolMetadata(r *rand.Rand) types.LiquidityPoolMetadata {
	return types.LiquidityPoolMetadata{}
}

// GenLiquidityPoolBatch randomized LiquidityPoolBatch
func GenLiquidityPoolBatch(r *rand.Rand) types.LiquidityPoolBatch {
	return types.LiquidityPoolBatch{}
}

// GenBatchPoolDepositMsgs randomized BatchPoolDepositMsg
func GenBatchPoolDepositMsgs(r *rand.Rand) []types.BatchPoolDepositMsg {
	return []types.BatchPoolDepositMsg{}
}

// GenBatchPoolWithdrawMsgs randomized BatchPoolWithdrawMsg
func GenBatchPoolWithdrawMsgs(r *rand.Rand) []types.BatchPoolWithdrawMsg {
	return []types.BatchPoolWithdrawMsg{}
}

// GenBatchPoolSwapMsgs randomized BatchPoolSwapMsg
func GenBatchPoolSwapMsgs(r *rand.Rand) []types.BatchPoolSwapMsg {
	return []types.BatchPoolSwapMsg{}
}

// RandomizedGenState generates a random GenesisState for distribution
func RandomizedGenState(simState *module.SimulationState) {
	var liquidityPoolTypes []types.LiquidityPoolType
	simState.AppParams.GetOrGenerate(
		simState.Cdc, LiquidityPoolTypes, &liquidityPoolTypes, simState.Rand,
		func(r *rand.Rand) { liquidityPoolTypes = GenLiquidityPoolTypes(r) },
	)

	var minInitDepositToPool sdk.Int
	simState.AppParams.GetOrGenerate(
		simState.Cdc, MinInitDepositToPool, &minInitDepositToPool, simState.Rand,
		func(r *rand.Rand) { minInitDepositToPool = GenMinInitDepositToPool(r) },
	)

	var initPoolCoinMintAmount sdk.Int
	simState.AppParams.GetOrGenerate(
		simState.Cdc, InitPoolCoinMintAmount, &initPoolCoinMintAmount, simState.Rand,
		func(r *rand.Rand) { initPoolCoinMintAmount = GenInitPoolCoinMintAmount(r) },
	)

	var liquidityPoolCreationFee sdk.Coins
	simState.AppParams.GetOrGenerate(
		simState.Cdc, LiquidityPoolCreationFee, &liquidityPoolCreationFee, simState.Rand,
		func(r *rand.Rand) { liquidityPoolCreationFee = GenLiquidityPoolCreationFee(r) },
	)

	var swapFeeRate sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, SwapFeeRate, &swapFeeRate, simState.Rand,
		func(r *rand.Rand) { swapFeeRate = GenSwapFeeRate(r) },
	)

	var withdrawFeeRate sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, WithdrawFeeRate, &withdrawFeeRate, simState.Rand,
		func(r *rand.Rand) { withdrawFeeRate = GenWithdrawFeeRate(r) },
	)

	var maxOrderAmountRatio sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, MaxOrderAmountRatio, &maxOrderAmountRatio, simState.Rand,
		func(r *rand.Rand) { maxOrderAmountRatio = GenMaxOrderAmountRatio(r) },
	)

	var unitBatchSize uint32
	simState.AppParams.GetOrGenerate(
		simState.Cdc, UnitBatchSize, &unitBatchSize, simState.Rand,
		func(r *rand.Rand) { unitBatchSize = GenUnitBatchSize(r) },
	)

	var liquidityPool types.LiquidityPool
	simState.AppParams.GetOrGenerate(
		simState.Cdc, LiquidityPool, &liquidityPool, simState.Rand,
		func(r *rand.Rand) { liquidityPool = GenLiquidityPool(r, simState) },
	)

	var liquidityPoolMetadata types.LiquidityPoolMetadata
	simState.AppParams.GetOrGenerate(
		simState.Cdc, LiquidityPoolMetadata, &liquidityPoolMetadata, simState.Rand,
		func(r *rand.Rand) { liquidityPoolMetadata = GenLiquidityPoolMetadata(r) },
	)

	var liquidityPoolBatch types.LiquidityPoolBatch
	simState.AppParams.GetOrGenerate(
		simState.Cdc, LiquidityPoolBatch, &liquidityPoolBatch, simState.Rand,
		func(r *rand.Rand) { liquidityPoolBatch = GenLiquidityPoolBatch(r) },
	)

	var batchPoolDepositMsgs []types.BatchPoolDepositMsg
	simState.AppParams.GetOrGenerate(
		simState.Cdc, BatchPoolDepositMsgs, &batchPoolDepositMsgs, simState.Rand,
		func(r *rand.Rand) { batchPoolDepositMsgs = GenBatchPoolDepositMsgs(r) },
	)

	var batchPoolWithdrawMsgs []types.BatchPoolWithdrawMsg
	simState.AppParams.GetOrGenerate(
		simState.Cdc, BatchPoolWithdrawMsgs, &batchPoolWithdrawMsgs, simState.Rand,
		func(r *rand.Rand) { batchPoolWithdrawMsgs = GenBatchPoolWithdrawMsgs(r) },
	)

	var batchPoolSwapMsgs []types.BatchPoolSwapMsg
	simState.AppParams.GetOrGenerate(
		simState.Cdc, BatchPoolSwapMsgs, &batchPoolSwapMsgs, simState.Rand,
		func(r *rand.Rand) { batchPoolSwapMsgs = GenBatchPoolSwapMsgs(r) },
	)

	var liquidityPoolRecords []types.LiquidityPoolRecord
	liquidityPoolRecord := types.LiquidityPoolRecord{
		LiquidityPool:         liquidityPool,
		LiquidityPoolMetadata: liquidityPoolMetadata,
		LiquidityPoolBatch:    liquidityPoolBatch,
		BatchPoolDepositMsgs:  batchPoolDepositMsgs,
		BatchPoolWithdrawMsgs: batchPoolWithdrawMsgs,
		BatchPoolSwapMsgs:     batchPoolSwapMsgs,
	}
	liquidityPoolRecords = append(liquidityPoolRecords, liquidityPoolRecord)

	liquidityGenesis := types.GenesisState{
		Params: types.Params{
			LiquidityPoolTypes:       liquidityPoolTypes,
			MinInitDepositToPool:     minInitDepositToPool,
			InitPoolCoinMintAmount:   initPoolCoinMintAmount,
			LiquidityPoolCreationFee: liquidityPoolCreationFee,
			SwapFeeRate:              swapFeeRate,
			WithdrawFeeRate:          withdrawFeeRate,
			MaxOrderAmountRatio:      maxOrderAmountRatio,
			UnitBatchSize:            unitBatchSize,
		},
		LiquidityPoolRecords: liquidityPoolRecords,
	}

	bz, err := json.MarshalIndent(&liquidityGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated distribution parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&liquidityGenesis)
}
