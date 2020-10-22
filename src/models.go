package main

type StreamEmission struct {
	OpenTime                 uint
	Open                     string
	High                     string
	Low                      string
	Close                    string
	Volume                   string
	CloseTime                uint
	QuoteAssetVolume         string
	NumberOfTrades           uint
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
	Ignore                   string
}

type MarketOrderAction int

type MarketOrder struct {
	Action MarketOrderAction
	Symbol string
	Price  float64
	Time   uint
}

type BotState struct {
	Active                bool
	PurchasePrice         float64
	MaxPriceSincePurchase float64
	MarketOrders          []MarketOrder
}

type Parameters struct {
	ChangeThresholdPercentage float64
	LossSellPercentage        float64
	GainSellPercentage        float64
}
