package main

import (
	"context"
	"crypto/ecdsa"
	"evm-balance-checker/contracts/balancechecker"
	"evm-balance-checker/contracts/metadatafetcher"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// 配置信息
const (
	rpcURL = "https://data-seed-prebsc-1-s1.bnbchain.org:8545" // BSC测试网
	// 请替换为您自己的私钥（不要在生产环境中硬编码私钥）
	privateKey = ""
)

func main() {
	deploy("metadatafetcher")
}

func deploy(name string) {
	// 连接到以太坊客户端
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("无法连接到以太坊客户端: %v", err)
	}
	defer client.Close()

	// 加载私钥
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("无法加载私钥: %v", err)
	}

	// 获取公钥地址
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法获取公钥")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("无法获取nonce: %v", err)
	}

	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("无法获取gas价格: %v", err)
	}

	// 创建交易选项
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("无法获取链ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("无法创建交易选项: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // 不发送ETH
	auth.GasLimit = uint64(3000000) // 设置gas限制
	auth.GasPrice = gasPrice

	// 部署合约
	fmt.Printf("正在部署%s合约...\n", name)
	var tx *types.Transaction
	var address common.Address

	if name == "balancechecker" {
		address, tx, _, err = balancechecker.DeployBalancechecker(auth, client)

	} else if name == "metadatafetcher" {
		address, tx, _, err = metadatafetcher.DeployMetadatafetcher(auth, client)
	} else {
		log.Fatalf("not supported")
	}
	if err != nil {
		log.Fatalf("无法部署合约: %v", err)
	}

	fmt.Printf("合约部署交易已发送: %s\n", tx.Hash().Hex())
	fmt.Printf("等待交易确认...\n")

	// 等待交易被确认
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("等待交易确认时出错: %v", err)
	}

	if receipt.Status == 0 {
		log.Fatal("交易失败")
	}

	fmt.Printf("合约已部署到地址: %s\n", address.Hex())
	fmt.Printf("Gas使用量: %d\n", receipt.GasUsed)

	fmt.Println("部署完成!")
}
