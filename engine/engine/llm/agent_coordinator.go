package llm

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/compose"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 投资建议结果
type InvestmentRecommendation struct {
	Industry         string        `json:"industry"`
	OverallScore     float64       `json:"overall_score"`    // 综合评分 -1.0到1.0
	Confidence       float64       `json:"confidence"`       // 置信度 0-1
	RecommendAction  string        `json:"recommend_action"` // "buy", "sell", "hold"
	SupportingEvents []EventImpact `json:"supporting_events"`
	Analysis         string        `json:"analysis"`
}

// Agent协调器
type AgentCoordinator struct {
	llmCtx *model.LLMContext
}

func NewAgentCoordinator(ctx context.Context) *AgentCoordinator {
	llmCtx := model.GetLLMContext(ctx)
	return &AgentCoordinator{
		llmCtx: llmCtx,
	}
}

// 投资建议生成链
func RecommendationChain[In, Out any](ctx context.Context) compose.Chain[In, Out] {
	chain := compose.NewChain[In, Out]()

	// 使用Lambda来实现自定义逻辑
	chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input In) (Out, error) {
		// 从上下文中获取事件影响分析结果
		impacts := model.GetEventImpacts(ctx)

		// 从输入中获取行业信息
		inputMap, ok := any(input).(map[string]any)
		if !ok {
			return any(nil).(Out), fmt.Errorf("invalid input type for recommendation")
		}

		industry, ok := inputMap["industry"].(string)
		if !ok {
			return any(nil).(Out), fmt.Errorf("industry not specified in input")
		}

		// 计算综合评分
		overallScore := 0.0
		var impactList []EventImpact

		if impacts != nil {
			switch v := impacts.(type) {
			case []EventImpact:
				impactList = v
				for _, impact := range v {
					overallScore += impact.ImpactScore * impact.Confidence
				}
				fmt.Printf("成功解析事件影响: %d个事件, 综合评分: %.2f\n", len(v), overallScore)
			default:
				fmt.Printf("无法解析事件影响类型: %T, 原始值: %#v\n", impacts, impacts)
			}
		} else {
			fmt.Println("事件影响分析结果为nil")
		}

		// 生成建议操作
		recommendAction := "hold"
		if overallScore > 0.3 {
			recommendAction = "buy"
		} else if overallScore < -0.3 {
			recommendAction = "sell"
		}

		// 创建投资建议
		recommendation := InvestmentRecommendation{
			Industry:         industry,
			OverallScore:     overallScore,
			Confidence:       0.8, // 示例置信度
			RecommendAction:  recommendAction,
			SupportingEvents: impactList,
			Analysis:         fmt.Sprintf("基于事件分析，对%s行业的综合评分为%.2f，建议%s", industry, overallScore, recommendAction),
		}

		// 将建议存储到上下文中
		ctx = model.SetInvestmentRecommendation(ctx, recommendation)

		fmt.Printf("生成的投资建议: %#v\n", recommendation)

		return any(&recommendation).(Out), nil
	}))

	return *chain
}

// 执行投资分析流程（直接调用Agent方法）
func RunInvestmentAnalysis(ctx context.Context, industry string) (*InvestmentRecommendation, error) {
	fmt.Println("=== 开始顺序执行投资分析流程 ===")

	// 1. 收集全球事件
	fmt.Println("1. 收集全球事件...")
	globalCollector := NewEventCollectorAgent(ctx)
	globalEvents, err := globalCollector.CollectGlobalEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("全球事件收集失败: %v", err)
	}
	ctx = model.SetGlobalEvents(ctx, globalEvents)
	fmt.Printf("   收集到 %d 个全球事件\n", len(globalEvents))

	// 2. 收集行业事件
	fmt.Println("2. 收集行业事件...")
	industryCollector := NewEventCollectorAgent(ctx)
	industryEvents, err := industryCollector.CollectIndustryEvents(ctx, industry)
	if err != nil {
		return nil, fmt.Errorf("行业事件收集失败: %v", err)
	}
	ctx = model.SetIndustryEvents(ctx, industryEvents)
	fmt.Printf("   收集到 %d 个行业事件\n", len(industryEvents))

	// 3. 分析事件影响
	fmt.Println("3. 分析事件影响...")
	analyzer := NewEventAnalyzerAgent(ctx)
	impacts, err := analyzer.AnalyzeEventImpact(ctx, append(globalEvents, industryEvents...), industry)
	if err != nil {
		return nil, fmt.Errorf("事件影响分析失败: %v", err)
	}
	ctx = model.SetEventImpacts(ctx, impacts)
	fmt.Printf("   分析出 %d 个事件影响\n", len(impacts))

	// 4. 生成投资建议
	fmt.Println("4. 生成投资建议...")
	var overallScore float64
	for _, impact := range impacts {
		overallScore += impact.ImpactScore * impact.Confidence
	}

	recommendAction := "hold"
	if overallScore > 0.3 {
		recommendAction = "buy"
	} else if overallScore < -0.3 {
		recommendAction = "sell"
	}

	recommendation := &InvestmentRecommendation{
		Industry:         industry,
		OverallScore:     overallScore,
		Confidence:       0.8, // 示例置信度
		RecommendAction:  recommendAction,
		SupportingEvents: impacts,
		Analysis:         fmt.Sprintf("基于事件分析，对%s行业的综合评分为%.2f，建议%s", industry, overallScore, recommendAction),
	}

	fmt.Println("=== 投资分析流程执行完成 ===")

	return recommendation, nil
}
