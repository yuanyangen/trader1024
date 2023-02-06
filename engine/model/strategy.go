package model

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/indicator"
)

type MarketStrategyContext struct {
	DailyData *indicator.DailyIndicators
	Broker    Broker
	Account   *account.Account
}

type Strategy interface {
	Init(ctx *MarketStrategyContext)
	OnBar(ctx *MarketStrategyContext, ts int64)
}
