package strategy_new

// 自定义的均线策略 ， 使用sma1， kama5, kama20
// 使用kama5+kama20 入场，使用 sma1 出场

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type CustomKAMASlopStrategy struct {
	kama5      *indicator.KAMAIndicator
	kama5_kama *indicator.KAMAIndicator
	slop       model.MarketIndicator
}

func NewCustomKAMASlopStrategyFactory() model.Strategy {
	return &CustomKAMASlopStrategy{}
}

func (es *CustomKAMASlopStrategy) Name() string {
	return "CustomKAMASlopStrategy"
}

func (es *CustomKAMASlopStrategy) Init(ec *model.ContractStrategyContext) {
	es.kama5 = indicator.NewKAMAIndicator(ec.Kline, 10, 5, 10)
	es.kama5_kama = indicator.NewKAMAIndicator(es.kama5, 10, 5, 10)
	es.slop = indicator.NewSlopIndicator(es.kama5, 2)
}

func (es *CustomKAMASlopStrategy) OnBar(ctx *model.ContractStrategyContext, ts int64) *model.StrategyResult {
	s := 0.001
	slopPeriod := 5
	currentKNode, err := ctx.Kline.GetByTs(ts)
	if err != nil || currentKNode == nil || currentKNode.GetValue() == 0 {
		return nil
	}
	curPrice := currentKNode.GetValue()
	slopI, _ := es.slop.GetLastByTsAndCount(ts, int64(slopPeriod))
	slops := utils.DataNodeSliceToFloat(slopI)
	if len(slops) == 0 {
		return nil
	}
	if utils.SliceFloatGt(slops, s) {
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice))

	} else if utils.SliceFloatLt(slops, -1*s) {
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice))

	} else {
		return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice))
	}
}
