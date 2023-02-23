package main

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/strategy"
)

func main() {
	e := engine.NewEngine()
	e.RegisterStrategy(strategy.NewDualSMAStrategyFactory)
	e.RegisterEventTrigger(event.NewBackTestDailyEventTrigger(1430494445, 1675697645))
	e.RegisterMarket("jdm", data_feed.NewCsvKLineDataFeed("jdm"))
	e.RegisterMarket("jm", data_feed.NewCsvKLineDataFeed("jm"))

	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
}
