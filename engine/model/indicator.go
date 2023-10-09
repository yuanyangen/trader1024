package model

import (
	"github.com/go-echarts/go-echarts/charts"
)

type IndicatorFactory func(period int64, parentIndicators ...ContractIndicator) ContractIndicator

type ContractIndicator interface {
	Name() string
	AddData(ts int64, node DataNode)
	GetAllSortedData() []DataNode
	GetLastByTsAndCount(ts int64, period int64) ([]DataNode, error)
	GetByTs(ts int64) (DataNode, error)
	DoPlot(page *charts.Kline, ratioLine *charts.Line)
	// common
	PlotChildren(kline *charts.Kline, ratioLine *charts.Line)
	AddChildrenIndicator(i ContractIndicator)
	TriggerChildren(ts int64, node DataNode)
}
