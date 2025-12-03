package model

type PortfolioStrategy func(req *TradeObjectEngineContext)

type TradeObjectEngineContext struct {
	TradeObject  *TradeObject
	Kline        *KLine
	CurrentKNode *KLineNode
	Ts           int64

	StrategyResult []*StrategyResult
	Orders         []*Order
}
