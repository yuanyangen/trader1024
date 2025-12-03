package analysis

import (
	"context"
	"github.com/yuanyangen/trader1024/engine/model"
	"log"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/yuanyangen/trader1024/config"
)

var fundamentalsMessage = `"你是一个乐于助人的 AI 助手，正在与其他助手协同工作。
请利用所提供的工具，逐步推进问题的解答。
如果你无法完全回答问题，没有关系；其他拥有不同工具的助手
将会接替你未完成的部分。请尽力执行你所能做的，以推动任务进展。
如果你或任何其他助手得出了最终交易建议：**买入/持有/卖出（BUY/HOLD/SELL）** 或交付成果，
请在你的回复开头加上：FINAL TRANSACTION PROPOSAL: **BUY/HOLD/SELL**，以便团队知道可以停止推进。
你可使用的工具包括：{tool_names}。
{system_message}
供你参考，当前日期是 {current_date}。我们要分析的公司是 {ticker}。

`

var fundamentalsSystemMessageTpl = `你是一名研究员，负责分析某家公司过去一周的基本面信息。请撰写一份关于该公司基本面信息的综合报告，内容包括但不限于：财务文件、公司概况、基础财务数据、公司财务历史、内部人士情绪以及内部人士交易记录，以便全面了解该公司的基本面情况，为交易者提供决策依据。
请确保报告包含尽可能详实的信息。不要笼统地描述趋势为“好坏参半'，而要提供详细、细致、粒度精细的分析与洞察，这些内容应有助于交易者做出更为明智的决策。
请务必在报告的末尾附上一份 Markdown 格式的表格，用于归纳和整理报告中的关键要点，要求表格结构清晰、内容条理分明、便于阅读。
`

func FundamentalChain[I, O any](ctx context.Context) *compose.Chain[I, O] {
	temperature := float32(0.7)
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		BaseURL:     config.GetArkBaseURL(),
		APIKey:      config.GetArkAPIKey(),
		Model:       config.GetArkModelName(),
		Temperature: &temperature,
	})
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	// 3. create an instance of tool.InvokableTool for Intent recognition and execution
	fundamentalTool := FundamentalTool()
	fundamentalToolObject, err := fundamentalTool.Info(ctx)
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	// 4. bind ToolInfo to ChatModel. ToolInfo will remain in effect until the next BindTools.
	err = chatModel.BindTools([]*schema.ToolInfo{fundamentalToolObject})
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	// 5. create an instance of ToolsNode as 3rd Graph Node
	toolsNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: []tool.BaseTool{fundamentalTool},
	})
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	// 1. create an instance of ChatTemplate as 1st Graph Node
	chatTpl := prompt.FromMessages(schema.FString,
		schema.SystemMessage(fundamentalsSystemMessageTpl),
		schema.MessagesPlaceholder("message_histories", true),
		schema.UserMessage(fundamentalsMessage),
	)
	//llmCtx := model.GetLLMContext(ctx)
	////err := chatTpl.Format(ctx, map[string]any{
	////	"uniq_id": llmCtx.UniqId,
	////	"date":    llmCtx.Date,
	////})

	c := compose.NewChain[I, O]().
		AppendChatTemplate(chatTpl).
		AppendChatModel(chatModel).
		AppendToolsNode(toolsNode)

	return c
}

func FundamentalTool() tool.BaseTool {
	return utils.NewTool(
		&schema.ToolInfo{
			Name: "get_fundamental",
			Desc: "根据唯一ID（比如股票ID）和时期，获取过去100天内的日维度的k线数据",
			ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
				"name": {
					Type: "ticker",
					Desc: "唯一交易物的ID",
				},
				"email": {
					Type: "curr_date",
					Desc: "当前时期",
				},
			}),
		}, getFundamentalForTool)
}

type FundamentalRequest struct {
	UniqId string `json:"uniq_id"`
	Date   string `json:"date"`
}
type FundamentalResponse struct {
	Data string
}

func getFundamentalForTool(ctx context.Context, input *FundamentalRequest) (output *FundamentalResponse, err error) {
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一名研究员，负责分析某家公司过去一周的基本面信息。"),
		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", false),
		// 用户消息模板
		schema.UserMessage("能否查找 {uniq_id} 在 {date} 所在月份的前一个月 至 {date} 所在月份期间，关于其基本面（Fundamental）的讨论内容。请确保只获取该时间段内发布的数据。并以表格形式列出，包含市盈率（PE）、市销率（PS）、现金流等指标。"),
	)

	llmCtx := model.GetLLMContext(ctx)

	// 使用模板生成消息
	messages, err := template.Format(ctx, map[string]any{
		"uniq_id": llmCtx.UniqId,
		"date":    llmCtx.Date,
	})
	if err != nil {
		log.Fatalf("format template failed: %v\n", err)
	}

	temperature := float32(0.7)

	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		BaseURL:     config.GetArkBaseURL(),
		APIKey:      config.GetArkAPIKey(),
		Model:       config.GetArkModelName(),
		Temperature: &temperature,
	})
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	result, err := chatModel.Generate(ctx, messages)
	if err != nil {

		panic("should not reach here" + err.Error())
	}

	return &FundamentalResponse{
		Data: result.Content,
	}, nil
}
