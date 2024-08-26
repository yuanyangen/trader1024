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

const collectionContractName = "contract_infos"

func SaveContractInfo(ctx context.Context, m *model.ContractDO) error {
	if m == nil {
		return fmt.Errorf("m nil")
	}
	if m.ContractCnName == "" || m.ContractDate == "" {
		return fmt.Errorf("cn name or date empty")
	}
	subject, err := QueryContract(ctx, m.ContractCnName, m.ContractDate)
	if err != nil {
		return err
	}
	if subject == nil {
		return InsertContractInfo(ctx, m)
	} else {
		mergo.Merge(m, subject, mergo.WithOverride)
		return UpdateContract(ctx, m.ContractCnName, m.ContractDate, m)
	}
}

func UpdateContract(ctx context.Context, cnName, contractDate string, s *model.ContractDO) error {
	collection := Trader1024Db.Collection(collectionContractName)
	_, err := collection.UpdateOne(ctx, bson.D{{"contractcnname", cnName}, {"contractdate", contractDate}}, bson.D{{"$set", s}})
	if err != nil {
		logs.Info("insert error %v %v", err, cnName)
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func InsertContractInfo(ctx context.Context, m *model.ContractDO) error {
	collection := Trader1024Db.Collection(collectionContractName)
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
func QueryContract(ctx context.Context, cnName, contractDate string) (*model.ContractDO, error) {
	collection := Trader1024Db.Collection(collectionContractName)
	var result *model.ContractDO
	err := collection.FindOne(ctx, bson.D{
		{"contractcnname", cnName},
		{"contractdate", contractDate},
	}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logs.Info(" No document was found with the cnName  %s %v", cnName, contractDate)
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func QueryAllContractBySubjectName(ctx context.Context, subjectName string) ([]*model.ContractDO, error) {
	collection := Trader1024Db.Collection(collectionContractName)
	var result []*model.ContractDO
	cursor, err := collection.Find(ctx, bson.D{{"contractcnname", subjectName}})
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
