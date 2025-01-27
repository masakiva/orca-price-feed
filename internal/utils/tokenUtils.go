package utils

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

// GetTokenAccountBalance retrieves and returns the balance of a specified
// token account in its respective unit.
func GetTokenAccountBalance(ctx context.Context, tokenAccountPubKey solana.PublicKey, rpcClient *rpc.Client) (balance float64) {
	balanceInfo, err := rpcClient.GetTokenAccountBalance(
		ctx,
		tokenAccountPubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		log.Fatalf("failed to get token account balance: %v", err)
	}
	balance, err = strconv.ParseFloat(balanceInfo.Value.Amount, 64)
	if err != nil {
		log.Fatalf("failed to parse token account balance: %v", err)
	}
	balance /= math.Pow(10, float64(balanceInfo.Value.Decimals))
	return
}

func GetTokenDecimals(ctx context.Context, tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) uint8 {
	var mint token.Mint
	err := rpcClient.GetAccountDataInto(
		context.TODO(),
		tokenMintPubKey,
		&mint,
	)
	if err != nil {
		log.Fatalf("failed to get mint account data: %v", err)
	}
	return mint.Decimals
}

func GetTokenSymbol(ctx context.Context, tokenMintPubKey solana.PublicKey, rpcClient *rpc.Client) string {
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
