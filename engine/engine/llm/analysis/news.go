//
//current_date = state["trade_date"]
//ticker = state["company_of_interest"]
//
//tools = [toolkit.get_global_news_ark, toolkit.get_google_news]

//system_message = (
//"你是一名新闻研究员，负责分析过去一周内的最新新闻与趋势。请撰写一份全面的报告，阐述当前与交易和宏观经济相关的全球形势。请查阅来自 EODHD 和 finnhub 的新闻，以确保信息全面。不要仅仅笼统地说趋势是“混合的”，而要提供详细、细粒度化的分析与洞察，这些分析应有助于交易者做出决策。"
//+"请务必在报告末尾附上一张 **Markdown 表格**，用于整理报告中的关键要点，要求条理清晰、易于阅读。"
//)
//
//prompt = ChatPromptTemplate.from_messages(
//[
//(
//"system",
//"你是一个乐于助人的 AI 助手，正在与其他助手协同工作。"
//+ "请利用所提供的工具，逐步推进问题的解答。"
//+ "如果你无法完全回答问题，没有关系；其他拥有不同工具的助手"
//+ "将会接替你未完成的部分。请尽力执行你所能做的，以推动任务进展。"
//+ "如果你或任何其他助手得出了最终交易建议：**买入/持有/卖出（BUY/HOLD/SELL）** 或交付成果，"
//+ "请在你的回复开头加上：FINAL TRANSACTION PROPOSAL: **BUY/HOLD/SELL**，以便团队知道可以停止。"
//+ "你可使用的工具包括：{tool_names}。"
//+ "{system_message}"
//+ "供你参考，当前日期是 {current_date}。我们正在分析的公司是 {ticker}。"
//),
//MessagesPlaceholder(variable_name="messages"),
//]
//)
//
//prompt = prompt.partial(system_message=system_message)
//prompt = prompt.partial(tool_names=", ".join([tool.name for tool in tools]))
//prompt = prompt.partial(current_date=current_date)
//prompt = prompt.partial(ticker=ticker)
//
//chain = prompt | llm.bind_tools(tools)
//result = chain.invoke(state["messages"])
//
//report = ""
//
//if len(result.tool_calls) == 0:
//report = result.content
//
//return {
//"messages": [result],
//"news_report": report,
//}

package analysis

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/yuanyangen/trader1024/config"
	"github.com/yuanyangen/trader1024/engine/model"
)

var newsUserMessageTPL = `你是一个乐于助人的 AI 助手，正在与其他助手协同工作。
请利用所提供的工具，逐步推进问题的解答。
如果你无法完全回答问题，没有关系；其他拥有不同工具的助手
将会接替你未完成的部分。请尽力执行你所能做的，以推动任务进展。
如果你或任何其他助手得出了最终交易建议：**买入/持有/卖出（BUY/HOLD/SELL）** 或交付成果，
请在你的回复开头加上：FINAL TRANSACTION PROPOSAL: **BUY/HOLD/SELL**，以便团队知道可以停止。
你可使用的工具包括：{tool_names}。
{system_message}
供你参考，当前日期是 {current_date}。我们正在分析的公司是 {ticker}。`

var newsSystemMessageTPL = `你是一名新闻研究员，负责分析过去一周内的最新新闻与趋势。请撰写一份全面的报告，阐述当前与交易和宏观经济相关的全球形势。请查阅来自 EODHD 和 finnhub 的新闻，以确保信息全面。不要仅仅笼统地说趋势是“混合的”，而要提供详细、细粒度化的分析与洞察，这些分析应有助于交易者做出决策。
请务必在报告末尾附上一张 **Markdown 表格**，用于整理报告中的关键要点，要求条理清晰、易于阅读。`

func NewsChain[I, O any](ctx context.Context) *compose.Chain[I, O] {
	//callbacks.AppendGlobalHandlers(&loggerCallbacks{})
	// 2. create an instance of ChatModel as 2nd Graph Node
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
	newsTool := NewsTool()
	info, err := newsTool.Info(ctx)
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	// 4. bind ToolInfo to ChatModel. ToolInfo will remain in effect until the next BindTools.
	err = chatModel.BindTools([]*schema.ToolInfo{info})
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	// 5. create an instance of ToolsNode as 3rd Graph Node
	toolsNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: []tool.BaseTool{newsTool},
	})
	if err != nil {
		panic("should not reach here" + err.Error())
	}

	// 1. create an instance of ChatTemplate as 1st Graph Node
	chatTpl := prompt.FromMessages(schema.FString,
		schema.SystemMessage(newsSystemMessageTPL),
		schema.MessagesPlaceholder("message_histories", true),
		schema.UserMessage(newsUserMessageTPL),
	)
	// TODO
	chatTpl.Format(ctx, map[string]any{})
	c := compose.NewChain[I, O]().
		AppendChatTemplate(chatTpl).
		AppendChatModel(chatModel).
		AppendToolsNode(toolsNode)
	return c
}

func NewsTool() tool.BaseTool {
	return utils.NewTool(
		&schema.ToolInfo{
			Name: "get_fundamental",
			Desc: "根据时期，获取最新的宏观经济信息",
			ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
				"date": {
					Type: "date",
					Desc: "当前时期",
				},
			}),
		}, getNewsForTool)
}

type NewsRequest struct {
	UniqId string `json:"uniq_id"`
	Date   string `json:"date"`
}
type NewsResponse struct {
	Data string
}

func getNewsForTool(ctx context.Context, input *NewsRequest) (string, error) {
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一名研究员，负责分析某家公司过去一周的基本面信息。"),
		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", false),
		// 用户消息模板
		schema.UserMessage("是否可以搜索从 {date} 前 7 天至 {date} 当天 的全球或宏观经济新闻，这些新闻应对交易决策具有参考价值？"),
	)

	llmCtx := model.GetLLMContext(ctx)

	// 使用模板生成消息
	messages, err := template.Format(ctx, map[string]any{
		"date": llmCtx.Date,
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
	return result.Content, nil
}
