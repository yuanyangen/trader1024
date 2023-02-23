package model

type MarKetType int64

const MarKetType_STOCK MarKetType = 1
const MarKetType_FUTURE MarKetType = 2

type Market struct {
	Name     string
	Type     MarKetType
	MarketId string
	Code     string
	SecId    string
	Exchange string
}
