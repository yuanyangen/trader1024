package execution

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
)

type ExecutionAction struct {
	MarketId   string
	StrategyId string
	Cmd        model.StrategyOut
	Count      decimal.Decimal
	Price      decimal.Decimal
	Reason     string
	Ts         int64
}

func Execute(actions ...*ExecutionAction) {
	broker := account.GetBackTestBroker()
	for _, action := range actions {
		var err error
		switch action.Cmd {
		case model.StrategyCmdClean:
			{
				position := broker.GetCurrentLivePositions(action.MarketId)
				if !position.Count.Equal(decimal.Zero) {
					if position.Count.GreaterThan(decimal.Zero) {
						err = broker.AddOrder(action.MarketId, account.OrderTypeSell, action.Count, action.Price, action.Reason, action.Ts)

					} else {
						err = broker.AddOrder(action.MarketId, account.OrderTypeBuy, action.Count, action.Price, action.Reason, action.Ts)
					}
				}
			}
		case model.StrategyCmdBuy:
			{
				err = broker.AddOrder(action.MarketId, account.OrderTypeBuy, action.Count, action.Price, action.Reason, action.Ts)
			}
		case model.StrategyCmdSell:
			{
				err = broker.AddOrder(action.MarketId, account.OrderTypeSell, action.Count, action.Price, action.Reason, action.Ts)
			}
		default:
			panic("should not reach here")
		}
		if err != nil {

		}
	}
}
