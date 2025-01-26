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
	vaultABalance := utils.GetPrettyTokenAccountBalance(
		*whirlpoolData.TokenVaultA,
		*whirlpoolData.TokenMintA,
		rpcClient,
	)
	vaultBBalance := utils.GetPrettyTokenAccountBalance(
		*whirlpoolData.TokenVaultB,
		*whirlpoolData.TokenMintB,
		rpcClient,
	)
	fmt.Println("Token Vault A balance:", vaultABalance)
	fmt.Println("Token Vault B balance:", vaultBBalance)
}
