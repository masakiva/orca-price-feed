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

func GetWhirlpoolCurrentPrice(ctx context.Context, whirlpoolAddress solana.PublicKey, rpcClient *rpc.Client) float64 {
	whirlpoolData := getWhirlpoolData(rpcClient, whirlpoolAddress)
	tokenADecimals := GetTokenDecimals(ctx, *whirlpoolData.TokenMintA, rpcClient)
	tokenBDecimals := GetTokenDecimals(ctx, *whirlpoolData.TokenMintB, rpcClient)
	return sqrtPriceToPrice(
		whirlpoolData.SqrtPrice.BigInt(),
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

// copied from gorca library, added proper error handling
func getWhirlpoolData(client *rpc.Client, whirlpoolAddress solana.PublicKey) gorca.WhirlpoolData {
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
