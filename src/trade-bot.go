package main

import (
	"fmt"
	"math"
	"strconv"
)

func processStreamData(streamData []StreamEmission) {
	// Create the initial trading bot state
	tradingBotState := BotState{
		Active:                false,
		PurchasePrice:         0,
		MaxPriceSincePurchase: 0,
		MarketOrders:          []MarketOrder{},
	}

	// Iterate through the stream data and determine whether to
	// consider selling or purchasing
	for _, streamEmission := range streamData {
		actionDetermination(streamEmission, &tradingBotState)
	}

}

func actionDetermination(streamEmission StreamEmission, tradingBotState *BotState) {
	// Set default action and order values
	action := Wait
	marketOrder := MarketOrder{
		Symbol: streamEmission.Symbol,
		Time:   streamEmission.CloseTime,
	}

	// Convert the open and close price to floats
	openPrice, _ := strconv.ParseFloat(streamEmission.Open, 32)
	closePrice, _ := strconv.ParseFloat(streamEmission.Close, 32)

	// Calculate the percent change
	percentChange := (closePrice - openPrice) / openPrice

	// Consider selling if the trade is active, otherwise consider a purchase
	switch tradingBotState.Active {
	case true:
		// SELL LOGIC
		// Update trading bot state values
		tradingBotState.MaxPriceSincePurchase = math.Max(tradingBotState.MaxPriceSincePurchase, closePrice)
		priceFallLossTriggered := closePrice <= (tradingBotState.PurchasePrice - (tradingBotState.PurchasePrice * BotParameters.LossSellPercentage))
		priceHasRisenEnough := closePrice >= (tradingBotState.PurchasePrice + (tradingBotState.PurchasePrice * BotParameters.GainSellPercentage))
		priceFallGainTriggered := closePrice <= tradingBotState.MaxPriceSincePurchase-(tradingBotState.MaxPriceSincePurchase*BotParameters.GainSellPercentage)

		// Sell if the price has fallen too far below the purchase point
		if priceFallLossTriggered {
			fmt.Println("sell loss")
			action = Sell
		} else if priceHasRisenEnough && priceFallGainTriggered {
			//	Sell if the price has risen enough from the purchase price and also fallen too far below the maximum price
			fmt.Println("sell profit")
			action = Sell
		}
	case false:
		// PURCHASE LOGIC
		// Purchase if the percent change has passed the defined threshold
		if percentChange >= BotParameters.ChangeThresholdPercentage {
			fmt.Println("purchase")
			action = Purchase
		}
	}

	// Create a market order if a purchase or sell action is set
	switch action {
	case Purchase:
		marketOrder.Action = Purchase

		// Set trading bot state values for sell processing
		tradingBotState.PurchasePrice = closePrice
		tradingBotState.MaxPriceSincePurchase = closePrice
		tradingBotState.Active = true
	case Sell:
		marketOrder.Action = Sell

		// Reset trading bot state for future purchases
		tradingBotState.Active = false
	}

	// If the action is not a Wait, fulfill the market order
	if action != Wait {
		marketOrder.Price = closePrice
		tradingBotState.MarketOrders = append(tradingBotState.MarketOrders, marketOrder)
	} else {
		fmt.Println("wait")
	}
}
