package strategy

//
//import (
//	"github.com/shopspring/decimal"
//	"github.com/yuanyangen/trader1024/engine/account"
//	"github.com/yuanyangen/trader1024/engine/model"
//	"github.com/yuanyangen/trader1024/strategy/indicator"
//)
//
//type SingleSMAStrategy struct {
//	sma *indicator.SMAIndicator
//	//sma5 *indicator.SMAIndicator
//	//crossover  *indicator.CrossOverIndicator
//	//crossunder *indicator.CrossUnderIndicator
//}
//
//func NewSingleSMAStrategy() model.Strategy {
//	return &SingleSMAStrategy{}
//}
//
//func (es *SingleSMAStrategy) CNName() string {
//	return "SingleSMA"
//}
//
//func (es *SingleSMAStrategy) Init(ec *model.MarketStrategyContext) {
//	es.sma = indicator.NewSMAIndicator(ec.DailyData.Kline, 5)
//}
//
//func (es *SingleSMAStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) []*model.StrategyResult {
//	currentKValue := model.NewKnodeFromAny(ctx.DailyData.Kline.GetByTs(ts))
//	if currentKValue == nil {
//		return nil
//	}
//	sma, ok := es.sma.GetByTs(ts).(float64)
//	if !ok {
//		return nil
//	}
//	if sma > currentKValue.Close && account.GetAccount().GetPositionByMarket(ctx.Contract.MarketId).IsEmpty() {
//		return []*model.StrategyResult{
//			model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
//			model.NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(currentKValue.Close)),
//		}
//	}
//	if sma < currentKValue.Close && account.GetAccount().GetPositionByMarket(ctx.Contract.MarketId).IsEmpty() {
//		return []*model.StrategyResult{
//			model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
//			model.NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(currentKValue.Close)),
//		}
//	}
//	//if under {
//	//	return []*model.StrategyResult{
//	//		NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(currentKValue.Close)),
//	//		NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(currentKValue.Close)),
//	//	}
//	//}
//	return nil
//}
