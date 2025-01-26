package utils

import (
	"context"
	"log"

	"github.com/Norbaeocystin/gorca"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetWhirlpoolCurrentPrice(ctx context.Context, whirlpoolAddress solana.PublicKey, rpcClient *rpc.Client) float64 {
	whirlpoolData := getWhirlpoolData(rpcClient, whirlpoolAddress)
	vaultABalance := GetTokenAccountBalance(ctx, *whirlpoolData.TokenVaultA, rpcClient)
	vaultBBalance := GetTokenAccountBalance(ctx, *whirlpoolData.TokenVaultB, rpcClient)
	return calculateWhirlpoolPrice(vaultABalance, vaultBBalance)
}

// copied from gorca library, with proper error handling
func getWhirlpoolData(client *rpc.Client, whirlpoolAddress solana.PublicKey) gorca.WhirlpoolData {
	account, err := client.GetAccountInfoWithOpts(context.TODO(),
		whirlpoolAddress,
		&rpc.GetAccountInfoOpts{
			Encoding:       solana.EncodingBase64,
			Commitment:     rpc.CommitmentFinalized,
			DataSlice:      nil,
			MinContextSlot: nil,
		},
	)
	if err != nil {
		log.Fatalf("failed to get account info: %v", err)
	}
	var wpData gorca.WhirlpoolData
	dataPos := account.GetBinary()
	borshDec := bin.NewBorshDecoder(dataPos)
	borshDec.Decode(&wpData)
	// log.Println(wpData)
	return wpData
}

func calculateWhirlpoolPrice(vaultABalance float64, vaultBBalance float64) (whirlpoolPrice float64) {
	if vaultABalance > vaultBBalance {
		whirlpoolPrice = vaultABalance / vaultBBalance
	} else {
		whirlpoolPrice = vaultBBalance / vaultABalance
	}
	return
}
