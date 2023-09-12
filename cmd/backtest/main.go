package main

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/strategy/strategy_new"
)

func main() {
	e := engine.NewEngine()
	//e.RegisterStrategy(strategy.NewDualSMAStrategyFactory)
	//e.RegisterStrategy(strategy.NewSingleSMAStrategy)
	e.RegisterStrategy(strategy_new.NewCustomLAMAStrategy2Factory)
	e.RegisterEventTrigger(event.NewBackTestDailyEventTrigger(1430494445, 1675697645))
	e.RegisterMarket("sp2305")
	//e.RegisterContract("fu2304")
	//e.RegisterContract("bu2307")
	//e.RegisterContract("b2305")
	//e.RegisterContract("eb2304")

	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
}
