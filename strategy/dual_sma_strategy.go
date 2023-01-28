package strategy

import (
	"github.com/yuanyangen/trader1024/engine/indicator/market"
	"github.com/yuanyangen/trader1024/engine/model"
)

type DualSMAStrategy struct {
	slowSMA *market.SimpleMovingAverageIndicator
	fastSMA *market.SimpleMovingAverageIndicator
}

func NewDualSMAStrategy() *DualSMAStrategy {
	return &DualSMAStrategy{}
}
func (es *DualSMAStrategy) Indicators() []model.MarketIndicator {
	return []model.MarketIndicator{
		es.slowSMA,
		es.fastSMA,
	}
}

func (es *DualSMAStrategy) Init(ec *model.MarketStrategyContext) {
	es.slowSMA = market.NewSMAIndicator(ec.DailyData.Line, 5)
	es.fastSMA = market.NewSMAIndicator(ec.DailyData.Line, 10)
}

func (es *DualSMAStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) {
	es.slowSMA.GetCurrentValue(ts)
}
