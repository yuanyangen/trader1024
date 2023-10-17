package portfolio

import (
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/engine"
	"sort"
)

// 执行资金管理动作，根据策略的输出，结合历史的仓位，决定下一步动作。
// 当前写死了， 只执行一次的策略。
func Evacuation(broker account.Broker, req *engine.ContractPortfolioReq) {
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
	if len(allOnlinePosition) == 0 {
		return
	}
	lastPosition := allOnlinePosition[len(allOnlinePosition)-1]
	if lastPosition.Type == account.PositionTypeLong {
		if req.StrategyResult.Price.LessThan(lastPosition.Buy.Price) {
			broker.AddOrder(req.Contract, account.OrderTypeSell, position.Count.Abs(), req.StrategyResult.Price, "evacuation_"+req.StrategyResult.Reason, req.Ts)

		}
	} else if lastPosition.Type == account.PositionTypeShort {
		if req.StrategyResult.Price.GreaterThan(lastPosition.Sell.Price) {
			broker.AddOrder(req.Contract, account.OrderTypeBuy, position.Count.Abs(), req.StrategyResult.Price, "evacuation_"+req.StrategyResult.Reason, req.Ts)
		}
	}
}
