package datasource

import (
	"context"
	"github.com/yuanyangen/trader1024/dal/mongo"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
)

type VendorMarket interface {
	GetMarketFromVendor(CNName string, date string) *model.Contract
}

func GetSubjectByCnNam(ctx context.Context, name string) *model.SubjectDO {
	v, _ := mongo.QuerySubject(ctx, name)
	return v
}

func GetContractByCnName(ctx context.Context, name string, contractDate string) (*model.Contract, error) {
	s, err := mongo.QuerySubject(ctx, name)
	if err != nil {
		return nil, err
	}
	if s == nil {
		panic(name + " not support")
	}
	c, err := mongo.QueryContract(ctx, name, contractDate)
	if err != nil {
		return nil, err
	}
	if c == nil {
		panic(name + " not support")
	}

	return &model.Contract{
		SubjectDO:  s,
		ContractDO: c,
	}, nil
}

func GetAllFutureSubjects(ctx context.Context) ([]*model.SubjectDO, error) {
	return mongo.QueryAllSubject(ctx)
}

func GetAllContractFromDb(ctx context.Context, subjectName string) []*model.Contract {
	subject, err := mongo.QuerySubject(ctx, subjectName)
	if err != nil {
		logs.Info("QuerySubject error " + err.Error())
		return nil
	}
	allContracts, err := mongo.QueryAllContractBySubjectName(ctx, subjectName)
	if err != nil {
		logs.Info("QueryAllContractBySubjectName error " + err.Error())
		return nil
	}
	out := make([]*model.Contract, len(allContracts))
	for i, v := range allContracts {
		out[i] = &model.Contract{
			SubjectDO:  subject,
			ContractDO: v,
		}
	}
	return out
}
