package llm

import (
	"context"
	"fmt"

	"github.com/yuanyangen/trader1024/engine/engine/llm/analysis"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 调试函数，打印上下文内容
func PrintContextDebugInfo(ctx context.Context) {
	// 获取LLM上下文
	llmCtx := model.GetLLMContext(ctx)
	fmt.Printf("LLM Context: UniqId=%s, Date=%s\n", llmCtx.UniqId, llmCtx.Date)

	// 获取全局事件
	globalEvents := model.GetGlobalEvents(ctx)
	if globalEvents != nil {
		fmt.Printf("Global Events: %#v\n", globalEvents)
	} else {
		fmt.Println("Global Events: nil")
	}

	// 获取行业事件
	industryEvents := model.GetIndustryEvents(ctx)
	if industryEvents != nil {
		fmt.Printf("Industry Events: %#v\n", industryEvents)
	} else {
		fmt.Println("Industry Events: nil")
	}

	// 获取事件影响
	eventImpacts := model.GetEventImpacts(ctx)
	if eventImpacts != nil {
		fmt.Printf("Event Impacts: %#v\n", eventImpacts)
	} else {
		fmt.Println("Event Impacts: nil")
	}

	// 获取投资建议
	investmentRecommendation := model.GetInvestmentRecommendation(ctx)
	if investmentRecommendation != nil {
		fmt.Printf("Investment Recommendation: %#v\n", investmentRecommendation)
	} else {
		fmt.Println("Investment Recommendation: nil")
	}
}

// 测试事件收集器的调试函数
func TestEventCollector(ctx context.Context) error {
	agent := NewEventCollectorAgent(ctx)

	fmt.Println("Testing Global Event Collection...")
	globalEvents, err := agent.CollectGlobalEvents(ctx)
	if err != nil {
		return fmt.Errorf("Global event collection failed: %v", err)
	}
	fmt.Printf("Collected %d global events\n", len(globalEvents))
	for i, event := range globalEvents {
		fmt.Printf("  Event %d: %s\n", i+1, event.Title)
	}

	fmt.Println("\nTesting Industry Event Collection...")
	industry := "technology"
	industryEvents, err := agent.CollectIndustryEvents(ctx, industry)
	if err != nil {
		return fmt.Errorf("Industry event collection failed: %v", err)
	}
	fmt.Printf("Collected %d industry events for %s\n", len(industryEvents), industry)
	for i, event := range industryEvents {
		fmt.Printf("  Event %d: %s\n", i+1, event.Title)
	}

	return nil
}

// 测试新闻分析器
func TestNewsAnalysis(ctx context.Context) error {
	llmCtx := model.GetLLMContext(ctx)
	newsAnalyzer := analysis.NewNewsAnalysis(llmCtx)

	fmt.Println("Testing News Analysis - Global News...")
	globalNews, err := newsAnalyzer.GetGlobalNews(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get global news: %v", err)
	}

	fmt.Printf("Got %d global news items\n", len(globalNews))
	for i, newsItem := range globalNews {
		if newsItem != nil {
			fmt.Printf("  News %d: %s\n", i+1, newsItem.Title)
			fmt.Printf("     Source: %s, Relevance: %.2f\n", newsItem.Source, newsItem.Relevance)
		}
	}

	fmt.Println("\nTesting News Analysis - Industry News...")
	industry := "technology"
	industryNews, err := newsAnalyzer.GetIndustryNews(ctx, industry)
	if err != nil {
		return fmt.Errorf("Failed to get industry news: %v", err)
	}

	fmt.Printf("Got %d industry news items for %s\n", len(industryNews), industry)
	for i, newsItem := range industryNews {
		if newsItem != nil {
			fmt.Printf("  News %d: %s\n", i+1, newsItem.Title)
			fmt.Printf("     Source: %s, Relevance: %.2f\n", newsItem.Source, newsItem.Relevance)
		}
	}

	return nil
}
