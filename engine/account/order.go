package account

import (
	"github.com/shopspring/decimal"
	"sync"
)

type OrderType int

const OrderTypeSell OrderType = 1
const OrderTypeBuy OrderType = 2

type OrderStatus int

const OrderStatusSucccess = 1
const OrderStatusFinished = 2

type Order struct {
	MarketId  string
	BrokerId  string
	OrderType OrderType
	Count     decimal.Decimal
	Price     decimal.Decimal
	Status    OrderStatus
}

type Broker interface {
	GetCurrentLivePositions(marketId string) *Position
	AddOrder(marketId string, t OrderType, count decimal.Decimal, price decimal.Decimal) error
}

type BackTestBroker struct {
	mu        sync.Mutex
	orders    []*Order
	positions map[string]*Position
}

var defaultBackTestBroker = &BackTestBroker{
	positions: map[string]*Position{},
}

func GetBackTestBroker() Broker {
	return defaultBackTestBroker
}

func (btb *BackTestBroker) GetCurrentLivePositions(marketId string) *Position {
	btb.mu.Lock()
	defer btb.mu.Unlock()
	position, _ := btb.positions[marketId]
	if position == nil {
		position = &Position{Count: decimal.NewFromInt(0)}
	}
	btb.positions[marketId] = position
	return position
}

func (btb *BackTestBroker) AddOrder(marketId string, t OrderType, count decimal.Decimal, price decimal.Decimal) error {
	btb.orders = append(btb.orders, &Order{OrderType: t, Count: count, MarketId: marketId})
	flag := int64(1)
	position := btb.GetCurrentLivePositions(marketId)
	if t == OrderTypeSell {
		position.Count.Sub(count)
		flag = 1
	} else if t == OrderTypeBuy {
		position.Count.Add(count)
		flag = -1
	}
	change := count.Mul(price).Mul(decimal.NewFromInt(flag))
	GetAccount().ChangeValue(change)
	return nil
}
