package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/Norbaeocystin/gorca"
	"github.com/gagliardetto/solana-go"
	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	ctx := context.Background()
	rpcClient := utils.GetRpcClient()
	// wallet := utils.GetWallet("./devnet-wallet.json")
	// balance := utils.GetSolBalance(ctx, wallet.PublicKey(), rpcClient)

	// Orca's SOL/USDC pool address
	// TODO: get addresses from user input
	whirlpoolAddress, err := solana.PublicKeyFromBase58("Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE")
	if err != nil {
		log.Fatalf("failed to parse whirlpool address: %v", err)
	}

	whirlpoolPrice := utils.GetWhirlpoolCurrentPrice(ctx, whirlpoolAddress, rpcClient)
	fmt.Printf("Whirlpool price: %f\n", whirlpoolPrice)

	// get price from whirlpool's sqrtPrice, for reference.
	whirlpoolData := gorca.GetWhirlpoolData(rpcClient, whirlpoolAddress)
	rawSqrtPrice := whirlpoolData.SqrtPrice.BigInt()
	twoPow64 := new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)
	sqrtPrice := new(big.Float).Quo(
		new(big.Float).SetInt(rawSqrtPrice),
		new(big.Float).SetInt(twoPow64),
	)
	price := new(big.Float).Mul(sqrtPrice, sqrtPrice)
	fmt.Printf("Price computed from whirlpool's sqrtPrice: %s\n", price.Text('f', 8))

	tickSpacing := whirlpoolData.TickSpacing
	fmt.Printf("Tick spacing: %d\n", tickSpacing)
	feeRate := whirlpoolData.FeeRate
	fmt.Printf("Fee rate: %v\n", feeRate)
}
