package strategy

// 自定义的均线策略 ， 使用sma1， sma5, smaFast
// 使用sma5+smaFast 入场，使用 sma1 出场

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type SingleSMAStrategy struct {
	smaQuick *indicator.SMAIndicator
	smaSlow  *indicator.SMAIndicator
	slop     *indicator.SlopIndicator
}

func NewSingleSMAlineStrategyFactory() model.Strategy {
	return &SingleSMAStrategy{}
}

func (es *SingleSMAStrategy) Name() string {
	return "SingleSMAStrategy"
}

func (es *SingleSMAStrategy) Init(ec *model.ContractStrategyContext) {
	//es.sma5 = indicator.NewsmaIndicator(ec.Kline, 2, 2, 5)
	//es.sma5_sma3 = indicator.NewsmaIndicator(es.sma5, 3, 5, 10)
	//es.sma5_sma2 = indicator.NewsmaIndicator(es.sma5, 2, 4, 6)
	es.smaQuick = indicator.NewSMAIndicator(ec.Kline, 1)
	es.smaSlow = indicator.NewSMAIndicator(ec.Kline, 5)
	es.slop = indicator.NewSlopIndicator(es.smaSlow, 2)
}

func (es *SingleSMAStrategy) OnBar(ctx *model.ContractStrategyContext, ts int64) *model.StrategyResult {
	//s := 0.01
	slopPeriod := 2
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
	var smaValue = 0.0
	v, _ := es.smaSlow.GetByTs(ts)
	if v != nil {
		smaValue = v.GetValue()
	}
	if smaValue == 0 {
		return nil
	}

	gap := 0.001
	if curPrice < smaValue && (smaValue-curPrice)/smaValue > gap {
		reason := fmt.Sprintf("(sma(%v)-cur(%v))/sma(%v)=%v>gap(%v)", smaValue, curPrice, smaValue, (smaValue-curPrice)/smaValue, gap)
		return model.NewStrategyResult(model.StrategyOutShort, decimal.NewFromFloat(curPrice), reason)
	} else if curPrice > smaValue && (curPrice-smaValue)/curPrice > gap {
		reason := fmt.Sprintf("(cur(%v)-sma(%v))/cur(%v)=%v>gap(%v)", curPrice, smaValue, curPrice, (curPrice-smaValue)/curPrice, gap)
		return model.NewStrategyResult(model.StrategyOutLong, decimal.NewFromFloat(curPrice), reason)
	} else {
		return model.NewStrategyResult(model.StrategyOutVolatility, decimal.NewFromFloat(curPrice), "not_long_or_short")
	}
}
