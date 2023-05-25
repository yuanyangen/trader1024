package engine

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/model"
)

type Engine struct {
	Markets        map[string]*MarketEngine
	EventTrigger   model.EventTrigger
	strategies     []func() model.Strategy
	watcherBackend *WatcherBackend
}

func NewEngine() *Engine {
	e := &Engine{
		Markets: map[string]*MarketEngine{},
	}
	e.watcherBackend = NewPlotterServers(e)
	return e
}

func (ec *Engine) RegisterEventTrigger(e model.EventTrigger) {
	ec.EventTrigger = e
}

func (ec *Engine) RegisterMarket(name string) {
	if len(ec.strategies) == 0 {
		panic("should register strategy first")
	}
	strategies := make([]model.Strategy, len(ec.strategies))
	for i, stFactory := range ec.strategies {
		strategies[i] = stFactory()
	}
	df := data_feed.NewCsvKLineDataFeed(name)

	m := NewMarket(name, df, strategies)
	ec.Markets[name] = m
}
func (ec *Engine) RegisterStrategy(stFactory func() model.Strategy) {
	ec.strategies = append(ec.strategies, stFactory)
}
func (ec *Engine) Start() error {
	if err := ec.checkEngine(); err != nil {
		panic(err)
	}

	ec.connectComponent()
	ec.doStart()
	return nil
}

func (ec *Engine) doStart() {
	ec.EventTrigger.Start()

	for _, market := range ec.Markets {
		market.Start()
	}
	ec.watcherBackend.Start()
}

func (ec *Engine) connectComponent() {
	for _, v := range ec.Markets {
		v.DataFeed.SetEventTrigger(ec.EventTrigger)
	}
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
