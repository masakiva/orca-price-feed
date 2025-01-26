package utils

import (
	"context"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetWallet(pathToPrivKey string) solana.Wallet {
	var wallet solana.Wallet
	var err error

	wallet.PrivateKey, err = solana.PrivateKeyFromSolanaKeygenFile(pathToPrivKey)
	if err != nil {
		log.Fatalf("failed to get private key from file: %v", err)
	}
	return wallet
}

func GetSolBalance(ctx context.Context, pubKey solana.PublicKey, rpcClient *rpc.Client) uint64 {
	balance, err := rpcClient.GetBalance(
		ctx,
		pubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		log.Fatalf("failed to get SOL balance: %v", err)
	}
	return balance.Value / solana.LAMPORTS_PER_SOL
}
