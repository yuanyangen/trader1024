package strategy

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
)

type DualSMAStrategy struct {
	slowSMA    *indicator.SMAIndicator
	fastSMA    *indicator.SMAIndicator
	crossover  *indicator.CrossOverIndicator
	crossunder *indicator.CrossUnderIndicator
}

func NewDualSMAStrategyFactory() Strategy {
	return &DualSMAStrategy{}
}

//func (es *DualSMAStrategy) Indicators() []indicator.MarketIndicator {
//	return []indicator.MarketIndicator{
//		es.sma10,
//		es.sma5,
//	}
//}

func (es *DualSMAStrategy) Name() string {
	return "DualSMA"
}

func (es *DualSMAStrategy) Init(ec *MarketStrategyContext) {
	es.slowSMA = indicator.NewSMAIndicator(ec.DailyData.Kline, 10)
	es.fastSMA = indicator.NewSMAIndicator(ec.DailyData.Kline, 5)
	es.crossover = indicator.NewCrossOverIndicator(ec.DailyData.Kline, es.fastSMA.SMALine, es.slowSMA.SMALine)
	es.crossunder = indicator.NewCrossUnderIndicator(ec.DailyData.Kline, es.fastSMA.SMALine, es.slowSMA.SMALine)
}

func (es *DualSMAStrategy) OnBar(ctx *MarketStrategyContext, ts int64) []*model.StrategyResult {
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
