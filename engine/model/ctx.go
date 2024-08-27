package model

type PortfolioStrategy func(req *ContractEngineContext)

type ContractEngineContext struct {
	Contract     *Contract
	Kline        *KLine
	CurrentKNode *KLineNode
	Ts           int64

	StrategyResult []*StrategyResult
	Orders         []*Order
}
