package model

import (
	"github.com/go-echarts/go-echarts/charts"
)

type IndicatorFactory func(period int64, parentIndicators ...MarketIndicator) MarketIndicator

type MarketIndicator interface {
	Name() string
	AddData(ts int64, node any)
	GetAllSortedData() []any
	GetLastByTsAndCount(ts int64, period int64) ([]any, error)
	GetByTs(ts int64) any
	DoPlot(page *charts.Kline, ratioLine *charts.Line)
	// common
	PlotChildren(kline *charts.Kline, ratioLine *charts.Line)
	AddChildrenIndicator(i MarketIndicator)
	TriggerChildren(ts int64, node any)
}
