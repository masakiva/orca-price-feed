package utils

import (
	"bufio"
	"log"
	"os"

	"github.com/gagliardetto/solana-go"
)

func ParseCmdLineArgs() string {
	if len(os.Args) < 2 {
		log.Fatal("Please supply a file path containing whirlpool addresses as an argument.")
	}
	addressFilePath := os.Args[1]
	return addressFilePath
}

func ReadAddressesFromFile(addressFilePath string) (addresses []string) {
	file, err := os.Open(addressFilePath)
	if err != nil {
		log.Fatalf("failed to open address file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addresses = append(addresses, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read address file: %v", err)
	}
	return
}

func GetSolanaAddressFromString(addressStr string) (solanaAddress solana.PublicKey) {
	solanaAddress, err := solana.PublicKeyFromBase58(addressStr)
	if err != nil {
		log.Fatalf("failed to parse whirlpool address %s: %v", addressStr, err)
	}
	return
}
