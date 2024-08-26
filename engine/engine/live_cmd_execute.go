package engine

import (
	"github.com/yuanyangen/trader1024/engine/model"
)

type LiveCmdExecutor struct {
	Contract          *model.Contract
	kline             *model.KLine
	portfolioStrategy []model.PortfolioStrategy
}

func NewLiveCmdExecutor(contract *model.Contract, kline *model.KLine, portfolioStrategy []model.PortfolioStrategy) CmdExecutor {
	t := &LiveCmdExecutor{
		Contract:          contract,
		kline:             kline,
		portfolioStrategy: portfolioStrategy,
	}
	return t
}

func (t *LiveCmdExecutor) ExecuteCmd(req *model.ContractPortfolioReq) {

	//account.GetAccount().EventTrigger(req.Ts)
}

func (t *LiveCmdExecutor) Report() {
}

func (t *LiveCmdExecutor) Init() {
}
