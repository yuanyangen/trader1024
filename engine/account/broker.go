package account

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
	"sync"
)

type Broker interface {
	GetCurrentLivePositions(contractId string) *ContractPosition
	AddOrder(contract *model.Contract, t OrderType, count decimal.Decimal, price decimal.Decimal, reason string, ts int64) error
}

type BackTestBroker struct {
	mu        sync.Mutex
	orders    []*Order
	positions map[string]*ContractPosition // marketId
}

var defaultBackTestBroker = &BackTestBroker{
	positions: map[string]*ContractPosition{},
}

func GetBackTestBroker() Broker {
	return defaultBackTestBroker
}

func (btb *BackTestBroker) GetCurrentLivePositions(marketId string) *ContractPosition {
	btb.mu.Lock()
	defer btb.mu.Unlock()
	position, _ := btb.positions[marketId]
	if position == nil {
		position = &ContractPosition{MarketId: marketId, Count: decimal.NewFromInt(0)}
	}
	btb.positions[marketId] = position
	return position
}

func (btb *BackTestBroker) AddOrder(contract *model.Contract, t OrderType, count decimal.Decimal, price decimal.Decimal, reason string, ts int64) error {
	//logs.Info("time=%v clean order_type=%v count=%v %v", utils.TsToString(ts), t, count, reason)

	order := &Order{OrderType: t, Price: price, Count: count, MarketId: contract.Id(), Reason: reason, CreateTimeStamp: ts}
	btb.orders = append(btb.orders, order)
	flag := int64(1)
	position := btb.GetCurrentLivePositions(contract.Id())
	if t == OrderTypeSell {
		flag = 1
	} else if t == OrderTypeBuy {
		flag = -1
	} else {
		panic("should not reach here")
	}

	position.ProcessOrder(order)
	GetAccount().GetPositionByMarket(contract.Id()).ProcessOrder(order)
	change := count.Mul(price).Mul(decimal.NewFromInt(flag))
	GetAccount().ChangeValue(change)

	//if position.IsEmpty() {
	//	//logs.Info("time=%v account_total order_type=%v count=%v %v", utils.TsToString(ts), t, count, reason)
	//}
	return nil
}
