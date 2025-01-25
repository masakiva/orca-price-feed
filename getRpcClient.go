package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gagliardetto/solana-go/rpc"
)

func getRpcClient() *rpc.Client {
	endpointUrl := os.Getenv("RPC_URL")
	if endpointUrl != "" {
		_, err := url.ParseRequestURI(endpointUrl)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("RPC_URL not defined, defaulting to Solana devnet")
		endpointUrl = rpc.DevNet_RPC
	}
	return rpc.New(endpointUrl)
}
