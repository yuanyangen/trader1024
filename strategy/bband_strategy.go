package strategy

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
)

type BBandStrategy struct {
	sma    *indicator.BBANDIndicator
	loaded bool // 只有
	//sma5 *indicator.SMAIndicator
	//crossover  *indicator.CrossOverIndicator
	//crossunder *indicator.CrossUnderIndicator
}

func NewBBandStrategy() Strategy {
	return &BBandStrategy{}
}

func (es *BBandStrategy) Name() string {
	return "BBand"
}

func (es *BBandStrategy) Init(ec *MarketStrategyContext) {
	es.sma = indicator.NewBBANDIndicator(ec.DailyData.Kline, 5)
}

func (es *BBandStrategy) OnBar(ctx *MarketStrategyContext, ts int64) []*model.StrategyResult {
	currentKValue, err := ctx.DailyData.Kline.GetKnodeByTs(ts)
	if err != nil {
		return nil
	}
	low := es.sma.GetLowFloat(ts)
	mid := es.sma.GetMidFloat(ts)
	upper := es.sma.GetUpperFloat(ts)
	if low == 0 || mid == 0 || upper == 0 {
		return nil
	}
	position := account.GetAccount().GetPositionByMarket(ctx.Market.MarketId)
	curPrice := (currentKValue.Open + currentKValue.Close) / 2
	if curPrice > upper {
		if position.IsEmpty() {
			return []*model.StrategyResult{
				NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(curPrice)),
			}
		} else if position.Count.LessThan(decimal.Zero) {
			return []*model.StrategyResult{
				NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
				NewStrategyResult(model.StrategyCmdBuy, decimal.NewFromFloat(curPrice)),
			}
		}
	} else if curPrice < low {
		if position.IsEmpty() && es.loaded {
			return []*model.StrategyResult{
				NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(curPrice)),
			}
		} else if position.Count.GreaterThan(decimal.Zero) {
			return []*model.StrategyResult{
				NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
				NewStrategyResult(model.StrategyCmdSell, decimal.NewFromFloat(curPrice)),
			}
		}
	} else if curPrice < mid && curPrice > low {
		if position.Count.GreaterThan(decimal.Zero) {
			return []*model.StrategyResult{
				NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
			}
		}
	} else {
		if position.Count.LessThan(decimal.Zero) {
			return []*model.StrategyResult{
				NewStrategyResult(model.StrategyCmdClean, decimal.NewFromFloat(curPrice)),
			}
		}
	}

	return nil
}
