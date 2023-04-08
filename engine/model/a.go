package model

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
