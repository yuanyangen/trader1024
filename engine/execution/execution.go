package execution

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
)

type ExecutionAction struct {
	MarketId   string
	StrategyId string
	Cmd        model.StrategyCmd
	Count      decimal.Decimal
	Price      decimal.Decimal
}

func Execute(actions ...*ExecutionAction) {
	broker := account.GetBackTestBroker()
	for _, action := range actions {
		switch action.Cmd {
		case model.StrategyCmdClean:
			{
				position := broker.GetCurrentLivePositions(action.MarketId)
				zero := decimal.NewFromInt(0)
				if !position.Count.Equal(zero) {
					if position.Count.GreaterThan(zero) {
						broker.AddOrder(action.MarketId, account.OrderTypeSell, action.Count, action.Price)
					} else {
						broker.AddOrder(action.MarketId, account.OrderTypeBuy, action.Count, action.Price)
					}
				}
			}
		case model.StrategyCmdBuy:
			{
				broker.AddOrder(action.MarketId, account.OrderTypeBuy, action.Count, action.Price)
			}
		case model.StrategyCmdSell:
			{
				broker.AddOrder(action.MarketId, account.OrderTypeSell, action.Count, action.Price)
			}

		}
	}
}
