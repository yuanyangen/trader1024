package engine

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/strategy/strategy"
	"testing"
)

func TestEngine(t *testing.T) {
	e := NewEngine()
	e.RegisterStrategy(strategy.NewDualSMAStrategyFactory)
	e.RegisterEventTrigger(event.NewBackTestDailyEventTrigger(1430494445, 1675697645))
	//e.RegisterMarket("ICM", data_feed.NewCsvKLineDataFeed("/home/yuanyangen/HomeData/go/trader1024/data/datas/daily/IC主力合约.storage"))
	e.RegisterMarket("不锈钢主力", data_feed.NewCsvKLineDataFeed("/home/yuanyangen/HomeData/go/trader1024/data/datas/daily/不锈钢主力.storage"))
	account.RegisterAccount(account.NewAccount(10000000))
	e.Start()
}
