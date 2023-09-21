package indicator_base

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/model"
)

type IndicatorCommon struct {
	name     string
	Children []model.MarketIndicator
}

func NewIndicatorCommon() *IndicatorCommon {
	return &IndicatorCommon{Children: []model.MarketIndicator{}}
}

func (bl *IndicatorCommon) Name() string {
	return bl.name
}

func (bl *IndicatorCommon) AddChildrenIndicator(i model.MarketIndicator) {
	bl.Children = append(bl.Children, i)
}

func (bl *IndicatorCommon) TriggerChildren(ts int64, node model.DataNode) {
	for _, i := range bl.Children {
		i.AddData(ts, node)
		i.TriggerChildren(ts, node)
	}
}

func (bl *IndicatorCommon) PlotChildren(kline *charts.Kline, ratioLine *charts.Line) {
	for _, i := range bl.Children {
		i.DoPlot(kline, ratioLine)
		i.PlotChildren(kline, ratioLine)
	}
}
