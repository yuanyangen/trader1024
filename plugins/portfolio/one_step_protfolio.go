package portfolio

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 执行资金管理动作，根据策略的输出，结合历史的仓位，决定下一步动作。
// 当前写死了， 只执行一次的策略。
func OneStepPortfolio(broker model.Broker, ctx *model.TradeObjectEngineContext) {
	count := decimal.NewFromInt(100)
	position := broker.GetPositionsByContract(ctx.TradeObjecct)
	for _, sr := range ctx.StrategyResult {
		switch sr.Cmd {
		case model.PositionTypeClear:
			{
				if !position.Count.Equal(decimal.Zero) {
					if position.Count.GreaterThan(decimal.Zero) {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeSell,
							Count:           position.Count,
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          "evacuation_" + sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					} else {
						ctx.Orders = append(ctx.Orders, &model.Order{
							OrderType:       model.OrderTypeBuy,
							Count:           position.Count,
							Price:           sr.Price,
							Contract:        ctx.TradeObjecct,
							Reason:          "evacuation_" + sr.Reason,
							CreateTimeStamp: ctx.Ts,
						})
					}
				}
			}
		case model.PositionTypeLong:
			{
				if position.IsEmpty() {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeBuy,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          "evacuation_" + sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				} else if position.Count.GreaterThan(decimal.Zero) {

				} else if position.Count.LessThan(decimal.Zero) {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeBuy,
						Count:           position.Count.Abs(),
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          "evacuation_" + sr.Reason,
						CreateTimeStamp: ctx.Ts,
					}, &model.Order{
						OrderType:       model.OrderTypeBuy,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          "evacuation_" + sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				}
			}

		case model.PositionTypeShort:
			{
				if position.IsEmpty() {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeSell,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          "evacuation_" + sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				} else if position.Count.GreaterThan(decimal.Zero) {
					ctx.Orders = append(ctx.Orders, &model.Order{
						OrderType:       model.OrderTypeSell,
						Count:           position.Count.Abs(),
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          "evacuation_" + sr.Reason,
						CreateTimeStamp: ctx.Ts,
					}, &model.Order{
						OrderType:       model.OrderTypeSell,
						Count:           count,
						Price:           sr.Price,
						Contract:        ctx.TradeObjecct,
						Reason:          "evacuation_" + sr.Reason,
						CreateTimeStamp: ctx.Ts,
					})
				} else if position.Count.LessThan(decimal.Zero) {
				}
			}
		default:
			panic("should not reach here")
		}
	}

}
