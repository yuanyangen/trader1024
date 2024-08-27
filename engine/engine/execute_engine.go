package engine

import (
	"context"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/account/local_account"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type ExecuteEngine struct {
	*baseEngine

	ContractEngines   map[string]*ContractExecuteEngine
	strategies        []func() model.Strategy
	portfolioStrategy []model.PortfolioStrategy
	brokers           []model.Broker
}

type CmdExecutor interface {
	ExecuteCmd(req *model.ContractEngineContext)
	Report()
}

func NewExecuteEngine(et model.EventTrigger, dataSource model.DateSource, strategies []func() model.Strategy, portfolioStrategy []model.PortfolioStrategy, brokers []model.Broker) *ExecuteEngine {
	e := &ExecuteEngine{
		baseEngine: &baseEngine{
			Contracts:      map[string]*model.Contract{},
			EventTrigger:   et,
			dataSource:     dataSource,
			watcherBackend: NewPlotterServers(),
		},
		strategies:        strategies,
		portfolioStrategy: portfolioStrategy,
		ContractEngines:   map[string]*ContractExecuteEngine{},
		brokers:           append([]model.Broker{local_account.GetLocalBroker()}, brokers...),
	}
	return e
}

func (ec *ExecuteEngine) Start() error {
	for _, contract := range ec.Contracts {
		strategies := make([]model.Strategy, len(ec.strategies))
		for i, stFactory := range ec.strategies {
			strategies[i] = stFactory()
		}

		kline := model.NewKLine(contract.CNName+contract.ContractDate, model.LineType_Day)
		if strategies != nil {
			for _, stra := range strategies {
				ctx := &model.ContractEngineContext{Kline: kline}
				stra.Init(ctx)
			}
		}
		ce := &ContractExecuteEngine{
			dataSource:        ec.dataSource,
			Contract:          contract,
			Line:              kline,
			Strategies:        strategies,
			portfolioStrategy: ec.portfolioStrategy,
			Brokers:           ec.brokers,
		}
		ec.EventTrigger.RegisterEventReceiver(ce)
		ec.ContractEngines[contract.ContractCnName+contract.ContractDate] = ce
		ec.watcherBackend.AddPlotter(ce)
	}
	ec.EventTrigger.RegisterEventReceiver(local_account.DefaultLocalAccount)

	ec.EventTrigger.Start()
	ec.watcherBackend.Start()
	return nil
}

// 处理某个具体的合约
type ContractExecuteEngine struct {
	Contract   *model.Contract
	dataSource model.DateSource
	Line       *model.KLine

	Strategies        []model.Strategy
	portfolioStrategy []model.PortfolioStrategy
	Brokers           []model.Broker
}

func (m *ContractExecuteEngine) DealEvent(event *model.EventMsg) {
	if event == nil || event.TimeStamp == 0 {
		return
	}
	if m.dataSource == nil {
		panic("no data_crawler source")
	}

	if event.TimeStamp >= m.Contract.ContractEndTime || event.TimeStamp < m.Contract.ContractStartTime {
		return
	}
	dataNode := m.dataSource.GetDataByTs(context.Background(), m.Contract.ContractCnName, m.Contract.ContractDate, model.LineType_Day, event.TimeStamp)
	if dataNode == nil {
		return
	}
	ctx := &model.ContractEngineContext{
		Contract:     m.Contract,
		Kline:        m.Line,
		CurrentKNode: dataNode,
		Ts:           event.TimeStamp,
	}
	m.Line.AddNodeData(event.TimeStamp, dataNode)
	m.runStrategies(ctx)
	m.runPortfolioStrategies(ctx)
	m.executeBuySellCmd(ctx)
}

func (m *ContractExecuteEngine) runStrategies(ctx *model.ContractEngineContext) {
	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx)
		if stResult != nil {
			stResult.StrategyName = st.Name()
			ctx.StrategyResult = append(ctx.StrategyResult, stResult)
		}
	}
}
func (m *ContractExecuteEngine) runPortfolioStrategies(ctx *model.ContractEngineContext) {
	if len(ctx.StrategyResult) == 0 {
		return
	}
	for _, p := range m.portfolioStrategy {
		p(ctx)
	}
}

func (m *ContractExecuteEngine) executeBuySellCmd(ctx *model.ContractEngineContext) {
	for _, order := range ctx.Orders {
		for _, broker := range m.Brokers {
			broker.AddOrder(order)
		}
	}
}

func (m *ContractExecuteEngine) Name() string {
	return m.Contract.ContractCnName + m.Contract.ContractDate
}

func (m *ContractExecuteEngine) DoPlot(p *charts.Page) {
	position := local_account.GetLocalBroker().GetCurrentLivePositions(m.Contract) //????
	position.ReportToCmd()
	DoPlot(p, m.Line)
}
