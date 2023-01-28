package indicator

type DailyData struct {
	Line           *KLineIndicator
	DataFeed       DataFeed
	ReceiveChannel chan *Data
	Indicators     []MarketIndicator
}
