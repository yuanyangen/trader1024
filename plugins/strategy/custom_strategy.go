package strategy

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/local_broker"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

//自定义的均线策略 ， 使用sma1， kama5, kamaFast
// 使用kama5+kamaFast 入场，使用 sma1 出场

type CustomStrategy2 struct {
	openKama  model.Indicator
	closeKama model.Indicator
	avgKama   model.Indicator
	closeSlop model.Indicator
	openSlop  model.Indicator
	avgSlop   model.Indicator
	powerSMA  model.Indicator
	//aocr      model.Indicator
}

func NewCustomStrategy2Factory() model.Strategy {
	return &CustomStrategy2{}
}

// 1. 找到合适的指标，降低掉波动的影响
// 2. 仓位管理： 如果没有合适的仓位，则空仓。 如果仓位没有预期的盈利， 也空仓。
func (es *CustomStrategy2) Name() string {
	return "CustomStrategy2"
}

func (es *CustomStrategy2) Init(ec *model.TradeObjectEngineContext) {
	es.openKama = indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))
	es.closeKama = indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))
	es.avgKama = indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceAVG))
	es.closeSlop = indicator.NewSlopIndicator(2, es.closeKama)
	es.openSlop = indicator.NewSlopIndicator(2, es.openKama)
	es.avgSlop = indicator.NewSlopIndicator(2, es.avgKama)
	es.powerSMA = indicator.NewSMAIndicator(4, indicator.NewRawKlineIndicator(indicator.DataSourceCloseSubOpen))

	//es.closeSlop = indicator.NewSlopIndicator(2, indicator.NewKAMAIndicator(5, indicator.DataSourceClose))
	//es.aocr = indicator.NewAOCRIndicator(5)

	ec.Kline.Indicators = []model.Indicator{
		//es.openKama,
		//es.closeKama,
		//es.avgKama,
		//es.powerSMA,
		indicator.NewRawKlineIndicator(indicator.DataSourceClose),
		indicator.NewRawKlineIndicator(indicator.DataSourceAVG),
		//indicator.NewRawKlineIndicator(indicator.DataSourceFullAVG),

		//indicator.NewRawKlineIndicator(indicator.DataSourceOpen),
		//indicator.NewRawKlineIndicator(indicator.DataSourceHigh),
		//indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
		//indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		//indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(20, 	indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewZDMAIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),

		//indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

		//indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		//indicator.NewZDMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewZDMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

		indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		//indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		//indicator.NewCCIIndicator(40),
		//indicator.NewCCIIndicator(20),
		//indicator.NewMFIIndicator(40),
		//indicator.NewMFIIndicator(20),
		//indicator.NewMFIIndicator(10),

		//indicator.NewMFIIndicator(5),
		//indicator.NewBLIndicator(2470),
		//indicator.NewBLIndicator(2450),

		//indicator.NewSMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewSMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),

		//indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		//indicator.NewJMAIndicator(0.15, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewJMAIndicator(0.15, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewJMAIndicator(0.2, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewJMAIndicator(0.2, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewJMAIndicator(0.3, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewJMAIndicator(0.5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewJMAIndicator(0.6, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewJMAIndicator(0.4, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//
		//indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),

		//indicator.NewSMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewSMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewSMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewSMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),

		//indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewLRIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),

		//indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		//indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),

		//indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		//indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		//indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),

		//indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewLRIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		//indicator.NewLRIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		//indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		//indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		//indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		//indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		//indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

		//indicator.NewKAMAIndicator(2, es.openKama.Name(), es.openKama.GetValueFromKNode),
		//indicator.NewKAMAIndicator(2, es.openKama.Name(), es.openKama.GetValueFromKNode),
		//indicator.NewKAMAIndicator(2, es.avgKama.Name(), es.avgKama.GetValueFromKNode),

		//indicator.NewKAMAIndicator(5, es.openKama.Name(), es.openKama.GetValueFromKNode),
		//indicator.NewKAMAIndicator(5, es.openKama.Name(), es.openKama.GetValueFromKNode),
		//indicator.NewKAMAIndicator(5, es.avgKama.Name(), es.avgKama.GetValueFromKNode),
		//indicator.NewZDMAIndicator(5, "close", indicator.DataFromClose),
		//indicator.NewZDMAIndicator(5, "open", indicator.DataFromOpen),
		//indicator.NewSlopIndicator(2, indicator.NewZDMAIndicator(20, "close", indicator.DataFromClose)),
		//indicator.NewSlopIndicator(2, indicator.NewZDMAIndicator(20, "open", indicator.DataFromOpen)),
		//indicator.NewSlopIndicator(2, indicator.NewZDMAIndicator(20, "avg", indicator.DataFromAVG)),

		//indicator.NewSMAIndicator(20, indicator.DataSourceClose),
		//es.closeSlop,
	}
}

func (es *CustomStrategy2) OnBar(ctx *model.TradeObjectEngineContext) *model.StrategyResult {
	ts := ctx.Ts
	currentKNode, err := ctx.Kline.GetNodeByTs(ts)
	if err != nil || currentKNode == nil {
		return nil
	}
	closePrice := es.closeKama.GetValueFromKNode(currentKNode)
	openPrice := es.openKama.GetValueFromKNode(currentKNode)

	if closePrice == 0 || openPrice == 0 {
		return nil
	}

	curPrice := currentKNode.Close
	lastPosition := local_broker.GetLocalBroker().GetPositionsByContract(ctx.TradeObjecct).GetLastPair()

	if lastPosition != nil && !lastPosition.Clear {
		//if lastPosition.Type == model.PositionTypeLong {
		//	targetPrice := lastPosition.Buy.Price.InexactFloat64()
		//	if curPrice <= targetPrice {
		//		return model.NewStrategyResult(model.PositionTypeClearLast, decimal.NewFromFloat(curPrice), fmt.Sprintf("quick_clear_last_long_%v<%v", curPrice, targetPrice))
		//	}
		//}
		//if lastPosition.Type == model.PositionTypeShort {
		//	targetPrice := lastPosition.Sell.Price.InexactFloat64()
		//	if curPrice >= targetPrice {
		//		return model.NewStrategyResult(model.PositionTypeClearLast, decimal.NewFromFloat(curPrice), fmt.Sprintf("quick_clear_last_short_%v>%v", curPrice, targetPrice))
		//	}
		//}
	}

	inOffset := 0.0015
	outOffset := 0.000
	currentPower := es.powerSMA.GetValueFromKNode(currentKNode)
	avgSlop := es.avgSlop.GetValueFromKNode(currentKNode)
	if closePrice > openPrice && avgSlop >= inOffset && currentPower > 0 {
		return model.NewStrategyResult(model.PositionTypeLong, decimal.NewFromFloat(curPrice), "trend_long")
	} else if closePrice < openPrice && avgSlop <= -1*inOffset && currentPower < 0 {
		return model.NewStrategyResult(model.PositionTypeShort, decimal.NewFromFloat(curPrice), "trend_short")
	}
	if lastPosition != nil && !lastPosition.Clear {
		if lastPosition.Type == model.PositionTypeLong {
			if avgSlop < outOffset {
				reason := fmt.Sprintf("trend_miss: closePrice:%v open:%v", closePrice, openPrice)
				return model.NewStrategyResult(model.PositionTypeClear, decimal.NewFromFloat(curPrice), reason)
			} else if currentPower < 0 {
				reason := fmt.Sprintf("long power down: closePrice:%v open:%v power:%v ", closePrice, openPrice, currentPower)
				return model.NewStrategyResult(model.PositionTypeClear, decimal.NewFromFloat(curPrice), reason)
			}
		} else if lastPosition.Type == model.PositionTypeShort {
			if avgSlop >= -1*outOffset {
				reason := fmt.Sprintf("trend_miss: closePrice:%v open:%v", closePrice, openPrice)
				return model.NewStrategyResult(model.PositionTypeClear, decimal.NewFromFloat(curPrice), reason)
			} else if currentPower > 0 {
				reason := fmt.Sprintf("short power down: closePrice:%v open:%v power:%v ", closePrice, openPrice, currentPower)
				return model.NewStrategyResult(model.PositionTypeClear, decimal.NewFromFloat(curPrice), reason)
			}
		}
	}
	return nil
}

func tmp() {
	//if closePrice > openPrice && closePrice-openPrice > 0.3*aocr {
	//	reason := fmt.Sprintf("long: %.3f>%.3f && %.3f ", closePrice, openPrice, aocr)
	//	return model.NewStrategyResult(model.PositionTypeLong, decimal.NewFromFloat(curPrice), reason)
	//} else if closePrice < openPrice && closePrice-openPrice < -0.3*aocr {
	//	reason := fmt.Sprintf("short: %.3f<%.3f && %.3f ", closePrice, openPrice, aocr)
	//	return model.NewStrategyResult(model.PositionTypeShort, decimal.NewFromFloat(curPrice), reason)
	//}
	{

		//if closePrice > openPrice && closeKamaSlop >= 0.01 && closePrice-openPrice > 0.3*aocr {
		//	reason := fmt.Sprintf("long: %.3f>%.3f && %.3f >= 0.01 aocr %v", closePrice, openPrice, closeKamaSlop, aocr)
		//	return model.NewStrategyResult(model.PositionTypeLong, decimal.NewFromFloat(curPrice), reason)
		//} else if closePrice < openPrice && closeKamaSlop <= -0.01 && closePrice-openPrice < -0.3*aocr {
		//	reason := fmt.Sprintf("short: %.3f<%.3f && %.3f <= -0.01 aocr%v ", closePrice, openPrice, closeKamaSlop, aocr)
		//	return model.NewStrategyResult(model.PositionTypeShort, decimal.NewFromFloat(curPrice), reason)
		//}
	}
	{
		//if currentKNode.Open > openPrice && currentKNode.Close > closePrice {
		//	reason := fmt.Sprintf("long: %.3f>%.3f  ", closePrice, openPrice)
		//	return model.NewStrategyResult(model.PositionTypeLong, decimal.NewFromFloat(curPrice), reason)
		//} else if currentKNode.Close < closePrice && currentKNode.Open < openPrice {
		//	reason := fmt.Sprintf("short: %.3f<%.3f  ", closePrice, openPrice)
		//	return model.NewStrategyResult(model.PositionTypeShort, decimal.NewFromFloat(curPrice), reason)
		//}
	}

}
