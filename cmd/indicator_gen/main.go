package main

import (
	"context"
	"github.com/yuanyangen/trader1024/data/datasource"
	"github.com/yuanyangen/trader1024/data/indicator"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/model"
	"time"
)

func main() {
	ctx := context.Background()
	e := engine.NewIndicatorEngine(
		engine.NewBackTestDailyEventTrigger("20200101", "20250101", "20060102"),
		datasource.NewMongoDataSource(),
		[]model.Indicator{
			indicator.NewSMAIndicator(2),
			indicator.NewSMAIndicator(5),
			indicator.NewSMAIndicator(10),
			indicator.NewSMAIndicator(20),
			indicator.NewSMAIndicator(40),
		},
	)
	e.RegisterContractBySubjectName(ctx, "玉米")
	e.Start()
	time.Sleep(time.Hour)
}
