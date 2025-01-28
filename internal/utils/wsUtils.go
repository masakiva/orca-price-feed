package utils

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func GetWsClient(ctx context.Context) *ws.Client {
	endpointUrl := os.Getenv("WS_URL")
	if endpointUrl != "" {
		_, err := url.ParseRequestURI(endpointUrl)
		if err != nil {
			log.Fatalf("Failed to parse WS_URL: %v", err)
		}
	} else {
		fmt.Println("WS_URL not defined, defaulting to Solana mainnet beta")
		endpointUrl = rpc.MainNetBeta_WS
	}
	wsClient, err := ws.Connect(ctx, endpointUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Solana WebSocket: %v", err)
	}
	return wsClient
}

func SubscribeToWsAccount(wsClient *ws.Client, pool Pool) *ws.AccountSubscription {
	sub, err := wsClient.AccountSubscribe(
		pool.PoolAddress,
		rpc.CommitmentFinalized, // You can use "Processed", "Confirmed", or "Finalized"
		// TODO:pass commitment level by env var?
	)
	if err != nil {
		log.Fatalf("Failed to subscribe to account: %v", err)
	}
	return sub
}
