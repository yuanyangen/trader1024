package analysis

//
//current_date = state["trade_date"]
//ticker = state["company_of_interest"]
//company_name = state["company_of_interest"]
//
//if toolkit.config["online_tools"]:
//tools = [toolkit.get_stock_news_ark]
//else:
//tools = [
//toolkit.get_reddit_stock_info,
//]
//
//system_message = (
//"你是一名专注于社交媒体与公司特定新闻的研究员/分析师，负责分析过去一周内，针对某家特定公司的社交媒体帖子、近期公司新闻以及公众情绪。你将获得一家公司的名称，你的目标是撰写一份详尽的长篇报告，通过分析该公司在社交媒体上的讨论内容、大众对该公司言论的看法、每日情绪数据，以及近期的公司新闻，深入挖掘该公司当前的状态，并为交易者和投资者提供详细的分析、洞察与潜在影响。"
//+ "请尽可能查阅所有可用信息来源，包括社交媒体、情绪数据、新闻报道等。不要只是笼统地指出“趋势复杂”或“好坏参半”，而要提供**详细且细粒度（fine-grained）的分析与洞察**，这些分析应有助于交易者做出更明智的决策。"
//+ "请务必在报告末尾附上一份 **Markdown 表格**，用于整理报告中的关键要点，使内容条理清晰、易于阅读。"
//)
//
//prompt = ChatPromptTemplate.from_messages(
//[
//(
//"system",
//"你是一个乐于助人的 AI 助手，正在与其他助手协同工作。"
//+"请利用所提供的工具，逐步推进问题的解答。"
//+"如果你无法完全回答问题，没有关系；其他拥有不同工具的助手"
//+"将会接替你未完成的部分。请尽力执行你所能做的，以推动任务进展。"
//+"如果你或任何其他助手得出了最终交易建议：**买入/持有/卖出（BUY/HOLD/SELL）** 或交付成果，"
//+"请在你的回复开头加上：FINAL TRANSACTION PROPOSAL: **BUY/HOLD/SELL**，以便团队知道可以停止。"
//+"你可使用的工具包括：{tool_names}。"
//+"{system_message}"
//+"供你参考，当前日期是 {current_date}。当前我们要分析的公司是 {ticker}。",
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
//
//result = chain.invoke(state["messages"])
//
//report = ""
//
//if len(result.tool_calls) == 0:
//report = result.content
//
//return {
//"messages": [result],
//"sentiment_report": report,
//}

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

var socialMediaUserMessageTPL = `"你是一个乐于助人的 AI 助手，正在与其他助手协同工作。
请利用所提供的工具，逐步推进问题的解答。
如果你无法完全回答问题，没有关系；其他拥有不同工具的助手
将会接替你未完成的部分。请尽力执行你所能做的，以推动任务进展。
如果你或任何其他助手得出了最终交易建议：**买入/持有/卖出（BUY/HOLD/SELL）** 或交付成果，
请在你的回复开头加上：FINAL TRANSACTION PROPOSAL: **BUY/HOLD/SELL**，以便团队知道可以停止。
你可使用的工具包括：{tool_names}。
{system_message}
供你参考，当前日期是 {current_date}。当前我们要分析的公司是 {ticker}。`

var socialMediaSystemMessageTPL = `你是一名专注于社交媒体与公司特定新闻的研究员/分析师，负责分析过去一周内，针对某家特定公司的社交媒体帖子、近期公司新闻以及公众情绪。你将获得一家公司的名称，你的目标是撰写一份详尽的长篇报告，通过分析该公司在社交媒体上的讨论内容、大众对该公司言论的看法、每日情绪数据，以及近期的公司新闻，深入挖掘该公司当前的状态，并为交易者和投资者提供详细的分析、洞察与潜在影响。
请尽可能查阅所有可用信息来源，包括社交媒体、情绪数据、新闻报道等。不要只是笼统地指出“趋势复杂”或“好坏参半”，而要提供**详细且细粒度（fine-grained）的分析与洞察**，这些分析应有助于交易者做出更明智的决策。
请务必在报告末尾附上一份 **Markdown 表格**，用于整理报告中的关键要点，使内容条理清晰、易于阅读。`

func SocialMediaChain[I, O any](ctx context.Context) *compose.Chain[I, O] {
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
	newsTool := SocialMediaTool()
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

func SocialMediaTool() tool.BaseTool {
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
		}, getSocialMediaForTool)
}

func getSocialMediaForTool(ctx context.Context, input *NewsRequest) (string, error) {
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一名研究员，负责分析某家公司过去一周的基本面信息。"),
		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", false),
		// 用户消息模板
		schema.UserMessage("请问您能否从 {curr_date} 前 7 天开始，到 {curr_date} 当天为止，在社交媒体上搜索股票代码为 {ticker} 的相关内容？请确保只获取该时间段内发布的帖子数据。"),
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
