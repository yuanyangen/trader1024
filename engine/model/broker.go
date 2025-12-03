package model

type Broker interface {
	GetPositionsByContract(contract *TradeObject) *ContractPositions
	AddOrder(order *Order) error
}
