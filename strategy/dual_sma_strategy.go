package strategy

import (
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
)

type DualSMAStrategy struct {
	slowSMA *indicator.SimpleMovingAverageIndicator
	fastSMA *indicator.SimpleMovingAverageIndicator
}

func NewDualSMAStrategy() *DualSMAStrategy {
	return &DualSMAStrategy{}
}
func (es *DualSMAStrategy) Indicators() []indicator.MarketIndicator {
	return []indicator.MarketIndicator{
		es.slowSMA,
		es.fastSMA,
	}
}

func (es *DualSMAStrategy) Init(ec *model.MarketStrategyContext) {
	es.slowSMA = indicator.NewSMAIndicator(ec.DailyData.Line, 15)
	es.fastSMA = indicator.NewSMAIndicator(ec.DailyData.Line, 5)
}

func (es *DualSMAStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) {
	es.slowSMA.GetCurrentValue(ts)
}
