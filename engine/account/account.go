package account

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"sync"
)

type Account struct {
	Total       decimal.Decimal // 单位是分
	Positions   map[string]*ContractPosition
	GlobalEvent chan *model.EventMsg
	indicator   *CashIndicator
	mu          sync.Mutex
}

func NewAccount(start int64) *Account {
	return &Account{
		Total:       decimal.NewFromInt(start),
		GlobalEvent: make(chan *model.EventMsg, 1024),
		Positions:   map[string]*ContractPosition{},
		indicator:   NewCashIndicator(),
	}
}

func (a *Account) DoPlot(p *charts.Page) {
	a.mu.Lock()
	defer a.mu.Unlock()
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
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Total = a.Total.Add(count)
}

//func (a *Account) AddPosition(marketId string, count decimal.Decimal) {
//	a.mu.Lock()
//	defer a.mu.Unlock()
//	position, ok := a.Positions[marketId]
//	if !ok {
//		position = &ContractPosition{ContractId: marketId, Count: decimal.NewFromInt(0)}
//	}
//	position.Count = position.Count.Add(count)
//	a.Positions[marketId] = position
//}

func (a *Account) GetPositionByMarket(marketId string) *ContractPosition {
	a.mu.Lock()
	defer a.mu.Unlock()
	position, ok := a.Positions[marketId]
	if !ok {
		position = &ContractPosition{MarketId: marketId, Count: decimal.NewFromInt(0)}
		a.Positions[marketId] = position
	}
	return position
}

var defaultAccount *Account

func RegisterAccount(account *Account) {
	defaultAccount = account
}

func GetAccount() *Account {
	return defaultAccount
}
