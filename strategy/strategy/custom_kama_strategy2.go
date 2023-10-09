package strategy

// 自定义的均线策略 ， 使用sma1， kama5, kamaFast
// 使用kama5+kamaFast 入场，使用 sma1 出场

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type CustomKAMAStrategy2 struct {
	kama32 *indicator.KAMAIndicator
	//kama5      *indicator.KAMAIndicator
	kama16      *indicator.KAMAIndicator
	kama16_kama *indicator.KAMAIndicator
	//crossover   *indicator.CrossOverIndicator
	//crossunder  *indicator.CrossUnderIndicator
	//loaded bool // 只有
}

func NewCustomLAMAStrategy2Factory() model.Strategy {
	return &CustomKAMAStrategy2{}
}

func (es *CustomKAMAStrategy2) Name() string {
	return "CustomKAMAStrategy2"
}

func (es *CustomKAMAStrategy2) Init(ec *model.ContractStrategyContext) {
	es.kama32 = indicator.NewKAMAIndicator(ec.Kline, 32, 2, 30)
	//es.kama5 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 5, 2, 30)
	es.kama16 = indicator.NewKAMAIndicator(ec.Kline, 16, 2, 30)
	//es.kama16_kama = indicator.NewKAMAIndicator(es.kama16, 16, 2, 30)
	//es.crossover = indicator.NewCrossOverIndicator(ec.Kline, es.kama16.KAMALine, es.kama32.KAMALine)
	//es.crossunder = indicator.NewCrossUnderIndicator(ec.Kline, es.kama16.KAMALine, es.kama32.KAMALine)
}

func (es *CustomKAMAStrategy2) OnBar(ctx *model.ContractStrategyContext, ts int64) *model.StrategyResult {
	currentKNode, err := ctx.Kline.GetByTs(ts)
	if err != nil || currentKNode == nil || currentKNode.GetValue() == 0 {
		return nil
	}
	curPrice := currentKNode.GetValue()
	//if utils.AnyToBool(es.crossover.GetByTs(ts)) || utils.AnyToBool(es.crossunder.GetByTs(ts)) {
	//	es.loaded = true
	//}
	fast := es.kama16.GetCurrentFloat(ts)
	slow := es.kama32.GetCurrentFloat(ts)

	if fast > slow && (fast-slow)/slow > 0.001 {
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice), "")
	} else if fast < slow && (slow-fast)/fast > 0.001 {
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice), "")
	}
	return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice), "")
}
