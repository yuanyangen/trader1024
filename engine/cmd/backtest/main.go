package main

import (
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/strategy"
)

func main() {
	engine := engine.NewEngine()
	st := strategy.NewDualSMAStrategy()
	df := data_feed.NewCsvKLineDataFeed("/home/yuanyangen/HomeData/go/trader1024/datas/datas/daily/IC主力合约.csv")
	df.SetDataMeta(&indicator.DataMeta{Name: "IC主力合约", Type: indicator.DataType_FUTURE, Source: indicator.SourceType_CSV})
	engine.RegisterStrategy(st)
	engine.RegisterMarket("IC主力合约", df)
	engine.RegisterAccount(&model.Account{Total: 100000})
	engine.Start()
}
