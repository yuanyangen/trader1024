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

func (bl *IndicatorCommon) TriggerChildren(ts int64, node any) {
	for _, i := range bl.Children {
		i.AddData(ts, node)
	}
}

func (bl *IndicatorCommon) PlotChildren(p *charts.Page, page *charts.Kline) {
	for _, i := range bl.Children {
		i.DoPlot(p, page)
		i.PlotChildren(p, page)
	}
}
