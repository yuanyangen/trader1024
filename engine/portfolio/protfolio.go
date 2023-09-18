package portfolio

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 执行资金管理动作，根据策略的输出，结合历史的仓位，决定下一步动作。
// 当前写死了， 只执行一次的策略。
func Portfolio(req *model.ContractPortfolioReq) {
	broker := account.GetBackTestBroker()
	count := decimal.NewFromInt(100)
	var err error
	switch req.StrategyResult.Cmd {
	case model.StrategyOutVolatility:
		{
			position := broker.GetCurrentLivePositions(req.Contract.Id())
			if !position.Count.Equal(decimal.Zero) {
				if position.Count.GreaterThan(decimal.Zero) {
					err = broker.AddOrder(req.Contract.Id(), account.OrderTypeSell, position.Count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
				} else {
					err = broker.AddOrder(req.Contract.Id(), account.OrderTypeBuy, position.Count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
				}
			}
		}
	case model.StrategyOutLong:
		{
			err = broker.AddOrder(req.Contract.Id(), account.OrderTypeBuy, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
		}
	case model.StrategyOutShort:
		{
			err = broker.AddOrder(req.Contract.Id(), account.OrderTypeSell, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
		}
	default:
		panic("should not reach here")
	}
	if err != nil {

	}
}
