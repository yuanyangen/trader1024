package model

import (
	"github.com/shopspring/decimal"
)

const OrderTypeSell OrderType = 1
const OrderTypeBuy OrderType = 2

type OrderType int

func (ot OrderType) String() string {
	if ot == OrderTypeSell {
		return "sell"
	} else {
		return "buy"
	}
}

type OrderStatus int

const OrderStatusSucccess = 1
const OrderStatusFinished = 2

type Order struct {
	Contract        *TradeObject
	BrokerId        string
	OrderType       OrderType
	Count           decimal.Decimal
	Price           decimal.Decimal
	Status          OrderStatus
	Reason          string
	CreateTimeStamp int64
}
