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
	Count    int64
}

type MarketPortfolioReq struct {
	Market     *Market
	Strategies []*StrategyReq
}
type StrategyReq struct {
	StrategyName string
	Cmds         []*StrategyResult
	Reason       string
	Ts           int64
}
