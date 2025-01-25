package main

import (
	"fmt"

	"github.com/Norbaeocystin/gorca"
	"github.com/gagliardetto/solana-go"
	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	rpcClient := utils.GetRpcClient()
	// wallet := utils.GetWallet("./devnet-wallet.json")

	// Orca's SOL/USDC pool address
	whirlpoolAddress := solana.MustPublicKeyFromBase58("Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE") // TODO: handle errors properly using PublicKeyFromBase58() instead
	balance := utils.GetBalanceInSol(whirlpoolAddress, rpcClient)
	fmt.Println("whirlpool's balance:", balance, "SOL")

	whirlpoolData := gorca.GetWhirlpoolData(rpcClient, whirlpoolAddress)
	fmt.Println(whirlpoolData.TokenMintA)
	fmt.Println(whirlpoolData.TokenMintB)
}
