package model

import (
	"github.com/shopspring/decimal"
)

type Broker interface {
	GetCurrentLivePositions(contractId string) *ContractPosition
	AddOrder(contract *Contract, t OrderType, count decimal.Decimal, price decimal.Decimal, reason string, ts int64) error
}
