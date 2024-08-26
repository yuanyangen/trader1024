package model

type PortfolioStrategy func(broker Broker, req *ContractPortfolioReq)

type ContractPortfolioReq struct {
	Contract       *Contract
	StrategyResult *StrategyResult
	Ts             int64
}
