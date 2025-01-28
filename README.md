## Usage

You can run the program by supplying the path to a file containing the
addresses of the pools you want to fetch the prices for, on Solana mainnet-beta:
```sh
go run cmd/orca-price-feed/main.go <path-to-address-file>
```

You can use this sample file for example:
```sh
go run cmd/orca-price-feed/main.go pool-addresses.txt
```

You can also set the `RPC_URL` environment variable to specify a custom Solana RPC endpoint
```sh
RPC_URL=https://api.devnet.solana.com go run cmd/orca-price-feed/main.go <path-to-address-file>
```
