package model

import "github.com/go-echarts/go-echarts/charts"

type Indicator interface {
	FillIndicatorToLine(line *KLine, ts int64, currentNode *KLineNode)
	Name() string
	GetValueFromKNode(dn *KLineNode) float64
}

type Plotter interface {
	DoPlot(page *charts.Page)
	Name() string
}
