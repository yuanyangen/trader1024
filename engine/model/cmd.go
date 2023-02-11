package model

import "github.com/shopspring/decimal"

type StrategyCmd int64

const StrategyCmdNothing StrategyCmd = 1
const StrategyCmdBuy StrategyCmd = 2
const StrategyCmdClean StrategyCmd = 3 // 平仓， 可以为平多仓， 也可以为平空仓
const StrategyCmdSell StrategyCmd = 4

type StrategyResult struct {
	Cmd   StrategyCmd
	Price decimal.Decimal
}
