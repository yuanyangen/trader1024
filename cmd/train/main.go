package main

import (
	"context"
	"github.com/yuanyangen/trader1024/data/datasource"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/model"
	"time"
)

func main() {
	ctx := context.Background()

	e := engine.NewTrainEngine(
		engine.NewBackTestDailyEventTrigger("20200101", "20270101", "20060102"),
		datasource.NewMongoDataSource(),
		[]func() model.Strategy{},
	)

	e.RegisterContractBySubjectName(ctx, "玉米")
	e.Start()
	time.Sleep(time.Hour)
}
