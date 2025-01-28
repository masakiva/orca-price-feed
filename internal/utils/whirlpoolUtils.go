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
	PoolAddress  string
	TokenASymbol string
	TokenBSymbol string
	Price        float64
	SwapFee      float64
}

func FetchAndUnmarshalPoolData(ctx context.Context, poolAddresses []string, rpcClient *rpc.Client) (pools []Pool) {
	for _, addressStr := range poolAddresses {
		poolData := GetWhirlpoolData(
			rpcClient,
			GetSolanaAddressFromString(addressStr),
		)
		poolPrice := GetPoolCurrentPrice(ctx, poolData, rpcClient)
		pool := Pool{
			PoolAddress:  addressStr,
			TokenASymbol: GetTokenSymbol(ctx, *poolData.TokenMintA, rpcClient),
			TokenBSymbol: GetTokenSymbol(ctx, *poolData.TokenMintB, rpcClient),
			Price:        poolPrice,
			SwapFee:      GetSwapFee(poolPrice, poolData.FeeRate),
		}
		pools = append(pools, pool)
	}
	return
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
	var wpData gorca.WhirlpoolData
	dataPos := account.GetBinary()
	borshDec := bin.NewBorshDecoder(dataPos)
	borshDec.Decode(&wpData)
	// log.Println(wpData)
	return wpData
}

func GetPoolCurrentPrice(ctx context.Context, poolData gorca.WhirlpoolData, rpcClient *rpc.Client) float64 {
	tokenADecimals := GetTokenDecimals(ctx, *poolData.TokenMintA, rpcClient)
	tokenBDecimals := GetTokenDecimals(ctx, *poolData.TokenMintB, rpcClient)
	return sqrtPriceToPrice(
		poolData.SqrtPrice.BigInt(),
		tokenADecimals,
		tokenBDecimals,
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
