package llm

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/compose"
	"github.com/yuanyangen/trader1024/engine/engine/llm/analysis"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 事件类型定义
type EventType string

const (
	EventTypeGlobal   EventType = "global"   // 全球大事件
	EventTypeIndustry EventType = "industry" // 行业事件
)

// 事件数据结构
type Event struct {
	ID          string    `json:"id"`
	Type        EventType `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Timestamp   int64     `json:"timestamp"`
	Source      string    `json:"source"`
	Relevance   float64   `json:"relevance"` // 相关性分数 0-1
}

// 事件收集器接口
type EventCollector interface {
	CollectGlobalEvents(ctx context.Context) ([]Event, error)
	CollectIndustryEvents(ctx context.Context, industry string) ([]Event, error)
}

// 事件收集Agent实现
type EventCollectorAgent struct {
	llmCtx       *model.LLMContext
	newsAnalyzer *analysis.NewsAnalysis
}

func NewEventCollectorAgent(ctx context.Context) *EventCollectorAgent {
	llmCtx := model.GetLLMContext(ctx)
	return &EventCollectorAgent{
		llmCtx:       llmCtx,
		newsAnalyzer: analysis.NewNewsAnalysis(llmCtx),
	}
}

// 全球事件收集链
func GlobalEventCollectorChain[In, Out any](ctx context.Context) compose.Chain[In, Out] {
	chain := compose.NewChain[In, Out]()
	chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input In) (Out, error) {
		agent := NewEventCollectorAgent(ctx)
		events, err := agent.CollectGlobalEvents(ctx)
		if err != nil {
			return any(nil).(Out), err
		}

		// 将事件存储到上下文中
		ctx = model.SetGlobalEvents(ctx, events)
		return any(input).(Out), nil // 直接返回输入，保持类型一致
	}))
	return *chain
}

// 行业事件收集链
func IndustryEventCollectorChain[In, Out any](ctx context.Context) compose.Chain[In, Out] {
	chain := compose.NewChain[In, Out]()
	chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input In) (Out, error) {
		agent := NewEventCollectorAgent(ctx)

		// 从输入中获取行业信息
		inputMap, ok := any(input).(map[string]any)
		if !ok {
			return any(nil).(Out), fmt.Errorf("invalid input type for industry event collector")
		}

		industry, ok := inputMap["industry"].(string)
		if !ok {
			return any(nil).(Out), fmt.Errorf("industry not specified in input")
		}

		events, err := agent.CollectIndustryEvents(ctx, industry)
		if err != nil {
			return any(nil).(Out), err
		}

		// 将事件存储到上下文中
		ctx = model.SetIndustryEvents(ctx, events)
		return any(input).(Out), nil // 直接返回输入，保持类型一致
	}))
	return *chain
}

// 实际实现全球事件收集
func (a *EventCollectorAgent) CollectGlobalEvents(ctx context.Context) ([]Event, error) {
	// 使用新闻分析器获取全球重要事件
	newsItems, err := a.newsAnalyzer.GetGlobalNews(ctx)
	if err != nil {
		return nil, err
	}

	var events []Event
	for i, item := range newsItems {
		if item != nil {
			events = append(events, Event{
				ID:          fmt.Sprintf("global-%d", i+1),
				Type:        EventTypeGlobal,
				Title:       item.Title,
				Description: item.Description,
				Timestamp:   item.Timestamp.Unix(),
				Source:      item.Source,
				Relevance:   item.Relevance,
			})
		}
	}

	return events, nil
}

// 实际实现行业事件收集
func (a *EventCollectorAgent) CollectIndustryEvents(ctx context.Context, industry string) ([]Event, error) {
	// 使用新闻分析器获取行业相关事件
	industryNewsItems, err := a.newsAnalyzer.GetIndustryNews(ctx, industry)
	if err != nil {
		return nil, err
	}

	var events []Event
	for i, item := range industryNewsItems {
		if item != nil {
			events = append(events, Event{
				ID:          fmt.Sprintf("industry-%s-%d", industry, i+1),
				Type:        EventTypeIndustry,
				Title:       item.Title,
				Description: item.Description,
				Timestamp:   item.Timestamp.Unix(),
				Source:      item.Source,
				Relevance:   item.Relevance,
			})
		}
	}

	return events, nil
}
