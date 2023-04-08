package strategy

import "github.com/yuanyangen/trader1024/engine/model"

type EmptyStrategy struct {
}

func NewEmptyStrategy() Strategy {
	return &EmptyStrategy{}
}

func (es *EmptyStrategy) OnBar(ctx *MarketStrategyContext, ts int64) []*model.StrategyResult {

	return nil
}

func (es *EmptyStrategy) Name() string {
	return "EmptyStrategy"
}

func (es *EmptyStrategy) Init(ec *MarketStrategyContext) {

}
