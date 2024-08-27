package strategy

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
)

//自定义的均线策略 ， 使用sma1， kama5, kamaFast
// 使用kama5+kamaFast 入场，使用 sma1 出场

type CustomStrategy2 struct {
}

func NewCustomStrategy2Factory() model.Strategy {
	return &CustomStrategy2{}
}

func (es *CustomStrategy2) Name() string {
	return "CustomStrategy2"
}

func (es *CustomStrategy2) Init(ec *model.ContractEngineContext) {
}

func (es *CustomStrategy2) OnBar(ctx *model.ContractEngineContext) *model.StrategyResult {
	ts := ctx.Ts
	currentKNode, err := ctx.Kline.GetNodeByTs(ts)
	if err != nil || currentKNode == nil {
		return nil
	}
	lastDayKNode, err := ctx.Kline.GetNodeByTs(ts)
	if err != nil || lastDayKNode == nil {
		return nil
	}
	if currentKNode.SmaData == nil {
		return nil
	}
	sma10 := currentKNode.SmaData[10]
	sma20 := currentKNode.SmaData[20]
	if sma20 == 0 || sma10 == 0 {
		return nil
	}
	fast := sma10
	slow := sma20
	curPrice := currentKNode.Close

	if fast > slow && (fast-slow)/slow > 0.001 {
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice), "")
	} else if fast < slow && (slow-fast)/fast > 0.001 {
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice), "")
	}
	return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice), "")
}
