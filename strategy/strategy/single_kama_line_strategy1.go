package strategy

// 自定义的均线策略 ， 使用sma1， kama5, kamaFast
// 使用kama5+kamaFast 入场，使用 sma1 出场

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type SingleKAMAStrategy struct {
	kama *indicator.KAMAIndicator
	sma  *indicator.SMAIndicator
}

func NewSingleKAMAlineStrategyFactory() model.Strategy {
	return &SingleKAMAStrategy{}
}

func (es *SingleKAMAStrategy) Name() string {
	return "SingleKAMAStrategy"
}

func (es *SingleKAMAStrategy) Init(ec *model.ContractStrategyContext) {
	//es.kama5 = indicator.NewKAMAIndicator(ec.Kline, 2, 2, 5)
	//es.kama5_kama3 = indicator.NewKAMAIndicator(es.kama5, 3, 5, 10)

	//es.kama5_kama2 = indicator.NewKAMAIndicator(es.kama5, 2, 4, 6)
	es.kama = indicator.NewKAMAIndicator(ec.Kline, 3, 3, 3)
	es.sma = indicator.NewSMAIndicator(ec.Kline, 1)
}

func (es *SingleKAMAStrategy) OnBar(ctx *model.ContractStrategyContext, ts int64) *model.StrategyResult {
	//s := 0.01
	currentKNode, err := ctx.Kline.GetByTs(ts)
	if err != nil || currentKNode == nil || currentKNode.GetValue() == 0 {
		return nil
	}
	curPrice := currentKNode.GetValue()

	kamaValue := es.kama.GetCurrentFloat(ts)
	v, _ := es.sma.GetByTs(ts)
	if v != nil {
		kamaValue = v.GetValue()
	}

	gap := 0.001
	if curPrice < kamaValue && (kamaValue-curPrice)/kamaValue > gap {
		reason := fmt.Sprintf("(kama(%v)-cur(%v))/kama(%v)=%v>gap(%v)", kamaValue, curPrice, kamaValue, (kamaValue-curPrice)/kamaValue, gap)
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice), reason)
	} else if curPrice > kamaValue && (curPrice-kamaValue)/curPrice > gap {
		reason := fmt.Sprintf("(cur(%v)-kama(%v))/cur(%v)=%v>gap(%v)", curPrice, kamaValue, curPrice, (curPrice-kamaValue)/curPrice, gap)
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice), reason)
	} else {
		return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice), "not_long_or_short")
	}
}
