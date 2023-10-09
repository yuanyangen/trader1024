package model

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type ContractStrategyContext struct {
	Contract *Contract
	Kline    ContractIndicator
}

func NewStrategyResult(Cmd StrategyOut, price decimal.Decimal, reason string) *StrategyResult {
	return &StrategyResult{
		Cmd:    Cmd,
		Price:  price,
		Reason: reason,
	}
}

type Strategy interface {
	Init(ctx *ContractStrategyContext)
	OnBar(ctx *ContractStrategyContext, ts int64) *StrategyResult
	Name() string
}

const StrategyOutLong StrategyOut = 1
const StrategyOutShort StrategyOut = 2
const StrategyOutVolatility StrategyOut = 3

type StrategyOut int64

func (so StrategyOut) String() string {
	if so == StrategyOutLong {
		return "long"
	} else if so == StrategyOutShort {
		return "short"
	} else {
		return "volatility"
	}
}

type StrategyResult struct {
	StrategyName string
	Reason       string
	Cmd          StrategyOut
	Price        decimal.Decimal
}

func (sr *StrategyResult) String() string {
	return fmt.Sprintf("strategyName=%v cmd=%v price=%v reason=%v", sr.StrategyName, sr.Cmd, sr.Price.String(), sr.Reason)
}
