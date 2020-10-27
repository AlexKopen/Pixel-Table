package main

import (
	"pixel-table/simulator/models"
	"strconv"
)

func ProcessTrade(tradingBotState models.BotState, streamEmission models.StreamEmission) models.MarketOrderAction {
	// Convert the open and close price to floats
	openPrice, _ := strconv.ParseFloat(streamEmission.Open, 32)
	closePrice, _ := strconv.ParseFloat(streamEmission.Close, 32)
	percentChange := (closePrice - openPrice) / openPrice

	// Wait by default
	var action = 2

	// Sell
	if tradingBotState.Active && percentChange > 0.01 {
		action = 1
	}

	// Purchase
	if !tradingBotState.Active && percentChange > 0.01 {
		action = 0
	}

	return models.MarketOrderAction(action)
}