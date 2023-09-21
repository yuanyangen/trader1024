package main

import (
	"github.com/yuanyangen/trader1024/data/storage_client"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/strategy/strategy_new"
	"time"
)

func main() {
	e := engine.NewTrainEngine(event.NewBackTestDailyEventTrigger(1030494445, 1675697645))
	//e.RegisterStrategy(strategy.NewDualSMAStrategyFactory)
	//e.RegisterStrategy(strategy.NewSingleSMAStrategy)
	e.RegisterStrategy(strategy_new.NewCustomKAMASlopStrategyFactory)
	e.RegisterContract("玉米", "", storage_client.SinaHttpStorage())
	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
	time.Sleep(time.Hour)
}
