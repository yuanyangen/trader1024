package engine

import (
	"context"

	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type IndicatorEngine struct {
	*baseEngine
	indicatorsChain    [][]model.Indicator
	contractIndicators map[string]*ContractIndicator
}

func NewIndicatorEngine(et model.EventTrigger, dataSource model.DateSource, indicatorsChain ...[]model.Indicator) *IndicatorEngine {
	e := &IndicatorEngine{
		baseEngine: &baseEngine{
			TradeObjects: map[string]*model.TradeObject{},
			EventTrigger: et,
			dataSource:   dataSource,
		},
		indicatorsChain:    indicatorsChain,
		contractIndicators: map[string]*ContractIndicator{},
	}
	return e
}

func (ec *IndicatorEngine) Start() error {
	for _, contract := range ec.TradeObjects {
		ce := &ContractIndicator{
			Contract:   contract,
			Line:       model.NewKLine(contract.CNName+contract.ContractDate, model.LineType_Day),
			Indicators: ec.indicatorsChain,
			dataSource: ec.dataSource,
		}
		ec.EventTrigger.RegisterEventReceiver(ce)
		ec.contractIndicators[contract.ContractCnName+contract.ContractDate] = ce
	}
	ec.EventTrigger.Start()
	return nil
}

// 处理某个具体的合约
type ContractIndicator struct {
	Contract   *model.TradeObject
	Line       *model.KLine
	Indicators [][]model.Indicator
	dataSource model.DateSource
}

func (m *ContractIndicator) DealEvent(event *model.EventMsg) {
	ctx := context.Background()
	if event == nil {
		return
	}
	if m.dataSource == nil {
		panic("no data_crawler source")
	}
	ts := event.TimeStamp

	if (m.Contract.ContractEndTime != 0 && ts >= m.Contract.ContractEndTime) || (ts < m.Contract.ContractStartTime && m.Contract.ContractStartTime != 0) {
		return
	}
	dataNode := m.dataSource.GetDataByTs(ctx, m.Contract.UniqueCode, model.LineType_Day, ts)
	if dataNode == nil {
		return
	}

	m.Line.AddNodeData(ts, dataNode)
	for _, indicators := range m.Indicators {
		for _, indi := range indicators {
			indi.FillIndicatorToLine(m.Line, dataNode.TimeStamp, dataNode)
		}
	}

	err := m.dataSource.SaveDataByTs(ctx, dataNode)
	if err != nil {
		logs.Info("save node to ts failed %v", err)
	}
}
