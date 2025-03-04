package main

import (
	"bytes"
	"context"
	balancechecker "evm-balance-checker/contracts"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const rpcURL = "https://data-seed-prebsc-1-s1.bnbchain.org:8545"
const contractAddress = "0x98b18FCe8B7A7c5f398CD3a2D5D6352eff14BC24"

func main() {
	// 连接到以太坊客户端
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 调用 tokenBalance 方法
	user := common.HexToAddress("0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14")
	token := common.HexToAddress("0x0000000000000000000000000000000000000000")

	bcabi, err := abi.JSON(bytes.NewReader([]byte(balancechecker.BalancecheckerMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	userToken := balancechecker.BalanceCheckerUserToken{
		User:  user,
		Token: token,
	}

	input, err := bcabi.Pack("balances", user, []balancechecker.BalanceCheckerUserToken{userToken})
	if err != nil {
		log.Fatal(err)
	}

	target := common.HexToAddress(contractAddress)

	output, err := client.CallContract(
		context.Background(),
		ethereum.CallMsg{
			From: user,
			To:   &target,
			Data: input,
		}, nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(output) == 0 {
		panic("call failed output is empty")
	}

	var out0 []balancechecker.BalanceCheckerTokenBalance
	err = bcabi.UnpackIntoInterface(&out0, "balances", output)
	if err != nil {
		log.Fatal(err)
	}

}
