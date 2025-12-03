package model

import "context"

type LLMContext struct {
	UniqId      string
	Date        string
	TradeObject *TradeObject
	Line        *KLine
}

const LLMContextKey = "LLMContextKey"

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
