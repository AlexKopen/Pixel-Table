package main

import "pixel-table/simulator/models"

var BotParameters = models.Parameters{
	OrderSize:                 100.00,
}

const (
	Purchase models.MarketOrderAction = iota
	Sell
	Wait
)

var Symbols = [...]string{
	"LINK",
	"ETH",
	"DOT",
	"BNB",
	"EOS",
	"ADA",
	"BTC",
	"TRX",
	"XRP",
	"XTZ",
	"UNI",
	"LTC",
}
