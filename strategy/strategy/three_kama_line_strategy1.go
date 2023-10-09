package strategy

// 自定义的均线策略 ， 使用sma1， kama5, kamaFast
// 使用kama5+kamaFast 入场，使用 sma1 出场

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type ThreeKAMAStrategy struct {
	//kama5            *indicator.KAMAIndicator
	//kama5_kama2      *indicator.KAMAIndicator
	//kama5_kama3      *indicator.KAMAIndicator
	kamaSlow         *indicator.KAMAIndicator
	kamaSlowKama     *indicator.KAMAIndicator
	kamaFast         *indicator.KAMAIndicator
	kamaFastKama     *indicator.KAMAIndicator
	kamaFastKamaKama *indicator.KAMAIndicator
	//crossover    *indicator.CrossOverIndicator
	//crossunder   *indicator.CrossUnderIndicator
	slop *indicator.SlopIndicator
}

func NewThreeKAMAlineStrategyFactory() model.Strategy {
	return &ThreeKAMAStrategy{}
}

func (es *ThreeKAMAStrategy) Name() string {
	return "SingleKAMAStrategy"
}

func (es *ThreeKAMAStrategy) Init(ec *model.ContractStrategyContext) {
	//es.kama5 = indicator.NewKAMAIndicator(ec.Kline, 2, 2, 5)
	//es.kama5_kama3 = indicator.NewKAMAIndicator(es.kama5, 3, 5, 10)
	//es.kama5_kama2 = indicator.NewKAMAIndicator(es.kama5, 2, 4, 6)
	es.kamaFast = indicator.NewKAMAIndicator(ec.Kline, 2, 5, 2)
	es.kamaSlow = indicator.NewKAMAIndicator(ec.Kline, 4, 5, 4)

	es.kamaFastKama = indicator.NewKAMAIndicator(es.kamaFast, 2, 4, 6)
	es.kamaSlowKama = indicator.NewKAMAIndicator(es.kamaSlow, 2, 4, 6)
	es.kamaFastKamaKama = indicator.NewKAMAIndicator(es.kamaFastKama, 2, 4, 6)
	//es.crossover = indicator.NewCrossOverIndicator(ec.Kline, es.kamaFast.KAMALine, es.kamaFastKama.KAMALine)
	//es.crossunder = indicator.NewCrossUnderIndicator(ec.Kline, es.kamaFast.KAMALine, es.kamaFastKama.KAMALine)

	es.slop = indicator.NewSlopIndicator(es.kamaSlowKama, 2)
}

func (es *ThreeKAMAStrategy) OnBar(ctx *model.ContractStrategyContext, ts int64) *model.StrategyResult {
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

	fast := es.kamaFast.GetCurrentFloat(ts)
	medium := es.kamaFastKama.GetCurrentFloat(ts)
	slow := es.kamaFastKamaKama.GetCurrentFloat(ts)
	slop := es.slop.GetCurrentFloat(ts)

	hitSlop := slop > 0.001 || slop < -0.001
	//gap := 0.005
	gap := 0.001
	if fast < medium && medium < slow && (slow-fast)/slow > gap && hitSlop {
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice), "")
	} else if fast > medium && medium > slow && (fast-slow)/fast > gap && hitSlop {
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice), "")
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
