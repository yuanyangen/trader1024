package indicator

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/data_feed"
)

type MarketIndicator interface {
	Name() string
	GetCurrentValue(int64) float64
	AddData(ts int64, node any)
	GetAllSortedData() []any
	DoPlot(page *charts.Kline)
}

type DailyIndicators struct {
	Line           *KLineIndicator
	ReceiveChannel chan *data_feed.Data
	Indicators     []MarketIndicator
}

type LineType int64

const LineType_Day = 1
const LineType_Minite = 2
const LineType_Hour = 3
