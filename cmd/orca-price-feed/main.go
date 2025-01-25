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
	balance := utils.GetSolBalance(whirlpoolAddress, rpcClient)
	fmt.Println("whirlpool's balance:", balance, "SOL")

	whirlpoolData := gorca.GetWhirlpoolData(rpcClient, whirlpoolAddress)
	fmt.Println("Token Vault A address:", whirlpoolData.TokenVaultA)
	fmt.Println("Token Vault B address:", whirlpoolData.TokenVaultB)
	vaultAAddress := solana.MustPublicKeyFromBase58(whirlpoolData.TokenVaultA.String())
	vaultBAddress := solana.MustPublicKeyFromBase58(whirlpoolData.TokenVaultB.String())
	fmt.Println("Token Vault A balance:", utils.GetSolBalance(vaultAAddress, rpcClient))
	fmt.Println("Token Vault B balance:", utils.GetSolBalance(vaultBAddress, rpcClient))
}
