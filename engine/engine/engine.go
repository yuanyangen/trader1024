package engine

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/engine/model"
)

type Engine struct {
	Markets        map[string]*Market
	EventTrigger   event.EventTrigger
	strategies     []model.Strategy
	watcherBackend *WatcherBackend
	Account        *account.Account
}

func NewEngine() *Engine {
	e := &Engine{
		Markets: map[string]*Market{},
	}
	e.watcherBackend = NewPlotterServers(e)
	return e
}

func (ec *Engine) RegisterAccount(account *account.Account) {
	ec.Account = account
}
func (ec *Engine) RegisterEventTrigger(e event.EventTrigger) {
	ec.EventTrigger = e
}

func (ec *Engine) RegisterMarket(name string, df data_feed.DataFeed) {
	if len(ec.strategies) == 0 {
		panic("should register strategy first")
	}
	m := NewMarket(name, df, ec.strategies)
	ec.Markets[name] = m
}
func (ec *Engine) RegisterStrategy(st model.Strategy) {
	ec.strategies = append(ec.strategies, st)
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
	ec.Account.RegisterEventTrigger(ec.EventTrigger)
}

func (ec *Engine) checkEngine() error {
	if len(ec.Markets) == 0 {
		return fmt.Errorf("market not configed")
	}
	if ec.Account == nil {
		panic("engine .account nil")
	}
	//if ec.Sizer == nil {
	//	return fmt.Errorf("sizer not configed")
	//}
	//if len(ec.Broker) == 0 {
	//	return fmt.Errorf("broker not configed")
	//}
	return nil
}
