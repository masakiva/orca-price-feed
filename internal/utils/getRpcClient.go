package utils

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gagliardetto/solana-go/rpc"
)

func GetRpcClient() *rpc.Client {
	endpointUrl := os.Getenv("RPC_URL")
	if endpointUrl != "" {
		_, err := url.ParseRequestURI(endpointUrl)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("RPC_URL not defined, defaulting to Solana mainnet beta")
		endpointUrl = rpc.MainNetBeta_RPC
	}
	return rpc.New(endpointUrl)
}
