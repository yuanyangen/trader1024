package engine

import (
	"context"

	"github.com/yuanyangen/trader1024/data/datasource"
	"github.com/yuanyangen/trader1024/engine/model"
)

// 最外层，处理全部合约
type baseEngine struct {
	dataSource     model.DateSource
	TradeObjects   map[string]*model.TradeObject
	EventTrigger   model.EventTrigger
	watcherBackend *WatcherBackend
}

func (be *baseEngine) RegisterContract(ctx context.Context, subjectCnName string, contractDate string) {
	contract, _ := datasource.GetContractByCnName(ctx, subjectCnName, contractDate)
	if contract == nil {
		panic("contract not define")
	}
	if be.TradeObjects == nil {
		be.TradeObjects = map[string]*model.TradeObject{}
	}
	be.TradeObjects[subjectCnName+contractDate] = contract
}

func (be *baseEngine) RegisterContractBySubjectName(ctx context.Context, subjectCnName string) {
	contracts := be.dataSource.GetAllContractBySubjectName(ctx, subjectCnName)
	if be.TradeObjects == nil {
		be.TradeObjects = map[string]*model.TradeObject{}
	}
	for _, contract := range contracts {
		be.TradeObjects[subjectCnName+contract.ContractDate] = contract
	}
}
