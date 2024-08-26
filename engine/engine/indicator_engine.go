package engine

import (
	"context"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type IndicatorEngine struct {
	*baseEngine
	indicators         []model.Indicator
	contractIndicators map[string]*ContractIndicator
}

func NewIndicatorEngine(et model.EventTrigger, dataSource model.DateSource, indicators []model.Indicator) *IndicatorEngine {
	e := &IndicatorEngine{
		baseEngine: &baseEngine{
			Contracts:    map[string]*model.Contract{},
			EventTrigger: et,
			dataSource:   dataSource,
		},
		indicators:         indicators,
		contractIndicators: map[string]*ContractIndicator{},
	}
	return e
}

func (ec *IndicatorEngine) Start() error {
	for _, contract := range ec.Contracts {
		ce := &ContractIndicator{
			Contract:   contract,
			Line:       model.NewKLine(contract.CNName+contract.ContractDate, model.LineType_Day),
			Strategies: ec.indicators,
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
	Contract   *model.Contract
	Line       *model.KLine
	Strategies []model.Indicator
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

	if ts >= m.Contract.ContractEndTime || ts < m.Contract.ContractStartTime {
		return
	}
	dataNode := m.dataSource.GetDataByTs(ctx, m.Contract.ContractCnName, m.Contract.ContractDate, model.LineType_Day, ts)
	if dataNode == nil {
		return
	}

	m.Line.AddNodeData(ts, dataNode)
	for _, st := range m.Strategies {
		st.FillIndicatorToLine(m.Line, dataNode.TimeStamp)
	}
	err := m.dataSource.SaveDataByTs(ctx, dataNode)
	if err != nil {
		logs.Info("save node to ts failed %v", err)
	}
}
