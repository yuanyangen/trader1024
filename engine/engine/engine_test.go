package engine

import (
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/strategy"
	"testing"
)

func TestEngine(t *testing.T) {
	engine := NewEngine()
	st := strategy.NewDualSMAStrategy()
	df := data_feed.NewCsvKLineDataFeed("/home/yuanyangen/HomeData/go/trader1024/datas/datas/daily/IC主力合约.csv")
	df.SetDataMeta(&data_feed.DataMeta{Name: "IC主力合约", Type: data_feed.DataType_FUTURE, Source: data_feed.SourceType_CSV})
	engine.RegisterStrategy(st)
	engine.RegisterMarket("IC主力合约", df)
	engine.RegisterAccount(&model.Account{Total: 100000})
	engine.Start()
}
