package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/masakiva/orca-price-feed/internal/utils"
	"github.com/olekukonko/tablewriter"
)

type Pool struct {
	PoolAddress  string
	TokenASymbol string
	TokenBSymbol string
	Price        float64
	SwapFee      float64
}

func main() {
	ctx := context.Background()
	rpcClient := utils.GetRpcClient()
	// wallet := utils.GetWallet("./devnet-wallet.json")
	// balance := utils.GetSolBalance(ctx, wallet.PublicKey(), rpcClient)

	// Orca's SOL/USDC pool address
	// TODO: get addresses from user input
	stringPoolAddresses := []string{
		"Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE",
		"4Ui9QdDNuUaAGqCPcDSp191QrixLzQiLxJ1Gnqvz3szP",
	}

	var pools []Pool
	for _, stringPoolAddress := range stringPoolAddresses {
		pools = append(pools, Pool{PoolAddress: stringPoolAddress})
	}

	for i := range pools {
		solanaAddress, err := solana.PublicKeyFromBase58(pools[i].PoolAddress)
		if err != nil {
			log.Fatalf("failed to parse whirlpool address %s: %v", pools[i].PoolAddress, err)
		}
		poolData := utils.GetWhirlpoolData(rpcClient, solanaAddress)

		pools[i].TokenASymbol = utils.GetTokenSymbol(ctx, *poolData.TokenMintA, rpcClient)
		pools[i].TokenBSymbol = utils.GetTokenSymbol(ctx, *poolData.TokenMintB, rpcClient)
		pools[i].Price = utils.GetPoolCurrentPrice(ctx, poolData, rpcClient)
		pools[i].SwapFee = utils.GetSwapFee(pools[i].Price, poolData.FeeRate)
	}

	// Create a new table writer instance
	table := tablewriter.NewWriter(os.Stdout)

	// Set the table header
	table.SetHeader([]string{"Pool Address", "Token A", "Token B", "Price without fees"})

	// Append rows to the table
	for _, pool := range pools {
		table.Append([]string{
			pool.PoolAddress,
			pool.TokenASymbol,
			pool.TokenBSymbol,
			fmt.Sprintf("%.4f", pool.Price-pool.SwapFee), // Format price with 4 decimal places
		})
	}

	// Customize the table (optional)
	table.SetBorder(true)                      // Enable borders
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Align text to the left

	// Render the table to standard output
	table.Render()
}
