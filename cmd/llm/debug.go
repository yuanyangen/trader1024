package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yuanyangen/trader1024/engine/engine/llm"
	"github.com/yuanyangen/trader1024/engine/model"
)

func runDebug() {
	// 创建上下文
	ctx := context.Background()

	// 初始化LLM上下文
	llmCtx := &model.LLMContext{
		UniqId: "test-run-1",
		Date:   time.Now().Format("2006-01-02"),
	}
	ctx = model.SetLLmContext(ctx, llmCtx)

	fmt.Println("=== 开始调试LLM AI投资助手 ===")

	// 1. 测试新闻分析器
	fmt.Println("\n1. 测试新闻分析器:")
	err := llm.TestNewsAnalysis(ctx)
	if err != nil {
		log.Fatalf("测试新闻分析器失败: %v", err)
	}

	// 2. 测试事件收集器
	fmt.Println("\n2. 测试事件收集器:")
	err = llm.TestEventCollector(ctx)
	if err != nil {
		log.Fatalf("测试事件收集器失败: %v", err)
	}

	// 3. 执行完整的投资分析
	fmt.Println("\n3. 执行完整的投资分析:")
	industry := "technology"
	fmt.Printf("分析 %s 行业...\n", industry)

	// 执行投资分析
	investmentRecommendation, err := llm.RunInvestmentAnalysis(ctx, industry)
	if err != nil {
		log.Fatalf("投资分析失败: %v", err)
	}

	// 打印结果
	fmt.Printf("综合评分: %.2f\n", investmentRecommendation.OverallScore)
	fmt.Printf("建议操作: %s\n", investmentRecommendation.RecommendAction)
	fmt.Printf("置信度: %.1f%%\n", investmentRecommendation.Confidence*100)
	fmt.Printf("分析: %s\n", investmentRecommendation.Analysis)

	if len(investmentRecommendation.SupportingEvents) > 0 {
		fmt.Println("支持事件:")
		for i, event := range investmentRecommendation.SupportingEvents {
			fmt.Printf("  %d. %s (影响分数: %.2f, 置信度: %.1f%%)\n", i+1, event.Analysis, event.ImpactScore, event.Confidence*100)
		}
	}

	fmt.Println("\n=== 调试完成 ===")
}

// 使用 go run cmd/llm/debug.go cmd/llm/... 运行时，需要指定正确的文件名
// 或者直接修改 cmd/llm/main.go 调用 runDebug 函数
