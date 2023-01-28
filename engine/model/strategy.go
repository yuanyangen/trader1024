package model

type Strategy interface {
	//Init(*EngineContext)
	OnBar(ctx *MarketStrategyContext)
}
