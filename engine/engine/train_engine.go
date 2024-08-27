package engine

import (
	"context"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/account/local_account"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type TrainEngine struct {
	*baseEngine
	ContractTrainEngines map[string]*ContractTrainEngine
	strategies           []func() model.Strategy
}

func NewTrainEngine(et model.EventTrigger, dataSource model.DateSource, strategies []func() model.Strategy) *TrainEngine {
	e := &TrainEngine{
		baseEngine: &baseEngine{
			Contracts:      map[string]*model.Contract{},
			EventTrigger:   et,
			dataSource:     dataSource,
			watcherBackend: NewPlotterServers(),
		},
		strategies:           strategies,
		ContractTrainEngines: map[string]*ContractTrainEngine{},
	}
	return e
}

func (ec *TrainEngine) doRegisterContract() {
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
		ce := &ContractTrainEngine{
			Contract:   contract,
			Line:       kline,
			Strategies: strategies,
		}
		ce.dataSource = ec.dataSource
		ec.EventTrigger.RegisterEventReceiver(ce)
		ec.ContractTrainEngines[contract.ContractCnName+contract.ContractDate] = ce
		ec.watcherBackend.AddPlotter(ce)
	}
}

func (ec *TrainEngine) Start() error {
	ec.doRegisterContract()
	ec.EventTrigger.Start()
	ec.watcherBackend.Start()
	return nil
}

// 处理某个具体的合约
type ContractTrainEngine struct {
	Contract    *model.Contract
	Line        *model.KLine
	Strategies  []model.Strategy
	dataSource  model.DateSource
	CmdExecutor CmdExecutor
}

func (m *ContractTrainEngine) DealEvent(event *model.EventMsg) {
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
	m.eventHandler(dataNode)
}

func (m *ContractTrainEngine) eventHandler(data *model.KLineNode) {
	ctx := &model.ContractEngineContext{
		Contract:     m.Contract,
		Kline:        m.Line,
		CurrentKNode: data,
		Ts:           data.TimeStamp,
	}

	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx)
		if stResult == nil {
			continue
		}
		//req := &ContractEngineContext{
		//	Contract:       m.Contract,
		//	Ts:             data_crawler.TimeStamp,
		//	StrategyResult: stResult,
		//}
		//m.CmdExecutor.ExecuteCmd(req)
	}
}

func (m *ContractTrainEngine) DoPlot(p *charts.Page) {
	position := local_account.GetLocalBroker().GetCurrentLivePositions(m.Contract) //????
	position.ReportToCmd()
	DoPlot(p, m.Line)
}
func (m *ContractTrainEngine) Name() string {
	return m.Contract.ContractCnName + m.Contract.ContractDate
}

func (m *ContractTrainEngine) plotKline() *charts.Kline {
	kline := charts.NewKLine()
	return kline
}
