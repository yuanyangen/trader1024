package engine

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/engine/llm/analysis"
	local_broker2 "github.com/yuanyangen/trader1024/engine/local_broker"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 处理某个具体的合约
type LLMTradeObjectExecuteEngine struct {
	TradeObject *model.TradeObject
	dataSource  model.DateSource
	Line        *model.KLine

	Strategies        []model.Strategy
	portfolioStrategy []model.PortfolioStrategy
	Brokers           []model.Broker
}

func (m *LLMTradeObjectExecuteEngine) DealEvent(event *model.EventMsg) {
	if event == nil || event.TimeStamp == 0 {
		return
	}
	if m.dataSource == nil {
		panic("no data_crawler source")
	}

	if (m.TradeObject.ContractEndTime != 0 && event.TimeStamp >= m.TradeObject.ContractEndTime) ||
		(event.TimeStamp < m.TradeObject.ContractStartTime && m.TradeObject.ContractStartTime != 0) {
		return
	}
	dataNode := m.dataSource.GetDataByTs(context.Background(), m.TradeObject.UniqueCode, model.LineType_Day, event.TimeStamp)
	if dataNode == nil {
		return
	}
	ctx := &model.TradeObjectEngineContext{
		TradeObject:  m.TradeObject,
		Kline:        m.Line,
		CurrentKNode: dataNode,
		Ts:           event.TimeStamp,
	}
	m.Line.AddNodeData(event.TimeStamp, dataNode)
	m.runStrategies(ctx)
	m.runPortfolioStrategies(ctx)
	m.executeBuySellCmd(ctx)
}

func (m *LLMTradeObjectExecuteEngine) runStrategies(ctx *model.TradeObjectEngineContext) {
	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx)
		if stResult != nil {
			stResult.StrategyName = st.Name()
			ctx.StrategyResult = append(ctx.StrategyResult, stResult)
		}
	}
}
func (m *LLMTradeObjectExecuteEngine) runPortfolioStrategies(ctx *model.TradeObjectEngineContext) {
	if len(ctx.StrategyResult) == 0 {
		return
	}
	for _, p := range m.portfolioStrategy {
		p(ctx)
	}
}

func (m *LLMTradeObjectExecuteEngine) executeBuySellCmd(ctx *model.TradeObjectEngineContext) {
	for _, order := range ctx.Orders {
		for _, broker := range m.Brokers {
			broker.AddOrder(order)
		}
	}
}

func (m *LLMTradeObjectExecuteEngine) Name() string {
	return m.TradeObject.ContractCnName + m.TradeObject.ContractDate
}

func (m *LLMTradeObjectExecuteEngine) DoPlot(p *charts.Page) {
	position := local_broker2.GetLocalBroker().GetPositionsByContract(m.TradeObject) //????
	position.ReportToCmd()
	m.Line.DoPlot(p)
}

func InitLLM(ctx context.Context, ticket string, date string) {
	const (
		nodeKeyFundamental = "fundamental"
		nodeKeyMarket      = "market"
		nodeKeyNews        = "news"
		nodeKeySocialMedia = "social_media"
	)

	ctx = model.SetLLmContext(ctx, &model.LLMContext{
		UniqId: ticket,
		Date:   date,
	})

	g := compose.NewGraph[map[string]any, []*schema.Message]()

	_ = g.AddGraphNode(nodeKeyMarket, analysis.MarketChain[map[string]any, []*schema.Message](ctx))
	_ = g.AddGraphNode(nodeKeyFundamental, analysis.FundamentalChain[map[string]any, []*schema.Message](ctx))
	_ = g.AddGraphNode(nodeKeyNews, analysis.NewsChain[map[string]any, []*schema.Message](ctx))
	_ = g.AddGraphNode(nodeKeySocialMedia, analysis.SocialMediaChain[map[string]any, []*schema.Message](ctx))

	_ = g.AddEdge(compose.START, nodeKeyMarket)
	_ = g.AddEdge(nodeKeyMarket, nodeKeyFundamental)
	_ = g.AddEdge(nodeKeyFundamental, nodeKeyNews)
	_ = g.AddEdge(nodeKeyNews, nodeKeySocialMedia)
	_ = g.AddEdge(nodeKeySocialMedia, compose.END)
	r, err := g.Compile(ctx)
	if err != nil {
		return
	}

	out, err := r.Invoke(ctx, map[string]any{
		"message_histories": []*schema.Message{},
		"tiket":             ticket,
		"user_query":        "我叫 zhangsan, 邮箱是 zhangsan@bytedance.com, 帮我推荐一处房产",
	})
	fmt.Println(out)

}
