package strategy

type EmptyStrategy struct {
}

func NewEmptyStrategy() *EmptyStrategy {
	return &EmptyStrategy{}
}

func (es *EmptyStrategy) OnBar(ctx *MarketStrategyContext, ts int64) {

}
