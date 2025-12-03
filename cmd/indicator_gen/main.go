package main

import (
	"context"
	"github.com/yuanyangen/trader1024/data/datasource"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

func main() {
	ctx := context.Background()
	eventTrigger := engine.NewBackTestDailyEventTrigger("20200101", "20250101", "20060102")
	e := engine.NewIndicatorEngine(
		eventTrigger,
		datasource.NewMongoDataSource(),
		[]model.Indicator{
			indicator.NewZDMAIndicator(5, indicator.NewBOPIndicator()),
			indicator.NewZDMAIndicator(10, indicator.NewBOPIndicator()),
			indicator.NewZDMAIndicator(20, indicator.NewBOPIndicator()),
			indicator.NewZDMAIndicator(40, indicator.NewBOPIndicator()),
		},
		[]model.Indicator{
			indicator.NewCCIIndicator(5),
			indicator.NewCCIIndicator(10),
			indicator.NewCCIIndicator(20),
			indicator.NewCCIIndicator(30),
			indicator.NewCCIIndicator(40),
			indicator.NewCCIIndicator(60),
		},
		[]model.Indicator{
			indicator.NewMFIIndicator(5),
			indicator.NewMFIIndicator(10),
			indicator.NewMFIIndicator(20),
			indicator.NewMFIIndicator(30),
			indicator.NewMFIIndicator(40),
			indicator.NewMFIIndicator(60),
		},
		[]model.Indicator{
			indicator.NewSMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewSMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewSMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewSMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewSMAIndicator(30, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewSMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewSMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewSMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewSMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewSMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewSMAIndicator(30, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewSMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		},
		[]model.Indicator{
			indicator.NewJMAIndicator(0.15, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewJMAIndicator(0.15, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewJMAIndicator(0.15, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewJMAIndicator(0.15, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewJMAIndicator(0.15, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewJMAIndicator(0.2, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewJMAIndicator(0.2, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewJMAIndicator(0.2, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewJMAIndicator(0.2, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewJMAIndicator(0.2, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewJMAIndicator(0.3, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewJMAIndicator(0.3, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewJMAIndicator(0.3, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewJMAIndicator(0.3, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewJMAIndicator(0.3, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewJMAIndicator(0.4, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewJMAIndicator(0.4, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewJMAIndicator(0.4, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewJMAIndicator(0.4, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewJMAIndicator(0.4, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewJMAIndicator(0.5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewJMAIndicator(0.5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewJMAIndicator(0.5, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewJMAIndicator(0.5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewJMAIndicator(0.5, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewJMAIndicator(0.6, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewJMAIndicator(0.6, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewJMAIndicator(0.6, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewJMAIndicator(0.6, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewJMAIndicator(0.6, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),
		},
		[]model.Indicator{
			indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewZDMAIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewZDMAIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewZDMAIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
			indicator.NewZDMAIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewZDMAIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

			indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
			indicator.NewZDMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		},
		getSlopIndictor(),
		[]model.Indicator{
			indicator.NewSMAIndicator(4, indicator.NewRawKlineIndicator(indicator.DataSourceCloseSubOpen)),
		},
		[]model.Indicator{
			indicator.NewLRIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewLRIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewLRIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),
			indicator.NewLRIndicator(80, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),
			indicator.NewLRIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),
			indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),
			indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
			indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
			indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),
			indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
			indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		},
		[]model.Indicator{
			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(5, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(10, indicator.NewLRIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewZDMAIndicator(20, indicator.NewLRIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		},
		[]model.Indicator{
			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewLRIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewLRIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		}, []model.Indicator{
			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(5, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewSMAIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(40, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),

			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceLow))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceHigh))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
			indicator.NewSMAIndicator(20, indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		},
		//getSlopIndictor(),
		//getKAMAForZDMAIndictor(),
	)
	//e.RegisterContractBySubjectName(ctx, "玉米")
	e.RegisterContract(ctx, "玉米", "202201")

	e.Start()
	eventTrigger.Wait()
}
