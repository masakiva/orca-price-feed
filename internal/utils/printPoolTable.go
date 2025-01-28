package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintPoolTable(pools []Pool) {
	// Create a new table writer instance
	table := tablewriter.NewWriter(os.Stdout)

	// Set the table header
	table.SetHeader([]string{"Pool Address", "Token A", "Token B", "Current price without fees"})

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
