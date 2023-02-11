package strategy

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
)

type MarketStrategyContext struct {
	Market    *model.Market
	DailyData *indicator.DailyIndicators
}

func NewStrategyResult(Cmd model.StrategyCmd, price decimal.Decimal) *model.StrategyResult {
	return &model.StrategyResult{
		Cmd:   Cmd,
		Price: price,
	}
}

type Strategy interface {
	Init(ctx *MarketStrategyContext)
	OnBar(ctx *MarketStrategyContext, ts int64) []*model.StrategyResult
	Name() string
}
