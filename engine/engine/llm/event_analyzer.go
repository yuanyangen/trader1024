package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/yuanyangen/trader1024/config"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 事件影响分析结果
type EventImpact struct {
	EventID     string  `json:"event_id"`
	Industry    string  `json:"industry"`
	ImpactScore float64 `json:"impact_score"` // 影响分数 -1.0到1.0，负数为负面影响，正数为正面影响
	Confidence  float64 `json:"confidence"`   // 置信度 0-1
	Analysis    string  `json:"analysis"`     // 详细分析文本
}

// 事件分析器接口
type EventAnalyzer interface {
	AnalyzeEventImpact(ctx context.Context, events []Event, industry string) ([]EventImpact, error)
}

// 事件分析Agent实现
type EventAnalyzerAgent struct {
	llmCtx *model.LLMContext
}

func NewEventAnalyzerAgent(ctx context.Context) *EventAnalyzerAgent {
	llmCtx := model.GetLLMContext(ctx)
	return &EventAnalyzerAgent{
		llmCtx: llmCtx,
	}
}

// 事件影响分析链
func EventImpactAnalysisChain[In, Out any](ctx context.Context) compose.Chain[In, Out] {
	chain := compose.NewChain[In, Out]()
	chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input In) (Out, error) {
		agent := NewEventAnalyzerAgent(ctx)

		// 从上下文中获取事件
		globalEvents := model.GetGlobalEvents(ctx)
		industryEvents := model.GetIndustryEvents(ctx)

		// 从输入中获取行业信息
		inputMap, ok := any(input).(map[string]any)
		if !ok {
			return any(nil).(Out), fmt.Errorf("invalid input type for event impact analysis")
		}

		industry, ok := inputMap["industry"].(string)
		if !ok {
			return any(nil).(Out), fmt.Errorf("industry not specified in input")
		}

		// 合并所有事件
		var allEvents []Event
		if globalEvents != nil {
			if eventsSlice, ok := globalEvents.([]Event); ok {
				allEvents = append(allEvents, eventsSlice...)
			}
		}
		if industryEvents != nil {
			if eventsSlice, ok := industryEvents.([]Event); ok {
				allEvents = append(allEvents, eventsSlice...)
			}
		}

		// 分析事件影响
		impacts, err := agent.AnalyzeEventImpact(ctx, allEvents, industry)
		if err != nil {
			return any(nil).(Out), err
		}

		// 将分析结果存储到上下文中
		ctx = model.SetEventImpacts(ctx, impacts)
		return any(input).(Out), nil // 直接返回输入，保持类型一致
	}))
	return *chain
}

// 实际实现事件影响分析
func (a *EventAnalyzerAgent) AnalyzeEventImpact(ctx context.Context, events []Event, industry string) ([]EventImpact, error) {
	var impacts []EventImpact

	// 使用豆包大模型进行事件影响分析
	temperature := float32(0.7)
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		BaseURL:     config.GetArkBaseURL(),
		APIKey:      config.GetArkAPIKey(),
		Model:       config.GetArkModelName(),
		Temperature: &temperature,
	})
	if err != nil {
		return nil, fmt.Errorf("创建聊天模型失败: %v", err)
	}

	for _, event := range events {
		// 构建事件影响分析的提示词
		eventTypeStr := func() string {
			if event.Type == EventTypeGlobal {
				return "全球事件"
			} else if event.Type == EventTypeIndustry {
				return "行业事件"
			}
			return "其他事件"
		}()

		userMsg := fmt.Sprintf(`请分析以下事件对%s行业的影响：

事件标题：%s
事件类型：%s
事件摘要：%s
事件相关性：%.2f

请从以下几个方面进行分析：
1. 事件对行业的影响方向（正面/负面/中性）
2. 影响程度（影响分数范围：-1.0到1.0，负数为负面影响，正数为正面影响）
3. 分析的置信度（范围：0-1）
4. 详细的影响分析

请直接返回JSON格式结果，不要包含其他任何文本，字段包括：
- impact_score: 影响分数（-1.0到1.0）
- confidence: 置信度（0-1）
- analysis: 详细分析文本`, industry, event.Title, eventTypeStr, event.Description, event.Relevance)

		template := prompt.FromMessages(schema.FString,
			schema.SystemMessage("你是一名专业的金融分析师，负责分析事件对特定行业的影响。请严格按照要求的格式返回结果。"),
			schema.UserMessage(userMsg),
		)

		// 格式化提示词
		messages, err := template.Format(ctx, map[string]any{})
		if err != nil {
			fmt.Printf("格式化提示词失败: %v，跳过事件 %s\n", err, event.ID)
			continue
		}

		// 调用豆包大模型进行分析
		result, err := chatModel.Generate(ctx, messages)
		if err != nil {
			fmt.Printf("调用LLM模型失败: %v，跳过事件 %s\n", err, event.ID)
			continue
		}

		// 解析LLM响应
		var impactResult struct {
			ImpactScore float64 `json:"impact_score"`
			Confidence  float64 `json:"confidence"`
			Analysis    string  `json:"analysis"`
		}

		err = json.Unmarshal([]byte(result.Content), &impactResult)
		if err != nil {
			// 如果解析JSON失败，使用默认值
			fmt.Printf("解析LLM响应失败: %v，使用默认值，跳过事件 %s\n", err, event.ID)

			// 使用简化的影响分析逻辑作为备用方案
			impactScore := 0.0
			if event.Type == EventTypeGlobal {
				impactScore = 0.3
			} else if event.Type == EventTypeIndustry {
				impactScore = 0.7
			}

			impacts = append(impacts, EventImpact{
				EventID:     event.ID,
				Industry:    industry,
				ImpactScore: impactScore * event.Relevance,
				Confidence:  0.85,
				Analysis:    fmt.Sprintf("事件 '%s' 可能对%s行业产生影响，影响评分为%.2f", event.Title, industry, impactScore*event.Relevance),
			})
			continue
		}

		// 确保影响分数在有效范围内
		if impactResult.ImpactScore < -1.0 {
			impactResult.ImpactScore = -1.0
		} else if impactResult.ImpactScore > 1.0 {
			impactResult.ImpactScore = 1.0
		}

		// 确保置信度在有效范围内
		if impactResult.Confidence < 0 {
			impactResult.Confidence = 0
		} else if impactResult.Confidence > 1 {
			impactResult.Confidence = 1
		}

		// 计算最终影响分数（考虑事件相关性）
		finalImpactScore := impactResult.ImpactScore * event.Relevance

		impacts = append(impacts, EventImpact{
			EventID:     event.ID,
			Industry:    industry,
			ImpactScore: finalImpactScore,
			Confidence:  impactResult.Confidence,
			Analysis:    impactResult.Analysis,
		})
	}

	return impacts, nil
}
