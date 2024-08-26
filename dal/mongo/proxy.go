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

const collectionNameProxies = "proxies"

func SaveProxy(ctx context.Context, m *model.HttpProxy) error {
	if m == nil {
		return fmt.Errorf("m nil")
	}
	if m.Addr == "" {
		return fmt.Errorf("addr  empty")
	}
	subject, err := QueryProxy(ctx, m.Addr)
	if err != nil {
		return err
	}
	if subject == nil {
		return InsertProxy(ctx, m)
	} else {
		mergo.Merge(m, subject, mergo.WithOverride)

		return UpdateProxy(ctx, m.Addr, m)
	}
}

func UpdateProxy(ctx context.Context, cnName string, s *model.HttpProxy) error {
	collection := Trader1024Db.Collection(collectionNameProxies)
	_, err := collection.UpdateOne(ctx, bson.D{{"addr", cnName}}, bson.D{{"$set", s}})
	if err != nil {
		logs.Info("insert error %v %v", err, cnName)
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func InsertProxy(ctx context.Context, m *model.HttpProxy) error {
	collection := Trader1024Db.Collection(collectionNameProxies)
	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		logs.Info("insert error %v %v\n", err, m.String())
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
func QueryProxy(ctx context.Context, cnName string) (*model.HttpProxy, error) {
	collection := Trader1024Db.Collection(collectionNameProxies)
	var result *model.HttpProxy
	err := collection.FindOne(ctx, bson.D{{"addr", cnName}}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logs.Info("No document was found with the addr  %s\n", cnName)
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
func QueryProxyWithExpireTime(ctx context.Context, expireTS int64) (*model.HttpProxy, error) {
	collection := Trader1024Db.Collection(collectionNameProxies)
	var result *model.HttpProxy
	err := collection.FindOne(ctx, bson.D{{"expiretime", bson.D{{"$gt", expireTS}}}}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//logs.Info("No proxy was found with the expiretime  %s", expireTS)
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
