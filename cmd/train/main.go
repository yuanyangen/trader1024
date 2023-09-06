package main

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/engine/train"
	"github.com/yuanyangen/trader1024/strategy/strategy_new"
	"time"
)

func main() {
	e := train.NewEngine()
	//e.RegisterStrategy(strategy.NewDualSMAStrategyFactory)
	//e.RegisterStrategy(strategy.NewSingleSMAStrategy)
	e.RegisterStrategy(strategy_new.NewCustomLAMAStrategy2Factory)
	e.RegisterEventTrigger(event.NewBackTestDailyEventTrigger(1430494445, 1675697645))
	e.RegisterMarket("pg2309")
	//e.RegisterMarket("fu2304")
	//e.RegisterMarket("bu2307")
	//e.RegisterMarket("b2305")
	//e.RegisterMarket("eb2304")

	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
	time.Sleep(time.Hour)
}
