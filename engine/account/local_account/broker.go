package local_account

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"sync"
)

type LocalBroker struct {
	mu        sync.Mutex
	orders    []*model.Order
	positions map[string]*model.ContractPosition // marketId
}

var defaultBackTestBroker = &LocalBroker{
	positions: map[string]*model.ContractPosition{},
}

func GetBackTestBroker() model.Broker {
	return defaultBackTestBroker
}

func (btb *LocalBroker) GetCurrentLivePositions(marketId string) *model.ContractPosition {
	btb.mu.Lock()
	defer btb.mu.Unlock()
	position, _ := btb.positions[marketId]
	if position == nil {
		position = &model.ContractPosition{ContractId: marketId, Count: decimal.NewFromInt(0)}
	}
	btb.positions[marketId] = position
	return position
}

func (btb *LocalBroker) AddOrder(contract *model.Contract, t model.OrderType, count decimal.Decimal, price decimal.Decimal, reason string, ts int64) error {
	logs.Info("time=%v clean order_type=%v count=%v %v", utils.TsToString(ts), t, count, reason)

	order := &model.Order{OrderType: t, Price: price, Count: count, MarketId: contract.Id(), Reason: reason, CreateTimeStamp: ts}
	btb.orders = append(btb.orders, order)
	flag := int64(1)
	position := btb.GetCurrentLivePositions(contract.Id())
	if t == model.OrderTypeSell {
		flag = 1
	} else if t == model.OrderTypeBuy {
		flag = -1
	} else {
		panic("should not reach here")
	}

	position.ProcessOrder(order)
	DefaultLocalAccount.GetPositionByMarket(contract.Id()).ProcessOrder(order)
	change := count.Mul(price).Mul(decimal.NewFromInt(flag))
	DefaultLocalAccount.ChangeValue(change)

	if position.IsEmpty() {
		//logs.Info("time=%v account_total order_type=%v count=%v %v", utils.TsToString(ts), t, count, reason)
	}
	return nil
}
