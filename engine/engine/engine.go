package engine

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/indicator/global"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type Engine struct {
	Markets        map[string]*Market
	Strategys      []model.Strategy
	watcherBackend *WatcherBackend
	Account        *model.Account
	globalEvent    chan *indicator.GlobalMsg
	globalWatchers []indicator.GlobalIndicator
}

func NewEngine() *Engine {
	e := &Engine{
		Markets:     map[string]*Market{},
		globalEvent: make(chan *indicator.GlobalMsg, 1024),
	}
	e.watcherBackend = NewPlotterServers(e)
	e.globalWatchers = []indicator.GlobalIndicator{
		global.NewCashIndicator(),
	}
	return e
}

//func (ec *Engine) RegisterWatcher(df model.GlobalIndicator) {
//	ec.Watchers = append(ec.Watchers, df)
//}

func (ec *Engine) RegisterAccount(account *model.Account) {
	ec.Account = account
}

func (ec *Engine) RegisterMarket(name string, df data_feed.DataFeed) {
	if len(ec.Strategys) == 0 {
		panic("should register strategy first")
	}
	m := NewMarket(name, df, ec.Strategys)
	ec.Markets[name] = m
}
func (ec *Engine) RegisterStrategy(st model.Strategy) {
	ec.Strategys = append(ec.Strategys, st)
}
func (ec *Engine) Start() error {
	if err := ec.checkEngine(); err != nil {
		panic(err)
	}
	for _, market := range ec.Markets {
		market.Start()
	}
	ec.globalWatcherDataDaemon()
	ec.watcherBackend.Start()
	return nil
}

func (ec *Engine) checkEngine() error {
	if len(ec.Markets) == 0 {
		return fmt.Errorf("market not configed")
	}
	//if ec.Sizer == nil {
	//	return fmt.Errorf("sizer not configed")
	//}
	//if len(ec.Broker) == 0 {
	//	return fmt.Errorf("broker not configed")
	//}
	return nil
}

func (ec *Engine) DoPlot(p *charts.Page) {
	for _, v := range ec.globalWatchers {
		v.DoPlot(p)
	}
}

func (ec *Engine) globalWatcherDataDaemon() {
	utils.AsyncRun(func() {
		for v := range ec.globalEvent {
			for _, w := range ec.globalWatchers {
				w.AddData(v)
			}
		}
	})
}
