package portfolio

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/model"
	"sort"
)

// 执行资金管理动作，根据策略的输出，结合历史的仓位，决定下一步动作。
// 当前写死了， 只执行一次的策略。
func MultiStepPortfolio(broker account.Broker, req *engine.ContractPortfolioReq) {
	count := decimal.NewFromInt(100)
	var err error
	position := broker.GetCurrentLivePositions(req.Contract.Id())
	if req.StrategyResult == nil {
		return
	}

	allOnlinePosition := []*account.PositionPair{}
	for _, v := range position.Details {
		if !v.Clear {
			allOnlinePosition = append(allOnlinePosition, v)
		}
	}
	sort.Slice(allOnlinePosition, func(i, j int) bool {
		return allOnlinePosition[i].CreateTimeStamp < allOnlinePosition[j].CreateTimeStamp
	})

	switch req.StrategyResult.Cmd {
	case model.StrategyOutVolatility:
		{
			if !position.Count.Equal(decimal.Zero) {
				if position.Count.GreaterThan(decimal.Zero) {
					err = broker.AddOrder(req.Contract, account.OrderTypeSell, position.Count.Abs(), req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
				} else {
					err = broker.AddOrder(req.Contract, account.OrderTypeBuy, position.Count.Abs(), req.StrategyResult.Price, req.StrategyResult.String(), req.Ts)
				}
			}
		}
	case model.StrategyOutLong:
		{
			if position.IsEmpty() {
				err = broker.AddOrder(req.Contract, account.OrderTypeBuy, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
			} else if position.Count.GreaterThan(decimal.Zero) {
				lastPosition := allOnlinePosition[len(allOnlinePosition)-1]
				if req.StrategyResult.Price.Sub(lastPosition.Buy.Price).Div(lastPosition.Buy.Price).GreaterThan(decimal.NewFromFloat(0.01)) {
					err = broker.AddOrder(req.Contract, account.OrderTypeBuy, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
				}
			} else if position.Count.LessThan(decimal.Zero) {
				err = broker.AddOrder(req.Contract, account.OrderTypeBuy, position.Count.Abs(), req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
				err = broker.AddOrder(req.Contract, account.OrderTypeBuy, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
			}
		}

	case model.StrategyOutShort:
		{
			if position.IsEmpty() {
				err = broker.AddOrder(req.Contract, account.OrderTypeSell, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
			} else if position.Count.GreaterThan(decimal.Zero) {
				err = broker.AddOrder(req.Contract, account.OrderTypeSell, position.Count.Abs(), req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
				err = broker.AddOrder(req.Contract, account.OrderTypeSell, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
			} else if position.Count.LessThan(decimal.Zero) {
				lastPosition := allOnlinePosition[len(allOnlinePosition)-1]
				if req.StrategyResult.Price.Sub(lastPosition.Sell.Price).Div(lastPosition.Sell.Price).GreaterThan(decimal.NewFromFloat(0.01)) {
					err = broker.AddOrder(req.Contract, account.OrderTypeSell, count, req.StrategyResult.Price, req.StrategyResult.Reason, req.Ts)
				}
			}
		}
	default:
		panic("should not reach here")
	}
	if err != nil {

	}
}
