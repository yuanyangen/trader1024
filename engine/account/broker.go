package account

import (
	"github.com/shopspring/decimal"
	"sync"
)

type Broker interface {
	GetCurrentLivePositions(marketId string) *MarketPosition
	AddOrder(marketId string, t OrderType, count decimal.Decimal, price decimal.Decimal, reason string, ts int64) error
}

type BackTestBroker struct {
	mu        sync.Mutex
	orders    []*Order
	positions map[string]*MarketPosition // marketId
}

var defaultBackTestBroker = &BackTestBroker{
	positions: map[string]*MarketPosition{},
}

func GetBackTestBroker() Broker {
	return defaultBackTestBroker
}

func (btb *BackTestBroker) GetCurrentLivePositions(marketId string) *MarketPosition {
	btb.mu.Lock()
	defer btb.mu.Unlock()
	position, _ := btb.positions[marketId]
	if position == nil {
		position = &MarketPosition{MarketId: marketId, Count: decimal.NewFromInt(0)}
	}
	btb.positions[marketId] = position
	return position
}

func (btb *BackTestBroker) AddOrder(marketId string, t OrderType, count decimal.Decimal, price decimal.Decimal, reason string, ts int64) error {
	order := &Order{OrderType: t, Price: price, Count: count, MarketId: marketId, Reason: reason, CreateTimeStamp: ts}
	btb.orders = append(btb.orders, order)
	flag := int64(1)
	position := btb.GetCurrentLivePositions(marketId)
	if t == OrderTypeSell {
		flag = 1
	} else if t == OrderTypeBuy {
		flag = -1
	} else {
		panic("should not reach here")
	}

	position.ProcessOrder(order)
	GetAccount().GetPositionByMarket(marketId).ProcessOrder(order)
	change := count.Mul(price).Mul(decimal.NewFromInt(flag))
	GetAccount().ChangeValue(change)
	return nil
}
