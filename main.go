package main

import (
	"context"
	"evm-balance-checker/tokens"
	"log"
	"log/slog"
)

func main() {
	var token = "0x84b9B910527Ad5C03A9Ca831909E21e236EA7b06"
	metadata, err := tokens.GetTokenMetadata(context.Background(), token)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("metadata: ", metadata)

	var account = "0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14"
	balances, err := tokens.GetAccountAllTokensBalances(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}

	for _, balance := range balances {
		slog.Info("balance: ", balance)
	}
}
