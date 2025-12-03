package local_broker

import (
	"sync"

	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type LocalBroker struct {
	mu        sync.Mutex
	orders    []*model.Order
	positions map[string]*model.ContractPositions // marketId
}

var defaultLocalBroker = &LocalBroker{
	positions: map[string]*model.ContractPositions{},
}

func GetLocalBroker() model.Broker {
	return defaultLocalBroker
}

func (lb *LocalBroker) GetPositionsByContract(contract *model.TradeObject) *model.ContractPositions {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	position, _ := lb.positions[contract.Id()]
	if position == nil {
		position = &model.ContractPositions{Contract: contract, Count: decimal.NewFromInt(0)}
	}
	lb.positions[contract.Id()] = position
	return position
}

func (lb *LocalBroker) AddOrder(order *model.Order) error {

	lb.orders = append(lb.orders, order)
	flag := int64(1)
	contractPosition := lb.GetPositionsByContract(order.Contract)
	if order.OrderType == model.OrderTypeSell {
		flag = 1
	} else if order.OrderType == model.OrderTypeBuy {
		flag = -1
	} else {
		panic("should not reach here")
	}

	logs.Info("order info: time=%v order_type=%v count=%v price=%v reason=[%v]", utils.TsToString(order.CreateTimeStamp), order.OrderType, order.Count, order.Price.String(), order.Reason)
	contractPosition.ProcessOrder(order)
	change := order.Count.Mul(order.Price).Mul(decimal.NewFromInt(flag))
	DefaultLocalAccount.ChangeValue(change)
	return nil
}
