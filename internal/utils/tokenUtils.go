package utils

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetPrettyTokenAccountBalance(tokenAccountPubKey solana.PublicKey, tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) string {
	amount, decimals := getTokenAccountBalance(tokenAccountPubKey, rpcClient)
	// Convert to a float for human-readable balance
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Fatalf("Failed to parse amount: %v", err)
	}
	amountFloat /= math.Pow(10, float64(decimals))
	symbol := getTokenSymbol(tokenMintPubKey, rpcClient)
	prettyStrAmount := fmt.Sprintf("%f %s", amountFloat, symbol)

	return prettyStrAmount
}

func getTokenAccountBalance(tokenAccountPubKey solana.PublicKey, rpcClient *rpc.Client) (string, uint8) {
	balance, err := rpcClient.GetTokenAccountBalance(
		context.TODO(),
		tokenAccountPubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	return balance.Value.Amount, balance.Value.Decimals
}

func getTokenSymbol(tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) string {
	metadata := getTokenMetadata(tokenMintPubKey, rpcClient)
	return metadata.Data.Symbol
}

func getTokenMetadata(tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) *token_metadata.Metadata {
	metadataAccount, _, err := solana.FindTokenMetadataAddress(tokenMintPubKey)
	if err != nil {
		panic(err)
	}

	var metadata token_metadata.Metadata
	rpcClient.GetAccountDataBorshInto(
		context.TODO(),
		metadataAccount,
		&metadata,
	)
	return &metadata
}
