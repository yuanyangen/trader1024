package strategy

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
)

type SingleSMAStrategy struct {
	sma *indicator.SimpleMovingAverageIndicator
	//fastSMA *indicator.SimpleMovingAverageIndicator
	//crossover  *indicator.CrossOverIndicator
	//crossunder *indicator.CrossUnderIndicator
}

func NewSingleSMAStrategy() *SingleSMAStrategy {
	return &SingleSMAStrategy{}
}

func (es *SingleSMAStrategy) Name() string {
	return "SingleSMA"
}

func (es *SingleSMAStrategy) Init(ec *MarketStrategyContext) {
	es.sma = indicator.NewSMAIndicator(ec.DailyData.Kline, 5)
}

func (es *SingleSMAStrategy) OnBar(ctx *MarketStrategyContext, ts int64) []*model.StrategyResult {
	currentKValue, err := ctx.DailyData.Kline.GetKnodeByTs(ts)
	if err != nil {
		return nil
	}
	sma, ok := es.sma.GetCurrentValue(ts).(float64)
	if !ok {
		return nil
	}
	if sma > currentKValue.Close && account.GetAccount().GetPositionByMarket(ctx.Market.MarketId) == nil {
		return []*model.StrategyResult{
			NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
			NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(currentKValue.Close)),
		}
	}
	//if under {
	//	return []*model.StrategyResult{
	//		NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
	//		NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(currentKValue.Close)),
	//	}
	//}
	return nil
}
