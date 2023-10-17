package engine

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
)

type LiveCmdExecutor struct {
	Contract          *model.Contract
	kline             model.ContractIndicator
	portfolioStrategy []PortfolioStrategy
}

func newLiveCmdExecutor(contract *model.Contract, kline model.ContractIndicator, portfolioStrategy []PortfolioStrategy) CmdExecutor {
	t := &LiveCmdExecutor{
		Contract:          contract,
		kline:             kline,
		portfolioStrategy: portfolioStrategy,
	}
	return t
}

func (t *LiveCmdExecutor) ExecuteCmd(req *ContractPortfolioReq) {
	broker := account.GetBackTestBroker()
	for _, p := range t.portfolioStrategy {
		p(broker, req)
	}
	account.GetAccount().EventTrigger(req.Ts)
}

func (t *LiveCmdExecutor) Report() {
}

func (t *LiveCmdExecutor) Init() {
}
