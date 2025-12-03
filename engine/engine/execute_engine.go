package engine

import (
	"context"

	"github.com/go-echarts/go-echarts/charts"
	local_broker2 "github.com/yuanyangen/trader1024/engine/local_broker"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type ExecuteEngine struct {
	*baseEngine

	ContractEngines   map[string]*TradeObjectExecuteEngine
	strategies        []func() model.Strategy
	portfolioStrategy []model.PortfolioStrategy
	brokers           []model.Broker
}

type CmdExecutor interface {
	ExecuteCmd(req *model.TradeObjectEngineContext)
	Report()
}

func NewEngine(et model.EventTrigger, dataSource model.DateSource, strategies []func() model.Strategy, portfolioStrategy []model.PortfolioStrategy, brokers []model.Broker) *ExecuteEngine {
	e := &ExecuteEngine{
		baseEngine: &baseEngine{
			TradeObjects:   map[string]*model.TradeObject{},
			EventTrigger:   et,
			dataSource:     dataSource,
			watcherBackend: NewPlotterServers(),
		},
		strategies:        strategies,
		portfolioStrategy: portfolioStrategy,
		ContractEngines:   map[string]*TradeObjectExecuteEngine{},
		brokers:           append([]model.Broker{local_broker2.GetLocalBroker()}, brokers...),
	}
	return e
}

func (ec *ExecuteEngine) Start() error {
	for _, tradeObject := range ec.TradeObjects {
		strategies := make([]model.Strategy, len(ec.strategies))
		for i, stFactory := range ec.strategies {
			strategies[i] = stFactory()
		}

		kline := model.NewKLine(tradeObject.CNName+tradeObject.ContractDate, model.LineType_Day)
		if strategies != nil {
			for _, stra := range strategies {
				ctx := &model.TradeObjectEngineContext{Kline: kline}
				stra.Init(ctx)
			}
		}
		ce := &TradeObjectExecuteEngine{
			dataSource:        ec.dataSource,
			TradeObject:       tradeObject,
			Line:              kline,
			Strategies:        strategies,
			portfolioStrategy: ec.portfolioStrategy,
			Brokers:           ec.brokers,
		}
		ec.EventTrigger.RegisterEventReceiver(ce)
		ec.ContractEngines[tradeObject.ContractCnName+tradeObject.ContractDate] = ce
		ec.watcherBackend.AddPlotter(ce)
	}
	ec.EventTrigger.RegisterEventReceiver(local_broker2.DefaultLocalAccount)

	ec.EventTrigger.Start()
	ec.watcherBackend.Start()
	return nil
}

// 处理某个具体的合约
type TradeObjectExecuteEngine struct {
	TradeObject *model.TradeObject
	dataSource  model.DateSource
	Line        *model.KLine

	Strategies        []model.Strategy
	portfolioStrategy []model.PortfolioStrategy
	Brokers           []model.Broker
}

func (m *TradeObjectExecuteEngine) DealEvent(event *model.EventMsg) {
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

func (m *TradeObjectExecuteEngine) runStrategies(ctx *model.TradeObjectEngineContext) {
	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx)
		if stResult != nil {
			stResult.StrategyName = st.Name()
			ctx.StrategyResult = append(ctx.StrategyResult, stResult)
		}
	}
}
func (m *TradeObjectExecuteEngine) runPortfolioStrategies(ctx *model.TradeObjectEngineContext) {
	if len(ctx.StrategyResult) == 0 {
		return
	}
	for _, p := range m.portfolioStrategy {
		p(ctx)
	}
}

func (m *TradeObjectExecuteEngine) executeBuySellCmd(ctx *model.TradeObjectEngineContext) {
	for _, order := range ctx.Orders {
		for _, broker := range m.Brokers {
			broker.AddOrder(order)
		}
	}
}

func (m *TradeObjectExecuteEngine) Name() string {
	return m.TradeObject.ContractCnName + m.TradeObject.ContractDate
}

func (m *TradeObjectExecuteEngine) DoPlot(p *charts.Page) {
	position := local_broker2.GetLocalBroker().GetPositionsByContract(m.TradeObject) //????
	position.ReportToCmd()
	m.Line.DoPlot(p)
}
