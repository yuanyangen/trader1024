package account

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/event"
	"sync"
)

type Account struct {
	Total       decimal.Decimal // 单位是分
	Positions   map[string]*Position
	GlobalEvent chan *event.EventMsg
	indicator   *CashIndicator
	mu          sync.Mutex
}

func NewAccount(start int64) *Account {
	return &Account{
		Total:       decimal.NewFromInt(start),
		GlobalEvent: make(chan *event.EventMsg, 1024),
		Positions:   map[string]*Position{},
		indicator:   NewCashIndicator(),
	}
}

func (a *Account) DoPlot(p *charts.Page) {
	a.indicator.DoPlot(p)
	a.showFinalNum()
}
func (a *Account) showFinalNum() {
	total := a.Total
	fmt.Println(total.String())
}

func (a *Account) EventTrigger(ts int64) {
	currentVal, _ := a.Total.Float64()
	a.indicator.AddData(ts, currentVal)
}

func (a *Account) ChangeValue(count decimal.Decimal) {
	a.Total = a.Total.Add(count)
}

func (a *Account) AddPosition(marketId string, count decimal.Decimal) {
	a.mu.Lock()
	defer a.mu.Unlock()
	position, ok := a.Positions[marketId]
	if !ok {
		position = &Position{Count: decimal.NewFromInt(0)}
	}
	position.Count = position.Count.Add(count)
	a.Positions[marketId] = position
}

func (a *Account) GetPositionByMarket(marketId string) *Position {
	a.mu.Lock()
	defer a.mu.Unlock()
	position, ok := a.Positions[marketId]
	if ok && position != nil && !position.IsEmpty() {
		return position
	}
	return nil
}

var defaultAccount *Account

func RegisterAccount(account *Account) {
	defaultAccount = account
}

func GetAccount() *Account {
	return defaultAccount
}
