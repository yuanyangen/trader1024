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
	e.RegisterMarket("IC主力合约", data_feed.NewCsvKLineDataFeed("IC主力合约", data_feed.DataType_FUTURE, "/home/yuanyangen/HomeData/go/trader1024/datas/datas/daily/IC主力合约.csv"))
	e.RegisterAccount(account.NewAccount(10000000))
	e.Start()
}
