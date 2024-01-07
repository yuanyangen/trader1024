package strategy

// 自定义的均线策略 ， 使用sma1， loess5, loessFast
// 使用loess5+loessFast 入场，使用 sma1 出场

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type SingleLOESSStrategy struct {
	loess *indicator.LoessIndicator
}

func NewSingleLoesslineStrategyFactory() model.Strategy {
	return &SingleLOESSStrategy{}
}

func (es *SingleLOESSStrategy) Name() string {
	return "SingleLOESSStrategy"
}

func (es *SingleLOESSStrategy) Init(ec *model.ContractStrategyContext) {
	//es.loess5 = indicator.NewLOESSIndicator(ec.Kline, 2, 2, 5)
	//es.loess5_loess3 = indicator.NewLOESSIndicator(es.loess5, 3, 5, 10)

	//es.loess5_loess2 = indicator.NewLOESSIndicator(es.loess5, 2, 4, 6)
	es.loess = indicator.NewLoessIndicator(ec.Kline, 3)
}

func (es *SingleLOESSStrategy) OnBar(ctx *model.ContractStrategyContext, ts int64) *model.StrategyResult {
	//s := 0.01
	currentKNode, err := ctx.Kline.GetByTs(ts)
	if err != nil || currentKNode == nil || currentKNode.GetValue() == 0 {
		return nil
	}
	curPrice := currentKNode.GetValue()

	loessValue := 0.0
	v, _ := es.loess.GetByTs(ts)
	if v != nil {
		loessValue = v.GetValue()
	}

	gap := 0.001
	if curPrice < loessValue && (loessValue-curPrice)/loessValue > gap {
		reason := fmt.Sprintf("(loess(%v)-cur(%v))/loess(%v)=%v>gap(%v)", loessValue, curPrice, loessValue, (loessValue-curPrice)/loessValue, gap)
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice), reason)
	} else if curPrice > loessValue && (curPrice-loessValue)/curPrice > gap {
		reason := fmt.Sprintf("(cur(%v)-loess(%v))/cur(%v)=%v>gap(%v)", curPrice, loessValue, curPrice, (curPrice-loessValue)/curPrice, gap)
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice), reason)
	} else {
		return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice), "not_long_or_short")
	}
}
