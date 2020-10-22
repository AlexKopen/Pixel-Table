package main

type StreamEmission struct {
	OpenTime                 int64
	Open                     string
	High                     string
	Low                      string
	Close                    string
	Volume                   string
	CloseTime                int64
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
	Time   int64
}

type BotState struct {
	Active                bool
	PurchasePrice         float64
	MaxPriceSincePurchase float64
	MarketOrders          []MarketOrder
	Profit                float64
}

type Parameters struct {
	ChangeThresholdPercentage float64
	LossSellPercentage        float64
	GainSellPercentage        float64
}
