package main

var BotParameters = Parameters{
	ChangeThresholdPercentage: 0.0035,
	LossSellPercentage:        0.015,
	GainSellPercentage:        0.0015,
}

const (
	Purchase MarketOrderAction = iota
	Sell
	Wait
)

var Symbols = [...]string{
	//"LINK",
	//"ETH",
	//"DOT",
	//"BNB",
	//"EOS",
	//"ADA",
	//"BTC",
	//"TRX",
	//"XRP",
	//"XTZ",
	"UNI",
	//"LTC",
}
