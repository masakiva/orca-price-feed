package main

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func getWallet(pathToPrivKey string) solana.Wallet {
	var wallet solana.Wallet
	var err error

	wallet.PrivateKey, err = solana.PrivateKeyFromSolanaKeygenFile(pathToPrivKey)
	if err != nil {
		panic(err)
	}
	return wallet
}

func getBalanceInSol(pubKey solana.PublicKey, rpcClient *rpc.Client) uint64 {
	balance, err := rpcClient.GetBalance(
		context.TODO(),
		pubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	return balance.Value / solana.LAMPORTS_PER_SOL
}
