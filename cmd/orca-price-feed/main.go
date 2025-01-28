package main

import (
	"context"

	"github.com/masakiva/orca-price-feed/internal/utils"
)

func main() {
	addressFilePath := utils.ParseCmdLineArgs()
	poolAddresses := utils.ReadAddressesFromFile(addressFilePath)

	ctx := context.Background()
	rpcClient := utils.GetRpcClient()
	pools := utils.UnmarshalPoolData(ctx, poolAddresses, rpcClient)

	utils.PrintPoolTable(pools)
}
