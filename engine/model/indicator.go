package model

import (
	"github.com/go-echarts/go-echarts/charts"
)

type IndicatorFactory func(period int64, parentIndicators ...MarketIndicator) MarketIndicator

type MarketIndicator interface {
	Name() string
	AddData(ts int64, node any)
	GetAllSortedData() []any
	GetByTsAndCount(ts int64, period int64) ([]any, error)
	GetByTs(ts int64) any
	DoPlot(p *charts.Page, page *charts.Kline)
	// common
	PlotChildren(p *charts.Page, page *charts.Kline)
	AddChildrenIndicator(i MarketIndicator)
}
