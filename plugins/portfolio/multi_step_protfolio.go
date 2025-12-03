package portfolio

import (
	"sort"

	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/local_broker"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 执行资金管理动作，根据策略的输出，结合历史的仓位，决定下一步动作。
// 当前写死了， 只执行一次的策略。
func MultiStepPortfolio(ctx *model.TradeObjectEngineContext) {
	count := decimal.NewFromInt(10)
	var err error
	contractPositions := local_broker.GetLocalBroker().GetPositionsByContract(ctx.TradeObjecct)

	allOnlinePosition := []*model.PositionPair{}
	for _, v := range contractPositions.Details {
		if !v.Clear {
			allOnlinePosition = append(allOnlinePosition, v)
		}
	}
	sort.Slice(allOnlinePosition, func(i, j int) bool {
		return allOnlinePosition[i].CreateTimeStamp < allOnlinePosition[j].CreateTimeStamp
	})
	for _, sr := range ctx.StrategyResult {
		switch sr.Cmd {
		case model.PositionTypeClearLast:
			{
				lastPosition := contractPositions.GetLastPair()
				if lastPosition != nil && !lastPosition.Clear {
					if lastPosition.Buy != nil && lastPosition.Buy.Count.GreaterThan(decimal.Zero) {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeSell,
							Count:           lastPosition.Buy.Count.Abs(),
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					} else if lastPosition.Sell != nil && lastPosition.Sell.Count.GreaterThan(decimal.Zero) {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeBuy,
							Count:           lastPosition.Sell.Count.Abs(),
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					}
				}

			}
		case model.PositionTypeClear:
			{
				if !contractPositions.Count.Equal(decimal.Zero) {
					if contractPositions.Count.GreaterThan(decimal.Zero) {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeSell,
							Count:           contractPositions.Count.Abs(),
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					} else {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeBuy,
							Count:           contractPositions.Count.Abs(),
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					}
				}
			}
		case model.PositionTypeLong:
			{
				if contractPositions.IsEmpty() {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeBuy,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				} else if contractPositions.Count.GreaterThan(decimal.Zero) {
					lastPosition := allOnlinePosition[len(allOnlinePosition)-1]
					if sr.Price.Sub(lastPosition.Buy.Price).Div(lastPosition.Buy.Price).GreaterThan(decimal.NewFromFloat(0.01)) {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeBuy,
							Count:           count,
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					}
				} else if contractPositions.Count.LessThan(decimal.Zero) {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeBuy,
						Count:           contractPositions.Count.Abs(),
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          sr.Reason,
						CreateTimeStamp: ctx.Ts,
					}, &model.Order{
						OrderType:       model.OrderTypeBuy,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				}
			}

		case model.PositionTypeShort:
			{
				if contractPositions.IsEmpty() {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeSell,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				} else if contractPositions.Count.GreaterThan(decimal.Zero) {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeSell,
						Count:           contractPositions.Count.Abs(),
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          sr.Reason,
						CreateTimeStamp: ctx.Ts,
					}, &model.Order{
						OrderType:       model.OrderTypeSell,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				} else if contractPositions.Count.LessThan(decimal.Zero) {
					lastPosition := allOnlinePosition[len(allOnlinePosition)-1]
					if sr.Price.Sub(lastPosition.Sell.Price).Div(lastPosition.Sell.Price).GreaterThan(decimal.NewFromFloat(0.01)) {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeSell,
							Count:           count,
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					}
				}
			}
		default:
			panic("should not reach here")
		}
		if err != nil {

		}
	}
}
