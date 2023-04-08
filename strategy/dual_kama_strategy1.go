package strategy

// 自定义的均线策略 ， 使用sma1， kama5, kama20
// 使用kama5+kama20 入场，使用 sma1 出场

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
)

type DualKAMAStrategy struct {
	kama10     *indicator.KAMAIndicator
	kama2      *indicator.KAMAIndicator
	crossover  *indicator.CrossOverIndicator
	crossunder *indicator.CrossUnderIndicator
	loaded     bool // 只有
}

func NewDualKAMAStrategyFactory() Strategy {
	return &DualKAMAStrategy{}
}

func (es *DualKAMAStrategy) Name() string {
	return "DualKAMAStrategy"
}

func (es *DualKAMAStrategy) Init(ec *MarketStrategyContext) {
	es.kama10 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 10, 2, 30)
	es.kama2 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 2, 2, 30)
	es.crossover = indicator.NewCrossOverIndicator(ec.DailyData.Kline, es.kama2.KAMALine, es.kama10.KAMALine)
	es.crossunder = indicator.NewCrossUnderIndicator(ec.DailyData.Kline, es.kama2.KAMALine, es.kama10.KAMALine)
}

func (es *DualKAMAStrategy) OnBar(ctx *MarketStrategyContext, ts int64) []*model.StrategyResult {
	over := es.crossover.GetCurrentValue(ts)
	under := es.crossunder.GetCurrentValue(ts)
	currentKValue, err := ctx.DailyData.Kline.GetKnodeByTs(ts)
	if err != nil {
		return nil
	}
	if over {
		return []*model.StrategyResult{
			NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
			NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(currentKValue.Close)),
		}
	}
	if under {
		return []*model.StrategyResult{
			NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
			NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(currentKValue.Close)),
		}
	}
	return nil
}
