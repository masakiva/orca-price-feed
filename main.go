package main

import (
	"fmt"
)

func main() {
	rpcClient := getRpcClient()
	wallet := getWallet("./devnet-wallet.json")

	walletBalance := getBalanceInSol(wallet.PublicKey(), rpcClient)
	fmt.Println("accountFrom's balance:", walletBalance, "SOL")
}
