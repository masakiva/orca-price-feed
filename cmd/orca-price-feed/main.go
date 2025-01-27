package main

import (
	"context"
	"fmt"
	"log"

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
	whirlpoolData := utils.GetWhirlpoolData(rpcClient, whirlpoolAddress)

	whirlpoolPrice := utils.GetWhirlpoolCurrentPrice(ctx, whirlpoolData, rpcClient)
	fmt.Printf("Whirlpool price: %f\n", whirlpoolPrice)

	swapFee := utils.GetSwapFee(whirlpoolPrice, whirlpoolData.FeeRate)
	fmt.Printf("Price without fees: %f\n", whirlpoolPrice-swapFee)
}
