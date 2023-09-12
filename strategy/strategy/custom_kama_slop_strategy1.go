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
//type CustomKAMASlopStrategy struct {
//	kama2      *indicator.KAMAIndicator
//	slop       model.MarketIndicator
//	volumnSlop model.MarketIndicator
//	//rsi        model.MarketIndicator
//	continous model.MarketIndicator
//	//continousKama model.MarketIndicator
//	//continousSma  model.MarketIndicator
//	loaded bool // 只有
//}
//
//func NewCustomLAMASlopStrategyFactory() model.Strategy {
//	return &CustomKAMASlopStrategy{}
//}
//
//func (es *CustomKAMASlopStrategy) CNName() string {
//	return "CustomKAMASlopStrategy"
//}
//
//func (es *CustomKAMASlopStrategy) Init(ec *model.MarketStrategyContext) {
//	es.kama2 = indicator.NewKAMAIndicator(ec.DailyData.Kline, 2, 2, 30)
//	es.slop = indicator.NewSlopIndicator(es.kama2, 2)
//	//es.rsi = indicator.NewRSIIndicator(ec.DailyData.Kline, 5)
//	es.continous = indicator.NewContinousIndicator(ec.DailyData.Kline)
//	//es.continousKama = indicator.NewLineKAMAIndicator(es.continous, 5, 2, 30)
//	es.volumnSlop = indicator.NewVolumnSlopIndicator(ec.DailyData.Kline)
//	//es.continousSma = indicator.NewLineSmaIndicator(es.continous, 5)
//}
//
//func (es *CustomKAMASlopStrategy) OnBar(ctx *model.MarketStrategyContext, ts int64) []*model.StrategyResult {
//
//	s := 0.005
//	maxS := 0.02
//	slopPeriod := 2
//	continousV := 0.015
//	currentKValue := model.NewKnodeFromAny(ctx.DailyData.Kline.GetByTs(ts))
//	if currentKValue == nil {
//		return nil
//	}
//	slopI, _ := es.slop.GetLastByTsAndCount(ts, int64(slopPeriod))
//	slops := utils.AnySliceToFloat(slopI)
//	if len(slops) == 0 {
//		return nil
//	}
//	position := account.GetAccount().GetPositionByMarket(ctx.Contract.MarketId)
//	curPrice := (currentKValue.Open + currentKValue.Close) / 2
//	continous := es.continous.GetByTs(ts)
//	v, ok := utils.AnyToFloat(continous)
//	if ok && v > continousV {
//		//return []*model.StrategyResult{
//		//	model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
//		//}
//	}
//	if position.IsEmpty() {
//		if utils.SliceFloatGt(slops, s) {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(curPrice)),
//			}
//		} else if utils.SliceFloatLt(slops, -1*s) {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(curPrice)),
//			}
//		}
//	} else if position.Count.GreaterThan(decimal.Zero) {
//		//if math.Abs(v) > continousV {
//		//	return nil
//		//}
//		if utils.SliceFloatLt(slops, 0) || slops[len(slops)-1] < -1*maxS {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
//			}
//		}
//	} else if position.Count.LessThan(decimal.Zero) {
//		//if math.Abs(v) > continousV {
//		//	return nil
//		//}
//		if utils.SliceFloatGt(slops, 0) || slops[len(slops)-1] > maxS {
//			return []*model.StrategyResult{
//				model.NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
//			}
//		}
//	}
//	return nil
//}
