package model

type GlobalMsg struct {
	TimeStamp    int64
	TotalAccount float64
}

type MarketStrategyContext struct {
	DailyData *DailyData
	Broker    Broker
	Account   *Account
}

type DailyData struct {
	Line           *KLineIndicator
	DataFeed       DataFeed
	ReceiveChannel chan *Data
	Indicators     []MarketIndicator
}
