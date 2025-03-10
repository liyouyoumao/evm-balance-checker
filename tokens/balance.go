package tokens

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TokenData struct {
	Balance         string `json:"balance"`
	ContractAddress string `json:"contract_address"`
	Decimals        int    `json:"decimals"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
}

type ApiResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     []TokenData `json:"data"`
	NextPage int         `json:"next_page"`
	Count    int         `json:"count"`
}

const (
	apiKey    = "2thYXJz54CmPOXvJu7M8m9oDQbt"
	baseURL   = "https://api.chainbase.online/v1/account/tokens"
	rateLimit = time.Second / 2
	chanId    = 97
	limit     = 100
)

func GetAccountAllTokensBalances(ctx context.Context, address string) ([]TokenData, error) {
	return fetchTokenBalances(ctx, chanId, address, limit)
}

func fetchTokenBalances(ctx context.Context, chainID int, address string, limit int) ([]TokenData, error) {
	if limit > 100 {
		return nil, fmt.Errorf("limit must be <= 100")
	}
	var allTokens []TokenData
	page := 1
	throttle := time.Tick(rateLimit)

	for {
		select {
		case <-ctx.Done():
			return allTokens, nil
		case <-throttle:
			<-throttle
			url := fmt.Sprintf("%s?chain_id=%d&address=%s&limit=%d&page=%d", baseURL, chainID, address, limit, page)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return nil, err
			}
			req.Header.Set("x-api-key", apiKey)
			req.Header.Set("accept", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			var apiResponse ApiResponse
			if err := json.Unmarshal(body, &apiResponse); err != nil {
				return nil, err
			}

			if apiResponse.Code != 0 {
				if apiResponse.Code == 429 {
					time.Sleep(100 * time.Millisecond)
					continue
				}
				return nil, fmt.Errorf("API error: %s", apiResponse.Message)
			}

			allTokens = append(allTokens, apiResponse.Data...)

			if apiResponse.NextPage == 0 {
				return allTokens, nil
			}
			log.Printf("page: %d success\n", page)
			page = apiResponse.NextPage
		}
	}
}
