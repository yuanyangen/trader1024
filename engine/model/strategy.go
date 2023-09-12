package model

import (
	"github.com/shopspring/decimal"
)

type ContractStrategyContext struct {
	Contract *Contract
	Kline    MarketIndicator
}

func NewStrategyResult(Cmd StrategyOut, price decimal.Decimal) *StrategyResult {
	return &StrategyResult{
		Cmd:   Cmd,
		Price: price,
	}
}

type Strategy interface {
	Init(ctx *ContractStrategyContext)
	OnBar(ctx *ContractStrategyContext, ts int64) *StrategyResult
	Name() string
}
