package strategy

//
//// 自定义的均线策略 ， 使用sma1， kama5, kama20
//// 使用kama5+kama20 入场，使用 sma1 出场
//
//import (
//	"github.com/shopspring/decimal"
//	"github.com/yuanyangen/trader1024/engine/account"
//	"github.com/yuanyangen/trader1024/engine/model"
//	"github.com/yuanyangen/trader1024/engine/utils"
//	"github.com/yuanyangen/trader1024/strategy/indicator"
//)
//
//type CustomKAMAStrategy2 struct {
//	kama10 *indicator.KAMAIndicator
//	//kama5      *indicator.KAMAIndicator
//	kama2      *indicator.KAMAIndicator
//	crossover  *indicator.CrossOverIndicator
//	crossunder *indicator.CrossUnderIndicator
//	loaded     bool // 只有
//}
//
//func NewCustomLAMAStrategy2Factory() model.Strategy {
//	return &CustomKAMAStrategy2{}
//}
//
//func (es *CustomKAMAStrategy2) CNName() string {
//	return "CustomKAMAStrategy2"
//}
//
//func (es *CustomKAMAStrategy2) Init(ec *model.MarketStrategyContext) {
//	es.kama10 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 8, 2, 30)
//	//es.kama5 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 5, 2, 30)
//	es.kama2 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 2, 2, 30)
//	es.crossover = indicator.NewCrossOverIndicator(ec.DailyData.Kline, es.kama2.KAMALine, es.kama10.KAMALine)
//	es.crossunder = indicator.NewCrossUnderIndicator(ec.DailyData.Kline, es.kama2.KAMALine, es.kama10.KAMALine)
//}
//
//func (es *CustomKAMAStrategy2) OnBar(ctx *model.MarketStrategyContext, ts int64) []*model.StrategyResult {
//	currentKValue := model.NewKnodeFromAny(ctx.DailyData.Kline.GetByTs(ts))
//	if currentKValue == nil {
//		return nil
//	}
//	if es.kama2.GetCurrentFloat(ts) == 0 || es.kama10.GetCurrentFloat(ts) == 0 {
//		return nil
//	}
//	position := account.GetAccount().GetPositionByMarket(ctx.Contract.ContractId)
//	curPrice := (currentKValue.Open + currentKValue.Close) / 2
//	data := []*model.StrategyResult{}
//
//	if utils.AnyToBool(es.crossover.GetByTs(ts)) || utils.AnyToBool(es.crossunder.GetByTs(ts)) {
//		es.loaded = true
//		data = append(data, model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)))
//	}
//
//	if es.long(es.kama2.GetCurrentFloat(ts), es.kama10.GetCurrentFloat(ts)) {
//		if position.IsEmpty() && es.loaded {
//			data = append(data, model.NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(curPrice)))
//		}
//	} else if es.short(es.kama2.GetCurrentFloat(ts), es.kama10.GetCurrentFloat(ts)) {
//		if position.IsEmpty() && es.loaded {
//			data = append(data, model.NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(curPrice)))
//		}
//	}
//
//	return data
//}
//
//func (es *CustomKAMAStrategy2) long(fast, slow float64) bool {
//	return fast > slow && (fast-slow)/slow > 0.001
//}
//func (es *CustomKAMAStrategy2) short(fast, slow float64) bool {
//	return fast < slow && (slow-fast)/fast > 0.001
//}
