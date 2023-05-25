package model

import (
	"github.com/shopspring/decimal"
)

type MarketStrategyContext struct {
	Market    *Market
	DailyData *DailyIndicators
}

func NewStrategyResult(Cmd StrategyCmd, price decimal.Decimal) *StrategyResult {
	return &StrategyResult{
		Cmd:   Cmd,
		Price: price,
	}
}

type Strategy interface {
	Init(ctx *MarketStrategyContext)
	OnBar(ctx *MarketStrategyContext, ts int64) []*StrategyResult
	Name() string
}
