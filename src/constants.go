package main

var UseTestData = false

var BotParameters = Parameters{
	ChangeThresholdPercentage: 0.0035,
	LossSellPercentage:        0.015,
	GainSellPercentage:        0.0015,
}

var BotParametersTest = Parameters{
	ChangeThresholdPercentage: 0.003,
	LossSellPercentage:        0.02,
	GainSellPercentage:        0.002,
}

const (
	Purchase MarketOrderAction = iota
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
