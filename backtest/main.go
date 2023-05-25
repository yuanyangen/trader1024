package main

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/strategy/strategy"
)

func main() {
	e := engine.NewEngine()
	//e.RegisterStrategy(strategy.NewDualSMAStrategyFactory)
	//e.RegisterStrategy(strategy.NewSingleSMAStrategy)
	e.RegisterStrategy(strategy.NewCustomLAMAStrategyFactory)
	e.RegisterEventTrigger(event.NewBackTestDailyEventTrigger(1430494445, 1675697645))
	e.RegisterMarket("sp2305")
	//e.RegisterMarket("fu2304")
	//e.RegisterMarket("bu2307")
	//e.RegisterMarket("b2305")
	//e.RegisterMarket("eb2304")

	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
}
