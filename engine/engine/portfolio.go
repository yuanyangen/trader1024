package engine

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
)

type PortfolioStrategy func(broker account.Broker, req *ContractPortfolioReq)

type ContractPortfolioReq struct {
	Contract       *model.Contract
	StrategyResult *model.StrategyResult
	Ts             int64
}
