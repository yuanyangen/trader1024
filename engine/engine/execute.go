package engine

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/portfolio"
)

type LiveCmdExecutor struct {
	Contract *model.Contract
	kline    model.MarketIndicator
}

func newLiveCmdExecutor(contract *model.Contract, kline model.MarketIndicator) CmdExecutor {
	t := &LiveCmdExecutor{
		Contract: contract,
		kline:    kline,
	}
	return t
}

func (t *LiveCmdExecutor) ExecuteCmd(req *model.MarketPortfolioReq) {
	portfolio.Portfolio(req)
	account.GetAccount().EventTrigger(req.Ts)
}

func (t *LiveCmdExecutor) Report() {
}

func (t *LiveCmdExecutor) Init() {
}
