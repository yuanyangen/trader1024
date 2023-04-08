package account

import (
	"github.com/shopspring/decimal"
)

type OrderType int

const OrderTypeSell OrderType = 1
const OrderTypeBuy OrderType = 2

type OrderStatus int

const OrderStatusSucccess = 1
const OrderStatusFinished = 2

type Order struct {
	MarketId        string
	BrokerId        string
	OrderType       OrderType
	Count           decimal.Decimal
	Price           decimal.Decimal
	Status          OrderStatus
	Reason          string
	CreateTimeStamp int64
}
