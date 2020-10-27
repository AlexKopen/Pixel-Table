package models

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
	Price  float64
	Time   int64
}

type BotState struct {
	Symbol                string
	Active                bool
	MarketOrders          []MarketOrder
}

type Parameters struct {
	OrderSize                 float64
}
