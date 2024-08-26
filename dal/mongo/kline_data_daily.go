package mongo

import (
	"context"
	"dario.cat/mergo"
	"errors"
	"fmt"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const kLineDataCollectionName = "kline_daily"

func SaveDataNode(ctx context.Context, m *model.KLineNode) error {
	if m == nil {
		return fmt.Errorf("m nil")
	}
	if m.ContractCnName == "" || m.ContractDate == "" || m.TimeStamp == 0 {
		return fmt.Errorf("cn name or date empty")
	}
	subject, err := QueryDataNode(ctx, m.ContractCnName, m.ContractDate, m.TimeStamp)
	if err != nil {
		return err
	}
	if subject == nil {
		return InsertDataNode(ctx, m)
	} else {
		mergo.Merge(m, subject, mergo.WithOverride)
		return UpdateDataNode(ctx, m.ContractCnName, m.ContractDate, m.TimeStamp, m)
	}
}

func UpdateDataNode(ctx context.Context, cnName, contractDate string, ts int64, s *model.KLineNode) error {
	collection := Trader1024Db.Collection(kLineDataCollectionName)
	_, err := collection.UpdateOne(ctx, bson.D{
		{"contractcnname", cnName},
		{"contractdate", contractDate},
		{"timestamp", ts}},
		bson.D{{"$set", s}})
	if err != nil {
		logs.Info("insert error %v %v", err, cnName)
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func InsertDataNode(ctx context.Context, m *model.KLineNode) error {
	collection := Trader1024Db.Collection(kLineDataCollectionName)
	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		logs.Info("insert error %v %v", err, m.String())
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
func QueryDataNode(ctx context.Context, cnName, contractDate string, ts int64) (*model.KLineNode, error) {
	collection := Trader1024Db.Collection(kLineDataCollectionName)
	var result *model.KLineNode
	err := collection.FindOne(ctx,
		bson.D{
			{"contractcnname", cnName},
			{"contractdate", contractDate},
			{"timestamp", ts},
		},
	).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//logs.Info("No document was found with the cnName  %s %v %v", cnName, contractDate, ts)
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func QueryAllDataNodeByContract(ctx context.Context, subjectName string, contractDate string) ([]*model.KLineNode, error) {
	collection := Trader1024Db.Collection(kLineDataCollectionName)
	var result []*model.KLineNode
	cursor, err := collection.Find(ctx,
		bson.D{
			{"contractcnname", subjectName},
			{"contractdate", contractDate},
		},
	)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logs.Info("No document was found ")
			return nil, nil
		}
		return nil, err
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
