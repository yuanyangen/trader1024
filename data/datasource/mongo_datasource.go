package datasource

import (
	"context"

	"github.com/yuanyangen/trader1024/dal/mongo"
	"github.com/yuanyangen/trader1024/engine/model"
)

type MongoDataSource struct {
}

func NewMongoDataSource() *MongoDataSource {
	return &MongoDataSource{}
}

func (mds *MongoDataSource) GetDataByTs(ctx context.Context, uniqueCode string, lineType model.LineType, ts int64) *model.KLineNode {
	nd, _ := mongo.QueryDataNode(ctx, uniqueCode, ts)
	return nd
}

func (mds *MongoDataSource) SaveDataByTs(ctx context.Context, data *model.KLineNode) error {
	err := mongo.SaveDataNode(ctx, data)
	return err
}

func (mds *MongoDataSource) GetAllContractBySubjectName(ctx context.Context, subjectName string) []*model.TradeObject {
	nd := GetAllContractFromDb(ctx, subjectName)
	out := []*model.TradeObject{}
	for _, v := range nd {
		if v.LastCrawlTime > 0 {
			out = append(out, v)
		}
	}
	return out
}
