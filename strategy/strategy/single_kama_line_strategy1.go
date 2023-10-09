package strategy

// 自定义的均线策略 ， 使用sma1， kama5, kamaFast
// 使用kama5+kamaFast 入场，使用 sma1 出场

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type SingleKAMAStrategy struct {
	kama     *indicator.KAMAIndicator
	kamaKama *indicator.KAMAIndicator
	slop     *indicator.SlopIndicator
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
	es.kama = indicator.NewKAMAIndicator(ec.Kline, 2, 5, 4)
	es.kamaKama = indicator.NewKAMAIndicator(es.kama, 2, 4, 6)
	es.slop = indicator.NewSlopIndicator(es.kamaKama, 2)
}

func (es *SingleKAMAStrategy) OnBar(ctx *model.ContractStrategyContext, ts int64) *model.StrategyResult {
	//s := 0.01
	slopPeriod := 10
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

	kamaValue := es.kama.GetCurrentFloat(ts)

	gap := 0.002
	if curPrice < kamaValue && (kamaValue-curPrice)/kamaValue > gap {
		reason := fmt.Sprintf("(kama %v -  cur %v)/kama %v =  gap %v > gap", kamaValue, curPrice, kamaValue, (kamaValue-curPrice)/kamaValue)
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice), reason)
	} else if curPrice > kamaValue && (curPrice-kamaValue)/curPrice > gap {
		reason := fmt.Sprintf("(cur %v -  kama %v)/cur %v =  gap %v > gap", curPrice, kamaValue, curPrice, (curPrice-kamaValue)/curPrice)
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice), reason)
	} else {
		return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice), "")
	}
	//
	//crossOver, _ := es.crossover.GetByTs(ts)
	//if crossOver != nil {
	//	return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice))
	//}
	//crossUnder, _ := es.crossunder.GetByTs(ts)
	//if crossUnder != nil {
	//	return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice))
	//
	//}
	return nil
	//if es.crossover.GetByTs(ts) utils.SliceFloatGt(slops, s) {
	//	return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice))
	//
	//} else if utils.SliceFloatLt(slops, -1*s) {
	//	return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice))
	//
	//} else {
	//	return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice))
	//}

	//if utils.SliceFloatGt(slops, s) {
	//	return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice))
	//
	//} else if utils.SliceFloatLt(slops, -1*s) {
	//	return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice))
	//
	//} else {
	//	return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice))
	//}
}
