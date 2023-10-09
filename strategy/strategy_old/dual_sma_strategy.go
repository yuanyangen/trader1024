package strategy_old

//
//import (
//	"github.com/shopspring/decimal"
//	"github.com/yuanyangen/trader1024/engine/model"
//	"github.com/yuanyangen/trader1024/engine/utils"
//	"github.com/yuanyangen/trader1024/strategy_old/indicator"
//)
//
//type DualSMAStrategy struct {
//	slowSMA    *indicator.SMAIndicator
//	fastSMA    *indicator.SMAIndicator
//	crossover  *indicator.CrossOverIndicator
//	crossunder *indicator.CrossUnderIndicator
//}
//
//func NewDualSMAStrategyFactory() model.Strategy {
//	return &DualSMAStrategy{}
//}
//
////func (es *DualSMAStrategy) Indicators() []indicator.MarketIndicator {
////	return []indicator.MarketIndicator{
////		es.sma10,
////		es.sma5,
////	}
////}
//
//func (es *DualSMAStrategy) CNName() string {
//	return "DualSMA"
//}
//
//func (es *DualSMAStrategy) Init(ec *model.MarketStrategyContext) {
//	es.slowSMA = indicator.NewSMAIndicator(ec.DailyData.Kline, 10)
//	es.fastSMA = indicator.NewSMAIndicator(ec.DailyData.Kline, 5)
//	es.crossover = indicator.NewCrossOverIndicator(ec.DailyData.Kline, es.fastSMA.SMALine, es.slowSMA.SMALine)
//	es.crossunder = indicator.NewCrossUnderIndicator(ec.DailyData.Kline, es.fastSMA.SMALine, es.slowSMA.SMALine)
//}
//
//func (es *DualSMAStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) []*model.StrategyResult {
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
