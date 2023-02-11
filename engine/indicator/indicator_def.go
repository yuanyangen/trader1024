package indicator

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/data_feed"
)

type MarketIndicator interface {
	Name() string
	//GetCurrentValue(int64) any
	AddData(ts int64, node any)
	GetAllSortedData() []any
	DoPlot(page *charts.Kline)
}

type DailyIndicators struct {
	Kline          *KLineIndicator
	ReceiveChannel chan *data_feed.Data
	Indicators     []MarketIndicator
}
