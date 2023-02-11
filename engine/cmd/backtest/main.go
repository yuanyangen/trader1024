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
	e.RegisterStrategy(strategy.NewDualSMAStrategy())
	e.RegisterEventTrigger(event.NewBackTestDailyEventTrigger(1430494445, 1675697645))
	e.RegisterMarket("ICM", data_feed.NewCsvKLineDataFeed("/home/yuanyangen/HomeData/go/trader1024/data/datas/daily/IC主力合约.csv"))
	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
}
