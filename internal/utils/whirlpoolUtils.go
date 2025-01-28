package utils

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/Norbaeocystin/gorca"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

const Q64_RESOLUTION float64 = 18446744073709551616.0

type Pool struct {
	PoolAddress    solana.PublicKey
	TokenASymbol   string
	TokenBSymbol   string
	TokenADecimals uint8
	TokenBDecimals uint8
	FeeRate        uint16
}

func FetchAndUnmarshalPoolData(ctx context.Context, poolAddress solana.PublicKey, rpcClient *rpc.Client) (pool Pool) {
	poolData := GetWhirlpoolData(
		rpcClient,
		poolAddress,
	)
	pool = Pool{
		PoolAddress:    poolAddress,
		TokenASymbol:   GetTokenSymbol(ctx, *poolData.TokenMintA, rpcClient), // TODO: get token mint data before
		TokenBSymbol:   GetTokenSymbol(ctx, *poolData.TokenMintB, rpcClient),
		TokenADecimals: GetTokenDecimals(ctx, *poolData.TokenMintA, rpcClient),
		TokenBDecimals: GetTokenDecimals(ctx, *poolData.TokenMintB, rpcClient),
		FeeRate:        poolData.FeeRate,
	}
	return
}

func CalculatePoolPriceAndFees(sqrtPrice bin.Uint128, poolData Pool) (float64, float64) {
	price := GetPoolCurrentPrice(sqrtPrice, poolData)
	fees := GetSwapFee(price, poolData.FeeRate)
	return price, fees
}

// copied from gorca library, added proper error handling
func GetWhirlpoolData(client *rpc.Client, whirlpoolAddress solana.PublicKey) gorca.WhirlpoolData {
	account, err := client.GetAccountInfoWithOpts(context.TODO(),
		whirlpoolAddress,
		&rpc.GetAccountInfoOpts{
			Encoding:       solana.EncodingBase64,
			Commitment:     rpc.CommitmentFinalized,
			DataSlice:      nil,
			MinContextSlot: nil,
		},
	)
	if err != nil {
		log.Fatalf("failed to get account info: %v", err)
	}
	poolData := DecodeWhirlpoolData(account.Value.Data)
	return poolData
}

func DecodeWhirlpoolData(data *rpc.DataBytesOrJSON) (poolData gorca.WhirlpoolData) {
	dataPos := data.GetBinary()
	borshDec := bin.NewBorshDecoder(dataPos)
	borshDec.Decode(&poolData)
	return
}

func GetPoolCurrentPrice(sqrtPrice bin.Uint128, poolData Pool) float64 {
	return sqrtPriceToPrice(
		sqrtPrice.BigInt(),
		poolData.TokenADecimals,
		poolData.TokenBDecimals,
	)
}

// logic copied from Orca Whirlpool's Core SDK
// at https://github.com/orca-so/whirlpools/blob/6210d2b46ed6741ce60b330bd2cfc4b8e74a55b4/rust-sdk/core/src/math/price.rs#L41
func sqrtPriceToPrice(sqrtPrice *big.Int, decimalsA uint8, decimalsB uint8) float64 {
	power := math.Pow(10.0, float64(decimalsA)-float64(decimalsB))
	sqrtPriceFloat := new(big.Float).SetInt(sqrtPrice)
	sqrtPriceFloat.Quo(sqrtPriceFloat, big.NewFloat(Q64_RESOLUTION))
	price, _ := sqrtPriceFloat.Float64()
	return math.Pow(price, 2.0) * power
}

func GetSwapFee(poolPrice float64, feeRate uint16) float64 {
	return poolPrice * float64(feeRate) / 1_000_000.0
}

//func GetWhirlpoolCurrentPrice(ctx context.Context, whirlpoolAddress solana.PublicKey, rpcClient *rpc.Client) float64 {
//	whirlpoolData := getWhirlpoolData(rpcClient, whirlpoolAddress)
//	vaultABalance := GetTokenAccountBalance(ctx, *whirlpoolData.TokenVaultA, rpcClient)
//	vaultBBalance := GetTokenAccountBalance(ctx, *whirlpoolData.TokenVaultB, rpcClient)
//	return calculateWhirlpoolPrice(vaultABalance, vaultBBalance)
//}

//func calculateWhirlpoolPrice(vaultABalance float64, vaultBBalance float64) (whirlpoolPrice float64) {
//	if vaultABalance > vaultBBalance {
//		whirlpoolPrice = vaultABalance / vaultBBalance
//	} else {
//		whirlpoolPrice = vaultBBalance / vaultABalance
//	}
//	return
//}
