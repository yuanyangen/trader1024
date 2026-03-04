package model

import "context"

type LLMContext struct {
	UniqId      string
	Date        string
	TradeObject *TradeObject
	Line        *KLine
}

const LLMContextKey = "LLMContextKey"

// 事件相关的上下文键
const (
	GlobalEventsKey   = "GlobalEventsKey"
	IndustryEventsKey = "IndustryEventsKey"
	EventImpactsKey   = "EventImpactsKey"
)

func GetLLMContext(ctx context.Context) *LLMContext {
	oi := ctx.Value(LLMContextKey)
	if oi == nil {
		panic("LLMContext not found in context, should not reach here")
	}
	o, ok := oi.(*LLMContext)
	if !ok {
		panic("LLMContext not found in context, should not reach here")
	}

	return o
}

func SetLLmContext(ctx context.Context, o *LLMContext) context.Context {
	if o == nil {
		panic("LLMContext is nil, should not reach here")
	}
	return context.WithValue(ctx, LLMContextKey, o)
}

// 全局事件上下文操作
func SetGlobalEvents(ctx context.Context, events interface{}) context.Context {
	return context.WithValue(ctx, GlobalEventsKey, events)
}

func GetGlobalEvents(ctx context.Context) interface{} {
	return ctx.Value(GlobalEventsKey)
}

// 行业事件上下文操作
func SetIndustryEvents(ctx context.Context, events interface{}) context.Context {
	return context.WithValue(ctx, IndustryEventsKey, events)
}

func GetIndustryEvents(ctx context.Context) interface{} {
	return ctx.Value(IndustryEventsKey)
}

// 事件影响分析上下文操作
func SetEventImpacts(ctx context.Context, impacts interface{}) context.Context {
	return context.WithValue(ctx, EventImpactsKey, impacts)
}

func GetEventImpacts(ctx context.Context) interface{} {
	return ctx.Value(EventImpactsKey)
}

// 投资建议上下文操作
const InvestmentRecommendationKey = "InvestmentRecommendationKey"

func SetInvestmentRecommendation(ctx context.Context, recommendation interface{}) context.Context {
	return context.WithValue(ctx, InvestmentRecommendationKey, recommendation)
}

func GetInvestmentRecommendation(ctx context.Context) interface{} {
	return ctx.Value(InvestmentRecommendationKey)
}
