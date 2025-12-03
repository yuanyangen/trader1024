package main

import (
	"context"
	"time"

	"github.com/yuanyangen/trader1024/data/datasource"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/local_broker"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/portfolio"
	"github.com/yuanyangen/trader1024/plugins/strategy"
)

func main() {
	ctx := context.Background()

	local_broker.InitLocalAccount(100000)
	e := engine.NewEngine(
		engine.NewBackTestDailyEventTrigger("20200101", "20270101", "20060102"),
		datasource.NewMongoDataSource(),
		[]func() model.Strategy{
			strategy.NewCustomStrategy2Factory,
		},
		[]model.PortfolioStrategy{
			portfolio.MultiStepPortfolio,
		},
		nil,
	)

	//e.RegisterContractBySubjectName(ctx, "玉米")
	e.RegisterContract(ctx, "玉米", "202201")
	e.Start()
	time.Sleep(time.Hour)
}
