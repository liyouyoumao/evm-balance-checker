package metadatafetcher_test

import (
	"bytes"
	"context"
	"evm-balance-checker/contracts/metadatafetcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"log/slog"
	"math/big"
	"testing"
)

const rpcURL = "https://bsc-testnet-dataseed.bnbchain.org"
const contractAddress = "0xb2C827D9d20dc523199F6C658029D7Ac2D56491F"

func TestGetERC20TokenMetadata(t *testing.T) {
	// 连接到以太坊客户端
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	// 加载合约ABI
	mfabi, err := abi.JSON(bytes.NewReader([]byte(metadatafetcher.MetadatafetcherMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}
	// 准备调用合约的参数
	target := common.HexToAddress(contractAddress)
	var address common.Address
	var expect metadatafetcher.ERC20MetadataFetcherERC20Metadata
	t.Run("eoa address", func(t *testing.T) {
		address = common.HexToAddress("0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14")
		expect = metadatafetcher.ERC20MetadataFetcherERC20Metadata{
			Symbol:      "",
			Decimals:    0,
			TotalSupply: new(big.Int).SetUint64(0),
			Valid:       false,
		}
	})

	t.Run("erc20 address", func(t *testing.T) {
		address = common.HexToAddress("0x84b9B910527Ad5C03A9Ca831909E21e236EA7b06")
		totalSupply, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
		expect = metadatafetcher.ERC20MetadataFetcherERC20Metadata{
			Symbol:      "LINK",
			Decimals:    18,
			TotalSupply: totalSupply,
			Valid:       true,
		}
	})

	// 调用合约方法
	input, err := mfabi.Pack("getTokenMetadata", address)
	if err != nil {
		log.Fatal(err)
	}
	output, err := client.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:   &target,
			Data: input,
		}, nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(output) == 0 {
		panic("call failed output is empty")
	}

	out0, err := mfabi.Unpack("getTokenMetadata", output)
	if err != nil {
		log.Fatal(err)
	}
	got := *abi.ConvertType(out0[0], new(metadatafetcher.ERC20MetadataFetcherERC20Metadata)).(*metadatafetcher.ERC20MetadataFetcherERC20Metadata)

	if expect.Symbol != got.Symbol {
		log.Fatal("unexpected symbol")
	}

	if expect.Decimals != got.Decimals {
		log.Fatal("unexpected decimals")
	}

	if expect.TotalSupply.Cmp(got.TotalSupply) != 0 {
		log.Fatal("unexpected total supply")
	}

	if expect.Valid != got.Valid {
		log.Fatal("unexpected validity")
	}

	slog.Info("output", "metadata", out0)

}
