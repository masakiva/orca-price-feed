package main

import (
	"context"
	"time"

	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	addressFilePath := utils.ParseCmdLineArgs()
	poolAddresses := utils.ReadAddressesFromFile(addressFilePath)

	ctx := context.Background()
	rpcClient := utils.GetRpcClient()

	for {
		pools := utils.FetchAndUnmarshalPoolData(ctx, poolAddresses, rpcClient)
		utils.PrintPoolTable(pools)
		time.Sleep(5 * time.Second)
	}
}
