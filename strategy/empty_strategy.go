package strategy

import (
	"github.com/yuanyangen/trader1024/engine/model"
)

type EmptyStrategy struct {
}

func NewEmptyStrategy() *EmptyStrategy {
	return &EmptyStrategy{}
}

func (es *EmptyStrategy) OnBar(ctx *model.MarketStrategyContext) {

}
