package main

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	rpcClient := utils.GetRpcClient()
	// wallet := utils.GetWallet("./devnet-wallet.json")

	// Orca's SOL/USDC pool address
	poolPubKey, err := solana.PublicKeyFromBase58("Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE")
	if err != nil {
		panic(err)
	}
	balance := utils.GetBalanceInSol(poolPubKey, rpcClient)
	fmt.Println("pool's balance:", balance, "SOL")
}
