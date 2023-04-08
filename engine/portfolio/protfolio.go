package portfolio

import (
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/execution"
	"github.com/yuanyangen/trader1024/engine/model"
)

func Portfolio(req *model.MarketPortfolioReq) error {

	count := decimal.NewFromInt(100)
	for _, st := range req.Strategies {
		for _, cmd := range st.Cmds {
			action := &execution.ExecutionAction{
				MarketId:   req.Market.MarketId,
				StrategyId: st.StrategyName,
				Cmd:        cmd.Cmd,
				Count:      count,
				Price:      cmd.Price,
				Reason:     st.Reason,
				Ts:         st.Ts,
			}

			execution.Execute(action)
		}
	}
	return nil
}
