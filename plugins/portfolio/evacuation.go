package portfolio

import (
	"sort"

	"github.com/yuanyangen/trader1024/engine/local_broker"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 执行资金管理动作，根据策略的输出，结合历史的仓位，决定下一步动作。
// 当前写死了， 只执行一次的策略。
func Evacuation(ctx *model.TradeObjectEngineContext) {
	position := local_broker.GetLocalBroker().GetPositionsByContract(ctx.TradeObjecct)
	if ctx.StrategyResult == nil {
		return
	}

	allOnlinePosition := []*model.PositionPair{}
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
	for _, sr := range ctx.StrategyResult {
		if lastPosition.Type == model.PositionTypeLong {
			if sr.Price.LessThan(lastPosition.Buy.Price) {
				ctx.Orders = append(ctx.Orders, &model.Order{
					OrderType:       model.OrderTypeSell,
					Count:           position.Count.Abs(),
					Price:           sr.Price,
					Contract:        ctx.TradeObjecct,
					Reason:          "evacuation_" + sr.Reason,
					CreateTimeStamp: ctx.Ts,
				})
			}
		} else if lastPosition.Type == model.PositionTypeShort {
			if sr.Price.GreaterThan(lastPosition.Sell.Price) {
				ctx.Orders = append(ctx.Orders, &model.Order{
					OrderType:       model.OrderTypeBuy,
					Count:           position.Count.Abs(),
					Price:           sr.Price,
					Contract:        ctx.TradeObjecct,
					Reason:          "evacuation_" + sr.Reason,
					CreateTimeStamp: ctx.Ts,
				})
			}
		}
	}
}
