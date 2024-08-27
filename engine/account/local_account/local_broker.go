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

var defaultLocalBroker = &LocalBroker{
	positions: map[string]*model.ContractPosition{},
}

func GetLocalBroker() model.Broker {
	return defaultLocalBroker
}

func (btb *LocalBroker) GetCurrentLivePositions(contract *model.Contract) *model.ContractPosition {
	btb.mu.Lock()
	defer btb.mu.Unlock()
	position, _ := btb.positions[contract.Id()]
	if position == nil {
		position = &model.ContractPosition{Contract: contract, Count: decimal.NewFromInt(0)}
	}
	btb.positions[contract.Id()] = position
	return position
}

func (btb *LocalBroker) AddOrder(order *model.Order) error {
	logs.Info("time=%v clean order_type=%v count=%v %v", utils.TsToString(order.CreateTimeStamp), order.OrderType, order.Count, order.Reason)
	btb.orders = append(btb.orders, order)
	flag := int64(1)
	position := btb.GetCurrentLivePositions(order.Contract)
	if order.OrderType == model.OrderTypeSell {
		flag = 1
	} else if order.OrderType == model.OrderTypeBuy {
		flag = -1
	} else {
		panic("should not reach here")
	}

	position.ProcessOrder(order)
	DefaultLocalAccount.GetPositionByMarket(order.Contract).ProcessOrder(order)
	change := order.Count.Mul(order.Price).Mul(decimal.NewFromInt(flag))
	DefaultLocalAccount.ChangeValue(change)

	if position.IsEmpty() {
		//logs.Info("time=%v account_total order_type=%v count=%v %v", utils.TsToString(ts), t, count, reason)
	}
	return nil
}
