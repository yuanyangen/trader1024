package model

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func NewStrategyResult(Cmd PositionType, price decimal.Decimal, reason string) *StrategyResult {
	return &StrategyResult{
		Cmd:    Cmd,
		Price:  price,
		Reason: reason,
	}
}

type Strategy interface {
	Init(ctx *TradeObjectEngineContext)
	OnBar(ctx *TradeObjectEngineContext) *StrategyResult
	Name() string
}

type StrategyResult struct {
	StrategyName string
	Reason       string
	Cmd          PositionType
	Price        decimal.Decimal
}

func (sr *StrategyResult) String() string {
	return fmt.Sprintf("strategyName=%v::cmd=%v::price=%v::reason=%v", sr.StrategyName, sr.Cmd, sr.Price.String(), sr.Reason)
}
