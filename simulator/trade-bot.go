package main

import (
	"math"
	"strconv"
)

func processStreamData(c chan BotState, streamData []StreamEmission, symbol string) {
	// Create the initial trading bot state
	tradingBotState := BotState{
		Symbol:                symbol,
		CurrentPrice:          0,
		Active:                false,
		PurchasePrice:         0,
		MaxPriceSincePurchase: 0,
		MarketOrders:          []MarketOrder{},
		ActiveAmount:          0,
		Profit:                0,
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

func actionDetermination(streamEmission StreamEmission, tradingBotState *BotState) {
	// Set default action and order values
	action := Wait
	marketOrder := MarketOrder{
		Time: streamEmission.CloseTime,
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
			action = Sell
		} else if priceHasRisenEnough && priceFallGainTriggered {
			//	Sell if the price has risen enough from the purchase price and also fallen too far below the maximum price
			action = Sell
		}
	case false:
		// PURCHASE LOGIC
		// Purchase if the percent change has passed the defined threshold
		if percentChange >= BotParameters.ChangeThresholdPercentage {
			action = Purchase
		}
	}

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
