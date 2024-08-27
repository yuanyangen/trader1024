package engine

import (
	"context"
	"github.com/yuanyangen/trader1024/data/datasource"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type baseEngine struct {
	dataSource     model.DateSource
	Contracts      map[string]*model.Contract
	EventTrigger   model.EventTrigger
	watcherBackend *WatcherBackend
}

func (be *baseEngine) RegisterContract(ctx context.Context, subjectCnName string, contractDate string) {
	contract, _ := datasource.GetContractByCnName(ctx, subjectCnName, contractDate)
	if contract == nil {
		panic("contract not define")
	}
	if be.Contracts == nil {
		be.Contracts = map[string]*model.Contract{}
	}
	be.Contracts[subjectCnName+contractDate] = contract
}

func (be *baseEngine) RegisterContractBySubjectName(ctx context.Context, subjectCnName string) {
	contracts := be.dataSource.GetAllContractBySubjectName(ctx, subjectCnName)
	if be.Contracts == nil {
		be.Contracts = map[string]*model.Contract{}
	}
	for _, contract := range contracts {
		be.Contracts[subjectCnName+contract.ContractDate] = contract
	}
}
