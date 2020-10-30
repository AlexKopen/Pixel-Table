package main

import "Pixel-Table/simulator/models"

var BotParameters = models.Parameters{
	OrderSize: 100.00,
}

const (
	Purchase models.MarketOrderAction = iota
	Sell
	Wait
)
