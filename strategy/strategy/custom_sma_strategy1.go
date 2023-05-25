package strategy

// 自定义的均线策略 ， 使用sma1， kama5, kama20
// 使用kama5+kama20 入场，使用 sma1 出场

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type CustomSMAStrategy struct {
	sma10  *indicator.SMAIndicator
	sma5   *indicator.SMAIndicator
	sma2   *indicator.SMAIndicator
	kama10 *indicator.KAMAIndicator
	kama5  *indicator.KAMAIndicator
	kama2  *indicator.KAMAIndicator
	loaded bool // 只有
}

func NewCustomSMAStrategyFactory() model.Strategy {
	return &CustomSMAStrategy{}
}

func (es *CustomSMAStrategy) Name() string {
	return "CustomSMAStrategy"
}

func (es *CustomSMAStrategy) Init(ec *model.MarketStrategyContext) {
	es.sma10 = indicator.NewSMAIndicator(ec.DailyData.Kline, 10)
	es.sma5 = indicator.NewSMAIndicator(ec.DailyData.Kline, 3)
	es.sma2 = indicator.NewSMAIndicator(ec.DailyData.Kline, 1)
	es.kama10 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 30, 2, 20)
	es.kama5 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 5, 2, 20)
	es.kama2 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 2, 2, 30)
}

func (es *CustomSMAStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) []*model.StrategyResult {
	currentKValue := model.NewKnodeFromAny(ctx.DailyData.Kline.GetByTs(ts))
	if currentKValue == nil {
		return nil
	}
	if es.kama2.GetCurrentFloat(ts) == 0 || es.kama5.GetCurrentFloat(ts) == 0 || es.kama10.GetCurrentFloat(ts) == 0 {
		return nil
	}
	position := account.GetAccount().GetPositionByMarket(ctx.Market.MarketId)
	curPrice := (currentKValue.Open + currentKValue.Close) / 2
	if long(es.kama2.GetCurrentFloat(ts), es.kama5.GetCurrentFloat(ts), es.kama10.GetCurrentFloat(ts)) {
		if position.IsEmpty() && es.loaded {
			return []*model.StrategyResult{
				model.NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(curPrice)),
			}
		}
	} else if short(es.kama2.GetCurrentFloat(ts), es.kama5.GetCurrentFloat(ts), es.kama10.GetCurrentFloat(ts)) {
		if position.IsEmpty() && es.loaded {
			return []*model.StrategyResult{
				model.NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(curPrice)),
			}
		}
	} else {
		es.loaded = true
		if !position.IsEmpty() {
			return []*model.StrategyResult{
				model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
			}
		}
	}

	return nil
}

func long(fast, mid, slow float64) bool {
	return fast > mid && mid > slow
}
func short(fast, mid, slow float64) bool {
	return fast < mid && mid < slow
}
