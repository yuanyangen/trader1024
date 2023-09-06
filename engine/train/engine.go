package train

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/model"
)

type TrainEngine struct {
	Markets      map[string]*TrainMarketEngine
	EventTrigger model.EventTrigger
	strategies   []func() model.Strategy
	//watcherBackend *WatcherBackend
}

func NewEngine() *TrainEngine {
	e := &TrainEngine{
		Markets: map[string]*TrainMarketEngine{},
	}
	return e
}

func (ec *TrainEngine) RegisterEventTrigger(e model.EventTrigger) {
	ec.EventTrigger = e
}

func (ec *TrainEngine) RegisterMarket(name string) {
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
func (ec *TrainEngine) RegisterStrategy(stFactory func() model.Strategy) {
	ec.strategies = append(ec.strategies, stFactory)
}
func (ec *TrainEngine) Start() error {
	if err := ec.checkEngine(); err != nil {
		panic(err)
	}

	ec.connectComponent()
	ec.doStart()
	return nil
}

func (ec *TrainEngine) doStart() {
	ec.EventTrigger.Start()

	for _, market := range ec.Markets {
		market.Start()
	}
}

func (ec *TrainEngine) connectComponent() {
	for _, v := range ec.Markets {
		v.DataFeed.SetEventTrigger(ec.EventTrigger)
	}
}

func (ec *TrainEngine) checkEngine() error {
	if len(ec.Markets) == 0 {
		return fmt.Errorf("market not configed")
	}
	return nil
}
