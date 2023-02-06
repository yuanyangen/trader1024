package account

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type Account struct {
	Total       int64 // 单位是分
	GlobalEvent chan *event.EventMsg
	indicator   *CashIndicator
}

func NewAccount(start int64) *Account {
	return &Account{
		Total:       start,
		GlobalEvent: make(chan *event.EventMsg, 1024),

		indicator: NewCashIndicator(),
	}
}

func (a *Account) DoPlot(p *charts.Page) {
	a.indicator.DoPlot(p)
}

func (a *Account) RegisterEventTrigger(et event.EventTrigger) {
	if a.GlobalEvent == nil {
		panic("account event chan nil")
	}
	et.RegisterEventReceiver(a.GlobalEvent)
	utils.AsyncRun(func() {
		for v := range a.GlobalEvent {
			a.indicator.AddData(v.TimeStamp, a.Total)
		}
	})
}
