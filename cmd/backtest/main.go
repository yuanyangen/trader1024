package main

import (
	"github.com/yuanyangen/trader1024/data/storage_client"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/strategy/portfolio"
	"github.com/yuanyangen/trader1024/strategy/strategy"
	"time"
)

func main() {

	e := engine.NewLiveExecuteEngine(
		event.NewBackTestDailyEventTrigger(1030494445, 1675697645),
		[]engine.PortfolioStrategy{
			portfolio.MultiStepPortfolio,
			portfolio.Evacuation,
		})
	//e.RegisterStrategy(strategy_old.NewDualSMAStrategyFactory)
	//e.RegisterStrategy(strategy_old.NewSingleSMAStrategy)
	//e.RegisterStrategy(strategy.NewCustomLAMAStrategy2Factory)
	e.RegisterStrategy(strategy.NewSingleKAMAlineStrategyFactory)
	//e.RegisterContract("玉米", "", storage_client.SinaHttpStorage())
	e.RegisterContract("橡胶", "", storage_client.SinaHttpStorage())
	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
	time.Sleep(time.Hour)
}
