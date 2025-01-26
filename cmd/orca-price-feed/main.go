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
	// balance := utils.GetSolBalance(ctx, wallet.PublicKey(), rpcClient)

	// Orca's SOL/USDC pool address
	// TODO: get addresses from user input
	whirlpoolAddress, err := solana.PublicKeyFromBase58("Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE")
	if err != nil {
		log.Fatalf("failed to parse whirlpool address: %v", err)
	}

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
