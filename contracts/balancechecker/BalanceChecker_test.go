package balancechecker_test

import (
	"bytes"
	"context"
	"evm-balance-checker/contracts/balancechecker"
	"log"
	"log/slog"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const rpcURL = "https://bsc-dataseed.bnbchain.org"
const contractAddress = "0x58D3b153E328fa3b93fdD122A942d8Ba50bAFA1D"

func TestBalances(t *testing.T) {
	// 连接到以太坊客户端
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	// 加载合约ABI
	bcabi, err := abi.JSON(bytes.NewReader([]byte(balancechecker.BalancecheckerMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}
	// 准备调用合约的参数
	target := common.HexToAddress(contractAddress)

	var userTokens []balancechecker.BalanceCheckerUserToken
	t.Run("single call", func(t *testing.T) {
		// 准备调用合约的参数
		cases := []struct {
			user  string
			token string
		}{
			{
				// native token
				user:  "0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14",
				token: "0x0000000000000000000000000000000000000000",
			},
			{
				// erc20
				user:  "0x3d69EeB2058555aE32B17bf916BfCfF536a225CA",
				token: "0xbBF8F565995c3fDF890120e6AbC48c4f818b03c2",
			},
			{
				user:  "0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14",
				token: "0xbBF8F565995c3fDF890120e6AbC48c4f818b03c2",
			},
		}
		userTokens = make([]balancechecker.BalanceCheckerUserToken, 0, len(cases))
		for _, c := range cases {
			userTokens = append(userTokens, balancechecker.BalanceCheckerUserToken{
				User:  common.HexToAddress(c.user),
				Token: common.HexToAddress(c.token),
			})
		}

	})

	t.Run("invalid token", func(t *testing.T) {
		userTokens = []balancechecker.BalanceCheckerUserToken{
			{
				// not erc20
				User:  common.HexToAddress("0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14"),
				Token: common.HexToAddress("0x58d3b153e328fa3b93fdd122a942d8ba50bafa1d"),
			},
			{
				// not erc20
				User:  common.HexToAddress("0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14"),
				Token: common.HexToAddress("0x0000000000000000000000000000000000000001"),
			},

			{
				// not erc20
				User:  common.HexToAddress("0xeE905F03b4fb722D58e44ee4E56fa92Ab72c2B14"),
				Token: common.HexToAddress("0x1000000000000000000000000000000000000000"),
			},
		}
	})

	// 调用合约方法
	input, err := bcabi.Pack("balances", userTokens)

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

	var out0 []balancechecker.BalanceCheckerTokenBalance
	err = bcabi.UnpackIntoInterface(&out0, "balances", output)
	if err != nil {
		log.Fatal(err)
	}

	for _, tb := range out0 {
		slog.Info("output", "TokenBalance", tb)
	}
}
