package strategy

//
//// 自定义的均线策略 ， 使用sma1， kama5, kama20
//// 使用kama5+kama20 入场，使用 sma1 出场
//
//import (
//	"github.com/shopspring/decimal"
//	"github.com/yuanyangen/trader1024/engine/model"
//	"github.com/yuanyangen/trader1024/engine/utils"
//	"github.com/yuanyangen/trader1024/strategy/indicator"
//)
//
//type DualKAMAStrategy struct {
//	kama10     *indicator.KAMAIndicator
//	kama2      *indicator.KAMAIndicator
//	crossover  *indicator.CrossOverIndicator
//	crossunder *indicator.CrossUnderIndicator
//	loaded     bool // 只有
//}
//
//func NewDualKAMAStrategyFactory() model.Strategy {
//	return &DualKAMAStrategy{}
//}
//
//func (es *DualKAMAStrategy) CNName() string {
//	return "DualKAMAStrategy"
//}
//
//func (es *DualKAMAStrategy) Init(ec *model.MarketStrategyContext) {
//	es.kama10 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 10, 2, 30)
//	es.kama2 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 2, 2, 30)
//	es.crossover = indicator.NewCrossOverIndicator(ec.DailyData.Kline, es.kama2.KAMALine, es.kama10.KAMALine)
//	es.crossunder = indicator.NewCrossUnderIndicator(ec.DailyData.Kline, es.kama2.KAMALine, es.kama10.KAMALine)
//}
//
//func (es *DualKAMAStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) []*model.StrategyResult {
//	over := es.crossover.GetByTs(ts)
//	under := es.crossunder.GetByTs(ts)
//	currentKValue := model.NewKnodeFromAny(ctx.DailyData.Kline.GetByTs(ts))
//	if currentKValue == nil {
//		return nil
//	}
//	if utils.AnyToBool(over) {
//		return []*model.StrategyResult{
//			model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
//			model.NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(currentKValue.Close)),
//		}
//	}
//	if utils.AnyToBool(under) {
//		return []*model.StrategyResult{
//			model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
//			model.NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(currentKValue.Close)),
//		}
//	}
//	return nil
//}
