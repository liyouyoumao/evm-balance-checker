package tokens

import (
	"bytes"
	"context"
	"errors"
	"evm-balance-checker/bigint"
	"evm-balance-checker/contracts/metadatafetcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log/slog"
)

type Metadata struct {
	Symbol      string `json:"symbol"`
	Precision   uint8  `json:"precision"`
	TotalSupply string `json:"totalSupply"`
	Contract    string `json:"contract"`
	Icon        string `json:"icon"`
}

const rpcURL = "https://bsc-testnet-dataseed.bnbchain.org"
const contractAddress = "0xb2C827D9d20dc523199F6C658029D7Ac2D56491F"

func GetTokenMetadata(ctx context.Context, address string) (*Metadata, error) {
	// 连接到以太坊客户端
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	// 加载合约ABI
	mfabi, err := abi.JSON(bytes.NewReader([]byte(metadatafetcher.MetadatafetcherMetaData.ABI)))
	if err != nil {
		return nil, err
	}
	// 准备调用合约的参数
	target := common.HexToAddress(contractAddress)

	// 调用合约方法
	input, err := mfabi.Pack("getTokenMetadata", common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	output, err := client.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:   &target,
			Data: input,
		}, nil)
	if err != nil {
		return nil, err
	}

	if len(output) == 0 {
		panic("call failed output is empty")
	}

	out, err := mfabi.Unpack("getTokenMetadata", output)
	if err != nil {
		return nil, err
	}
	result := *abi.ConvertType(out[0], new(metadatafetcher.ERC20MetadataFetcherERC20Metadata)).(*metadatafetcher.ERC20MetadataFetcherERC20Metadata)

	slog.Info("output", "metadata", result)

	if !result.Valid {
		return nil, errors.New("non-standard ERC20 tokens")
	}

	return &Metadata{
		Symbol:      result.Symbol,
		Precision:   result.Decimals,
		TotalSupply: decimal.NewFromFloat(bigint.FromBig(result.TotalSupply).Readable(result.Decimals)).String(),
		Contract:    address,
		Icon:        "",
	}, nil
}
