package model

type Broker interface {
	GetCurrentLivePositions(contract *Contract) *ContractPosition
	AddOrder(order *Order) error
}
