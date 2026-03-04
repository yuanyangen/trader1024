package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/yuanyangen/trader1024/engine/engine/llm"
	"github.com/yuanyangen/trader1024/engine/model"
)

func main() {
	// 创建上下文
	ctx := context.Background()

	// 初始化LLM上下文
	llmCtx := &model.LLMContext{
		UniqId: "test-run-1",
		Date:   time.Now().Format("2006-01-02"),
	}
	ctx = model.SetLLmContext(ctx, llmCtx)

	// 示例：分析科技行业的投资机会
	industry := "technology"

	// 执行投资分析
	fmt.Printf("正在分析 %s 行业的投资机会...\n\n", industry)

	recommendation, err := llm.RunInvestmentAnalysis(ctx, industry)
	if err != nil {
		log.Fatalf("分析失败: %v", err)
	}

	// 打印分析结果
	fmt.Printf("行业：%s\n", recommendation.Industry)
	fmt.Printf("综合评分：%.2f\n", recommendation.OverallScore)
	fmt.Printf("置信度：%.1f%%\n", recommendation.Confidence*100)
	fmt.Printf("建议操作：%s\n", recommendation.RecommendAction)
	fmt.Printf("分析：%s\n\n", recommendation.Analysis)

	// 打印支持性事件
	fmt.Println("支持性事件：")
	for _, impact := range recommendation.SupportingEvents {
		fmt.Printf("- 事件ID：%s\n", impact.EventID)
		fmt.Printf("  影响评分：%.2f (置信度：%.1f%%)\n", impact.ImpactScore, impact.Confidence*100)
		fmt.Printf("  分析：%s\n\n", impact.Analysis)
	}

	// 输出JSON格式结果
	jsonResult, err := json.MarshalIndent(recommendation, "", "  ")
	if err != nil {
		log.Fatalf("JSON编码失败: %v", err)
	}

	fmt.Println("JSON格式结果：")
	fmt.Println(string(jsonResult))
}
