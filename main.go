package main

import (
	"context"
	"encoding/json"
	"fmt"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/k0kubun/pp"
	"math"
	"time"
	conf "tinkoff-invest-status/config"
)

func main() {
	pp.ColoringEnabled = false

	config := conf.NewConfig()
	restClient := sdk.NewRestClient(config.Token)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accounts, err := restClient.Accounts(ctx)
	if err != nil {
		config.Logger.Fatalln("get accounts:", err)
	}

	result := map[string]map[string]interface{}{}

	for _, account := range accounts {
		positions, err := restClient.PositionsPortfolio(ctx, account.ID)
		if err != nil {
			config.Logger.Fatalln("get portfolio positions:", err)
		}
		account_id := "account_" + account.ID

		currencies, err := restClient.CurrenciesPortfolio(ctx, account.ID)
		if err != nil {
			config.Logger.Fatalln("get portfolio currencies:", err)
		}

		result[account_id] = map[string]interface{}{}
		result[account_id]["Positions"] = positions
		result[account_id]["Currencies"] = currencies

		summaryByCurrency := map[string]float64{}
		for _, item := range currencies {
			currency := string(item.Currency)
			summaryByCurrency[currency] += item.Balance
		}
		for _, item := range positions {
			if item.Ticker == "USD000UTSTOM" {
				continue
			}
			currency := string(item.AveragePositionPrice.Currency)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			orderbook, err := restClient.Orderbook(ctx, 1, item.FIGI)
			if err != nil {
				config.Logger.Fatalln("get instrument by figi:", err)
			}

			summaryByCurrency[currency] += item.Balance * orderbook.LastPrice
			summaryByCurrency[currency] = math.Floor(summaryByCurrency[currency]*100) / 100
		}

		result[account_id]["Summary"] = summaryByCurrency
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		config.Logger.Fatalln(err)
	}

	fmt.Printf("%v\n", string(jsonBytes))
}
