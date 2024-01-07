package engine

import (
	"fmt"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type Engine struct {
	Contracts          map[string]*ContractEngine
	EventTrigger       model.EventTrigger
	strategies         []func() model.Strategy
	portfolioStrategy  []PortfolioStrategy
	cmdExecutorFactory CmdExecutorFactory // 决定
	watcherBackend     *WatcherBackend
}

type CmdExecutorFactory func() func(contract *model.Contract, kline model.ContractIndicator, portfolioStrategy []PortfolioStrategy) CmdExecutor

type CmdExecutor interface {
	ExecuteCmd(req *ContractPortfolioReq)
	Report()
}

func NewTrainEngine(et model.EventTrigger) *Engine {
	e := &Engine{
		Contracts: map[string]*ContractEngine{},
		cmdExecutorFactory: func() func(contract *model.Contract, kline model.ContractIndicator, portfolioStrategy []PortfolioStrategy) CmdExecutor {
			return newTrain
		},
		EventTrigger: et,
	}
	e.watcherBackend = NewPlotterServers(e)
	return e
}

func NewLiveExecuteEngine(et model.EventTrigger, portfolioStrategy []PortfolioStrategy) *Engine {
	e := &Engine{
		Contracts: map[string]*ContractEngine{},
		cmdExecutorFactory: func() func(contract *model.Contract, kline model.ContractIndicator, portfolioStrategy []PortfolioStrategy) CmdExecutor {
			return newLiveCmdExecutor
		},
		EventTrigger:      et,
		portfolioStrategy: portfolioStrategy,
	}
	e.watcherBackend = NewPlotterServers(e)
	return e
}

func (ec *Engine) RegisterContract(subjectCnName string, contractTimes []string, dataSource model.DateSource) {
	if len(ec.strategies) == 0 {
		panic("should register strategy_old first")
	}
	for _, contractTime := range contractTimes {

		contract := markets.GetContractByCnNam(subjectCnName, contractTime)
		if contract == nil {
			panic("contract not define")
		}
		if ec.cmdExecutorFactory == nil {
			panic("engine_mode not specify")
		}

		strategies := make([]model.Strategy, len(ec.strategies))
		for i, stFactory := range ec.strategies {
			strategies[i] = stFactory()
		}

		ce := NewContractEngine(contract, strategies, ec.cmdExecutorFactory, dataSource, ec.portfolioStrategy)
		ec.EventTrigger.RegisterEventReceiver(ce.EventTriggerChan)
		ec.Contracts[subjectCnName+contractTime] = ce
	}
}

func (ec *Engine) RegisterStrategy(stFactory func() model.Strategy) {
	ec.strategies = append(ec.strategies, stFactory)
}

func (ec *Engine) Start() error {
	if err := ec.checkEngine(); err != nil {
		panic(err)
	}

	for _, contract := range ec.Contracts {
		contract.Start()
	}
	ec.EventTrigger.Start()
	ec.watcherBackend.Start()
	return nil
}

func (ec *Engine) checkEngine() error {
	if len(ec.Contracts) == 0 {
		return fmt.Errorf("market not configed")
	}
	if len(ec.strategies) == 0 {
		panic("no strategies")
	}
	return nil
}
