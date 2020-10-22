package main

var BotParameters = Parameters{
	ChangeThresholdPercentage: 0.003,
	LossSellPercentage:        0.002,
	GainSellPercentage:        0.002,
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
