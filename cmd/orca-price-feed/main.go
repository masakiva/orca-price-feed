package main

import (
	"fmt"

	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	rpcClient := utils.GetRpcClient()
	wallet := utils.GetWallet("./devnet-wallet.json")

	balance := utils.GetBalanceInSol(wallet.PublicKey(), rpcClient)
	fmt.Println("wallet's balance:", balance, "SOL")
}
