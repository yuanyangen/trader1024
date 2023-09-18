package model

import "github.com/shopspring/decimal"

type StrategyOut int64

const StrategyOutLong StrategyOut = 1
const StrategyOutShort StrategyOut = 2
const StrategyOutVolatility StrategyOut = 3

type StrategyResult struct {
	StrategyName string
	Reason       string
	Cmd          StrategyOut
	Price        decimal.Decimal
}
