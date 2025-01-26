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

// GetPrettyTokenAccountBalance returns the formatted balance of a token account,
// with its symbol.
func GetPrettyTokenAccountBalance(ctx context.Context, tokenAccountPubKey solana.PublicKey, tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) string {
	rawAmount, tokenDecimals := fetchTokenAccountBalance(ctx, tokenAccountPubKey, rpcClient)
	amountFloat, err := strconv.ParseFloat(rawAmount, 64)
	if err != nil {
		log.Fatalf("failed to parse token account balance: %v", err)
	}
	amountFloat /= math.Pow(10, float64(tokenDecimals))

	tokenSymbol := fetchTokenSymbol(ctx, tokenMintPubKey, rpcClient)

	prettyBalance := fmt.Sprintf("%.6f %s", amountFloat, tokenSymbol)
	return prettyBalance
}

// fetchTokenAccountBalance retrieves the balance and decimals of a token account.
func fetchTokenAccountBalance(ctx context.Context, tokenAccountPubKey solana.PublicKey, rpcClient *rpc.Client) (string, uint8) {
	balance, err := rpcClient.GetTokenAccountBalance(
		ctx,
		tokenAccountPubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		log.Fatalf("failed to get token account balance: %v", err)
	}
	return balance.Value.Amount, balance.Value.Decimals
}

func fetchTokenSymbol(ctx context.Context, tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) string {
	metadata, err := fetchTokenMetadata(ctx, tokenMintPubKey, rpcClient)
	if err != nil {
		log.Fatalf("failed to get token symbol: %v", err)
	}
	return metadata.Data.Symbol
}

func fetchTokenMetadata(ctx context.Context, tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) (*token_metadata.Metadata, error) {
	metadataAccountPubKey, _, err := solana.FindTokenMetadataAddress(tokenMintPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to find token metadata address: %w", err)
	}

	var metadata token_metadata.Metadata
	err = rpcClient.GetAccountDataBorshInto(
		ctx,
		metadataAccountPubKey,
		&metadata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata account data: %w", err)
	}
	return &metadata, nil
}
