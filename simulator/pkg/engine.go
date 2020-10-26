package main

import (
	"path/filepath"
	models "pixel-table/simulator/shared"
	"plugin"
	"strconv"
)

func processStreamData(c chan models.BotState, streamData []models.StreamEmission, symbol string) {
	// Create the initial trading bot state
	tradingBotState := models.BotState{
		Symbol:                symbol,
		CurrentPrice:          0,
		Active:                false,
		PurchasePrice:         0,
		MaxPriceSincePurchase: 0,
		MarketOrders:          []models.MarketOrder{},
		ActiveAmount:          0,
		Profit:                0,
		PercentChange:         0,
	}

	// Iterate through the stream data and determine whether to
	// consider selling or purchasing
	for _, streamEmission := range streamData {
		tradingBotState.CurrentPrice, _ = strconv.ParseFloat(streamEmission.Close, 32)
		actionDetermination(streamEmission, &tradingBotState)
	}

	// Adjust the profit if a trade hasn't sold after the emissions are done being processed
	if tradingBotState.Active {
		tradingBotState.Profit += tradingBotState.ActiveAmount * tradingBotState.CurrentPrice
	}

	// Output the final bot state
	c <- tradingBotState
}

func actionDetermination(streamEmission models.StreamEmission, tradingBotState *models.BotState) {
	// Set default action and order values
	action := Wait
	marketOrder := models.MarketOrder{
		Time: streamEmission.CloseTime,
	}

	// Convert the open and close price to floats
	openPrice, _ := strconv.ParseFloat(streamEmission.Open, 32)
	closePrice, _ := strconv.ParseFloat(streamEmission.Close, 32)
	tradingBotState.PercentChange = (closePrice - openPrice) / openPrice

	// PLUGIN
	action = determineTradeAction(tradingBotState)

	// Create a market order if a purchase or sell action is set
	switch action {
	case Purchase:
		marketOrder.Action = action

		// Set trading bot state values for sell processing
		tradingBotState.PurchasePrice = closePrice
		tradingBotState.MaxPriceSincePurchase = closePrice
		tradingBotState.Active = true
		tradingBotState.ActiveAmount = BotParameters.OrderSize / closePrice
		tradingBotState.Profit -= BotParameters.OrderSize
	case Sell:
		marketOrder.Action = action

		// Set trading bot state values for future purchases
		tradingBotState.Active = false
		tradingBotState.Profit += tradingBotState.ActiveAmount * closePrice
	}

	// If the action is not a Wait, fulfill the market order
	if action != Wait {
		marketOrder.Price = closePrice
		tradingBotState.MarketOrders = append(tradingBotState.MarketOrders, marketOrder)
	}
}

func determineTradeAction(tradingBotState *models.BotState) models.MarketOrderAction {
	// Glob – Gets the plugin to be loaded
	plugins, err := filepath.Glob("simulator/plugins/trade-bot/trade-bot.so")
	if err != nil {
		panic(err)
	}
	// Open – Loads the plugin
	p, loadErr := plugin.Open(plugins[0])
	if loadErr != nil {
		panic(loadErr)
	}
	// Lookup – Searches for a symbol name in the plugin
	symbol, lookupErr := p.Lookup("ProcessTrade")
	if lookupErr != nil {
		panic(lookupErr)
	}
	// symbol – Checks the function signature
	processFunc, ok := symbol.(func(*models.BotState) models.MarketOrderAction)
	if !ok {
		panic("Plugin has no 'ProcessTrade' function")
	}
	// Uses the function to return results
	return processFunc(tradingBotState)
}
