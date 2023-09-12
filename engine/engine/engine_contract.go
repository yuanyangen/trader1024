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
	Kline            model.MarketIndicator
	dataSource       model.DateSource
	//DailyIndicators  *model.DailyIndicators
	Strategies  []model.Strategy
	CmdExecutor CmdExecutor
}

func NewContractEngine(contract *model.Contract, strategies []model.Strategy, executor CmdExecutorFactory, dataSource model.DateSource) *ContractEngine {
	kline := indicator.NewKLine(contract.CNName+contract.ContractTime, model.LineType_Day)
	return &ContractEngine{
		EventTriggerChan: make(chan *model.EventMsg, 1024),
		Contract:         contract,
		Kline:            kline,
		Strategies:       strategies,
		dataSource:       dataSource,
		CmdExecutor:      executor()(contract, kline),
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
	kData := m.dataSource.GetDataByTs(m.Contract.ContractId, model.LineType_Day, ts)
	if kData != nil {
		data := &model.Data{
			DataType: model.DataTypeKLine,
			KData:    kData,
		}
		m.Kline.AddData(ts, data)
		m.eventHandler(data)
	}
}

func (m *ContractEngine) DoPlot(p *charts.Page) {
	position := account.GetBackTestBroker().GetCurrentLivePositions(m.Contract.ContractId) //????
	position.Report()
	kline := charts.NewKLine()
	line := charts.NewLine()

	m.Kline.DoPlot(kline, line)
	p.Add(kline)
	p.Add(line)
}

func (m *ContractEngine) eventHandler(data *model.Data) {
	ctx := &model.ContractStrategyContext{
		Contract: m.Contract,
	}
	req := &model.MarketPortfolioReq{
		Contract: m.Contract,
		Ts:       data.KData.TimeStamp,
	}

	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx, data.KData.TimeStamp)
		if stResult != nil {
			req.Strategies = append(req.Strategies, &model.StrategyReq{
				StrategyName: st.Name(),
				Cmd:          stResult,
				Reason:       st.Name(),
			})
		}
	}
	m.CmdExecutor.ExecuteCmd(req)
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
