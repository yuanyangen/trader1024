package model

import "github.com/shopspring/decimal"

type StrategyOut int64

const StrategyOutLong StrategyOut = 1
const StrategyOutShort StrategyOut = 2

type StrategyResult struct {
	DecisionDesc string
	Cmd          StrategyOut
	Price        decimal.Decimal
}
