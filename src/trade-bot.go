package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func processStreamData(streamData []StreamEmission, symbol string) {
	// Create the initial trading bot state
	tradingBotState := BotState{
		Active:                false,
		PurchasePrice:         0,
		MaxPriceSincePurchase: 0,
		MarketOrders:          []MarketOrder{},
		Profit:                0,
	}

	// Iterate through the stream data and determine whether to
	// consider selling or purchasing
	for _, streamEmission := range streamData {
		actionDetermination(streamEmission, &tradingBotState, symbol)
	}

	fmt.Println(tradingBotState.Profit)
}

func actionDetermination(streamEmission StreamEmission, tradingBotState *BotState, symbol string) {
	// Set default action and order values
	action := Wait
	marketOrder := MarketOrder{
		Symbol: symbol,
		Time:   streamEmission.CloseTime,
	}
	// Test data parameters detection
	botParameters := BotParameters
	if UseTestData {
		botParameters = BotParametersTest
	} else {
		botParameters = BotParameters
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
		priceFallLossTriggered := closePrice <= (tradingBotState.PurchasePrice - (tradingBotState.PurchasePrice * botParameters.LossSellPercentage))
		priceHasRisenEnough := closePrice >= (tradingBotState.PurchasePrice + (tradingBotState.PurchasePrice * botParameters.GainSellPercentage))
		priceFallGainTriggered := closePrice <= tradingBotState.MaxPriceSincePurchase-(tradingBotState.MaxPriceSincePurchase*botParameters.GainSellPercentage)

		// Sell if the price has fallen too far below the purchase point
		if priceFallLossTriggered {
			fmt.Printf("sell loss - %s\n", timeFormatted(streamEmission.CloseTime))
			action = Sell
		} else if priceHasRisenEnough && priceFallGainTriggered {
			//	Sell if the price has risen enough from the purchase price and also fallen too far below the maximum price
			fmt.Printf("sell profit - %s\n", timeFormatted(streamEmission.CloseTime))
			action = Sell
		}
	case false:
		// PURCHASE LOGIC
		// Purchase if the percent change has passed the defined threshold
		if percentChange >= botParameters.ChangeThresholdPercentage {
			fmt.Printf("purchase - %s\n", timeFormatted(streamEmission.CloseTime))
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
		tradingBotState.Profit = tradingBotState.Profit - closePrice
	case Sell:
		marketOrder.Action = action

		// Set trading bot state values for future purchases
		tradingBotState.Active = false
		tradingBotState.Profit = tradingBotState.Profit + closePrice
	}

	// If the action is not a Wait, fulfill the market order
	if action != Wait {
		marketOrder.Price = closePrice
		tradingBotState.MarketOrders = append(tradingBotState.MarketOrders, marketOrder)
	} else if UseTestData {
		fmt.Println("wait")
	}
}

func timeFormatted(timestamp int64) time.Time {
	// Shave off the last 3 digits from the timestamp for the Unix() function to work properly
	tm := time.Unix(timestamp/1e3, 0)
	return tm
}
