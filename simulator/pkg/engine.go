package main

import (
	"Pixel-Table/simulator/models"
	"math/rand"
	"path/filepath"
	"plugin"
	"strconv"
)

func processStreamData(c chan models.BotState, streamData []models.StreamEmission, symbol string) {
	// Create the initial trading bot state
	tradingBotState := models.BotState{
		Symbol:       symbol,
		Active:       false,
		MarketOrders: []models.MarketOrder{},
		Id:           rand.Float64(),
	}

	// Iterate through the stream data and determine whether to
	// consider selling or purchasing
	for _, streamEmission := range streamData {
		actionDetermination(streamEmission, &tradingBotState)
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

	// PLUGIN
	action = determineTradeAction(*tradingBotState, streamEmission)

	// Create a market order if a purchase or sell action is set
	switch action {
	case Purchase:
		marketOrder.Action = action

		// Set trading bot state values for sell processing
		tradingBotState.Active = true
	case Sell:
		marketOrder.Action = action

		// Set trading bot state values for future purchases
		tradingBotState.Active = false
	}

	// If the action is not a Wait, fulfill the market order
	if action != Wait {
		marketOrder.Price, _ = strconv.ParseFloat(streamEmission.Close, 32)
		tradingBotState.MarketOrders = append(tradingBotState.MarketOrders, marketOrder)
	}
}

func determineTradeAction(tradingBotState models.BotState, streamEmission models.StreamEmission) models.MarketOrderAction {
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
	processFunc, ok := symbol.(func(models.BotState, models.StreamEmission) models.MarketOrderAction)
	if !ok {
		panic("Plugin has no 'ProcessTrade' function")
	}
	// Uses the function to return results
	return processFunc(tradingBotState, streamEmission)
}
