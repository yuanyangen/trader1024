package model

import (
	"github.com/yuanyangen/trader1024/engine/indicator"
)

type MarketStrategyContext struct {
	DailyData *indicator.DailyData
	Broker    Broker
	Account   *Account
}

type Strategy interface {
	Init(ctx *MarketStrategyContext)
	OnBar(ctx *MarketStrategyContext, ts int64)
}
