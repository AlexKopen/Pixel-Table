package main

import (
	models "pixel-table/simulator/shared"
)

func ProcessTrade(tradingBotState *models.BotState) models.MarketOrderAction {
	// Wait by default
	var action = 2
	if tradingBotState.Active && tradingBotState.PercentChange > 0.01 {
		// Sell
		action = 1
	}

	if !tradingBotState.Active && tradingBotState.PercentChange > 0.01 {
		// Purchase
		action = 0
	}

	return models.MarketOrderAction(action)
}
