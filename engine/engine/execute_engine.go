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

	ContractEngines    map[string]*ContractExecuteEngine
	strategies         []func() model.Strategy
	portfolioStrategy  []model.PortfolioStrategy
	cmdExecutorFactory CmdExecutor // 决定
	watcherBackend     *WatcherBackend
}

type CmdExecutor interface {
	ExecuteCmd(req *model.ContractPortfolioReq)
	Report()
}

func NewExecuteEngine(et model.EventTrigger, dataSource model.DateSource, strategies []func() model.Strategy, portfolioStrategy []model.PortfolioStrategy) *ExecuteEngine {
	e := &ExecuteEngine{
		baseEngine: &baseEngine{
			Contracts:    map[string]*model.Contract{},
			EventTrigger: et,
			dataSource:   dataSource,
		},
		strategies:         strategies,
		portfolioStrategy:  portfolioStrategy,
		ContractEngines:    map[string]*ContractExecuteEngine{},
		cmdExecutorFactory: nil,
	}
	e.watcherBackend = NewPlotterServers()
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
				ctx := &model.ContractStrategyContext{Kline: kline}
				stra.Init(ctx)
			}
		}
		ce := &ContractExecuteEngine{
			dataSource:        ec.dataSource,
			Contract:          contract,
			Line:              kline,
			Strategies:        strategies,
			portfolioStrategy: ec.portfolioStrategy,
		}
		ec.EventTrigger.RegisterEventReceiver(ce)
		ec.ContractEngines[contract.ContractCnName+contract.ContractDate] = ce
		ec.watcherBackend.AddPlotter(ce)
	}

	ec.EventTrigger.Start()
	ec.watcherBackend.Start()
	return nil
}

// 处理某个具体的合约
type ContractExecuteEngine struct {
	Contract          *model.Contract
	Line              *model.KLine
	Strategies        []model.Strategy
	dataSource        model.DateSource
	portfolioStrategy []model.PortfolioStrategy
}

func (m *ContractExecuteEngine) DealEvent(event *model.EventMsg) {
	if event == nil {
		return
	}
	if m.dataSource == nil {
		panic("no data_crawler source")
	}
	ts := event.TimeStamp

	if ts >= m.Contract.ContractEndTime || ts < m.Contract.ContractStartTime {
		return
	}
	dataNode := m.dataSource.GetDataByTs(context.Background(), m.Contract.ContractCnName, m.Contract.ContractDate, model.LineType_Day, ts)
	if dataNode == nil {
		return
	}

	m.Line.AddNodeData(ts, dataNode)
	ctx := &model.ContractStrategyContext{
		Contract: m.Contract,
		Kline:    m.Line,
	}

	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx, dataNode.TimeStamp)
		if stResult == nil {
			continue
		}
		req := &model.ContractPortfolioReq{
			Contract:       m.Contract,
			Ts:             dataNode.TimeStamp,
			StrategyResult: stResult,
		}
		broker := local_account.GetBackTestBroker()
		for _, p := range m.portfolioStrategy {
			p(broker, req)
		}
	}
}
func (m *ContractExecuteEngine) Name() string {
	return m.Contract.ContractCnName + m.Contract.ContractDate
}

func (m *ContractExecuteEngine) DoPlot(p *charts.Page) {
	position := local_account.GetBackTestBroker().GetCurrentLivePositions(m.Contract.Id()) //????
	position.Report()
	DoPlot(p, m.Line)
}

func (m *ContractExecuteEngine) plotKline() *charts.Kline {
	kline := charts.NewKLine()
	return kline
}
