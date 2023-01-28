package model

import (
	"github.com/go-echarts/go-echarts/charts"
)

type MarketIndicator interface {
	Name() string
	GetCurrentValue(int64) float64
	AddData(ts int64, node any)
	GetAllSortedData() []any
	DoPlot(page *charts.Kline)
}

type GlobalIndicator interface {
	AddData(ctx *GlobalMsg)
	DoPlot(page *charts.Page)
}

type KLineIndicator struct {
	*BaseLine
	Indicators []MarketIndicator
}
