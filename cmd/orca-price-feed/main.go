package main

import (
	"context"
	"log"

	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	ctx := context.Background()

	addressFilePath := utils.ParseCmdLineArgs()
	poolAddresses := utils.ReadAddressesFromFile(addressFilePath)

	rpcClient := utils.GetRpcClient()
	wsClient := utils.GetWsClient(ctx)
	defer wsClient.Close()

	pool := utils.FetchAndUnmarshalPoolData(
		ctx,
		utils.GetSolanaAddressFromString(poolAddresses[0]),
		rpcClient,
	)

	sub := utils.SubscribeToWsAccount(wsClient, pool)
	defer sub.Unsubscribe()

	for {
		notification, err := sub.Recv()
		if err != nil {
			log.Fatalf("Error receiving notification: %v", err)
		}
		newData := utils.DecodeWhirlpoolData(notification.Value.Data)
		price, swapFee := utils.CalculatePoolPriceAndFees(newData.SqrtPrice, pool)
		utils.PrintPoolTable(pool, price, swapFee)
	}
}
