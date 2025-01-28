package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Norbaeocystin/gorca"
	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	addressFilePath := utils.ParseCmdLineArgs()
	poolAddresses := utils.ReadAddressesFromFile(addressFilePath)

	wsClient, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
	if err != nil {
		log.Fatalf("Failed to connect to Solana WebSocket: %v", err)
	}
	defer wsClient.Close()

	// Define the account to monitor
	accountPubKey := solana.MustPublicKeyFromBase58(poolAddresses[0])

	// Subscribe to account changes
	sub, err := wsClient.AccountSubscribe(
		accountPubKey,
		rpc.CommitmentFinalized, // You can use "Processed", "Confirmed", or "Finalized"
	)
	if err != nil {
		log.Fatalf("Failed to subscribe to account: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Printf("Subscribed to account: %s\n", accountPubKey)

	// Listen for events
	for {
		notification, err := sub.Recv()
		if err != nil {
			log.Fatalf("Error receiving notification: %v", err)
		}

		// Handle account data changes
		fmt.Printf("Account data changed for %s:\n", accountPubKey)
		fmt.Printf("Slot: %d\n", notification.Context.Slot)
		fmt.Printf("Account Data (raw): %x\n", notification.Value.Data.GetBinary())

		// Decode the data here to extract the price or relevant fields
		// You can implement a decoding function based on the account structure

		accountData := notification.Value.Data
		var wpData gorca.WhirlpoolData
		dataPos := accountData.GetBinary()
		borshDec := bin.NewBorshDecoder(dataPos)
		borshDec.Decode(&wpData)
		spew.Dump(wpData)
	}
}
