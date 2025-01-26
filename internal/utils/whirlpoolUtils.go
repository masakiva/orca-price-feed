package utils

import (
	"context"

	"github.com/Norbaeocystin/gorca"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetWhirlpoolCurrentPrice(ctx context.Context, whirlpoolAddress solana.PublicKey, rpcClient *rpc.Client) float64 {
	whirlpoolData := gorca.GetWhirlpoolData(rpcClient, whirlpoolAddress)
	vaultABalance := GetTokenAccountBalance(ctx, *whirlpoolData.TokenVaultA, rpcClient)
	vaultBBalance := GetTokenAccountBalance(ctx, *whirlpoolData.TokenVaultB, rpcClient)
	return calculateWhirlpoolPrice(vaultABalance, vaultBBalance)
}

func calculateWhirlpoolPrice(vaultABalance float64, vaultBBalance float64) (whirlpoolPrice float64) {
	if vaultABalance > vaultBBalance {
		whirlpoolPrice = vaultABalance / vaultBBalance
	} else {
		whirlpoolPrice = vaultBBalance / vaultABalance
	}
	return
}
