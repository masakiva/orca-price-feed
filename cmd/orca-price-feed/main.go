package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Norbaeocystin/gorca"
	"github.com/gagliardetto/solana-go"
	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	ctx := context.Background()
	rpcClient := utils.GetRpcClient()
	// wallet := utils.GetWallet("./devnet-wallet.json")

	// Orca's SOL/USDC pool address
	whirlpoolAddress, err := solana.PublicKeyFromBase58("Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE")
	if err != nil {
		log.Fatalf("failed to parse whirlpool address: %v", err)
	}

	balance := utils.GetSolBalance(ctx, whirlpoolAddress, rpcClient)
	fmt.Println("whirlpool's balance:", balance, "SOL")

	whirlpoolData := gorca.GetWhirlpoolData(rpcClient, whirlpoolAddress)

	vaultABalance := utils.GetPrettyTokenAccountBalance(
		ctx,
		*whirlpoolData.TokenVaultA,
		*whirlpoolData.TokenMintA,
		rpcClient,
	)
	vaultBBalance := utils.GetPrettyTokenAccountBalance(
		ctx,
		*whirlpoolData.TokenVaultB,
		*whirlpoolData.TokenMintB,
		rpcClient,
	)
	fmt.Println("Token Vault A balance:", vaultABalance)
	fmt.Println("Token Vault B balance:", vaultBBalance)
}
