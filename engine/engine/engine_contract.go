package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

// 处理某个具体的合约
type ContractEngine struct {
	EventTriggerChan chan *model.EventMsg
	Contract         *model.Contract
	Kline            model.ContractIndicator
	dataSource       model.DateSource
	//DailyIndicators  *model.DailyIndicators
	Strategies  []model.Strategy
	CmdExecutor CmdExecutor
}

func NewContractEngine(contract *model.Contract, strategies []model.Strategy, executor CmdExecutorFactory, dataSource model.DateSource, portfolioStrategy []PortfolioStrategy) *ContractEngine {
	kline := indicator.NewKLine(contract.CNName+contract.ContractTime, model.LineType_Day)
	return &ContractEngine{
		EventTriggerChan: make(chan *model.EventMsg, 1024),
		Contract:         contract,
		Kline:            kline,
		Strategies:       strategies,
		dataSource:       dataSource,
		CmdExecutor:      executor()(contract, kline, portfolioStrategy),
	}
}

func (m *ContractEngine) Start() {
	if m.EventTriggerChan == nil {
		panic("event chan error")
	}
	m.initStrategy()
	utils.AsyncRun(func() {
		for event := range m.EventTriggerChan {
			m.dealEvent(event)
		}
	})
}

func (m *ContractEngine) dealEvent(event *model.EventMsg) {
	ts := event.TimeStamp
	kData := m.dataSource.GetDataByTs(m.Contract.Id(), model.LineType_Day, ts)
	if kData != nil {
		m.Kline.AddData(ts, kData)
		m.eventHandler(kData)
	}
}

func (m *ContractEngine) DoPlot(p *charts.Page) {
	position := account.GetBackTestBroker().GetCurrentLivePositions(m.Contract.Id()) //????
	position.Report()
	kline := charts.NewKLine()
	line := charts.NewLine()

	m.Kline.DoPlot(kline, line)
	p.Add(kline)
	p.Add(line)
}

func (m *ContractEngine) eventHandler(data *model.KNode) {
	ctx := &model.ContractStrategyContext{
		Contract: m.Contract,
		Kline:    m.Kline,
	}

	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx, data.TimeStamp)
		if stResult == nil {
			continue
		}
		req := &ContractPortfolioReq{
			Contract:       m.Contract,
			Ts:             data.TimeStamp,
			StrategyResult: stResult,
		}
		m.CmdExecutor.ExecuteCmd(req)
	}
}

func (m *ContractEngine) initStrategy() {
	if m.Strategies != nil {
		for _, stra := range m.Strategies {
			ctx := &model.ContractStrategyContext{Kline: m.Kline}
			stra.Init(ctx)
		}
	}
}

func (m *ContractEngine) plotKline() *charts.Kline {
	kline := charts.NewKLine()
	return kline
}
