package strategy

//
//import (
//	"github.com/shopspring/decimal"
//	"github.com/yuanyangen/trader1024/engine/account"
//	"github.com/yuanyangen/trader1024/engine/model"
//	"github.com/yuanyangen/trader1024/strategy/indicator"
//)
//
//type BBandStrategy struct {
//	sma    *indicator.BBANDIndicator
//	loaded bool // 只有
//}
//
//func NewBBandStrategy() model.Strategy {
//	return &BBandStrategy{}
//}
//
//func (es *BBandStrategy) Name() string {
//	return "BBand"
//}
//
//func (es *BBandStrategy) Init(ec *model.MarketStrategyContext) {
//	es.sma = indicator.NewBBANDIndicator(ec.DailyData.Kline, 5)
//}
//
//func (es *BBandStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) []*model.StrategyResult {
//	currentKValue := model.NewKnodeFromAny(ctx.DailyData.Kline.GetByTs(ts))
//	if currentKValue == nil {
//		return nil
//	}
//
//	low := es.sma.GetLowFloat(ts)
//	mid := es.sma.GetMidFloat(ts)
//	upper := es.sma.GetUpperFloat(ts)
//	if low == 0 || mid == 0 || upper == 0 {
//		return nil
//	}
//	position := account.GetAccount().GetPositionByMarket(ctx.Market.MarketId)
//	curPrice := (currentKValue.Open + currentKValue.Close) / 2
//	if curPrice > upper {
//		if position.IsEmpty() {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(curPrice)),
//			}
//		} else if position.Count.LessThan(decimal.Zero) {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
//				model.NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(curPrice)),
//			}
//		}
//	} else if curPrice < low {
//		if position.IsEmpty() && es.loaded {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(curPrice)),
//			}
//		} else if position.Count.GreaterThan(decimal.Zero) {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
//				model.NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(curPrice)),
//			}
//		}
//	} else if curPrice < mid && curPrice > low {
//		if position.Count.GreaterThan(decimal.Zero) {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
//			}
//		}
//	} else {
//		if position.Count.LessThan(decimal.Zero) {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
//			}
//		}
//	}
//
//	return nil
//}
