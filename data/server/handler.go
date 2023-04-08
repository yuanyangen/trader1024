package main

import (
	"fmt"
	model2 "github.com/yuanyangen/trader1024/data/model"
	storage2 "github.com/yuanyangen/trader1024/data/storage"
	"github.com/yuanyangen/trader1024/engine/model"
)

func GetAllData(body string) (any, error) {
	req := model2.NewGetAllDataReqFromJson(body)

	storage, err := getDb(req.DbName)
	if err != nil {
		return nil, err
	}
	v := storage.GetAllData(req.MarketId, model.LineType(req.LineType))
	return v, nil
}
func GetDataByTs(body string) (any, error) {
	req := model2.NewGetDataByTsReqFromJson(body)

	storage, err := getDb(req.DbName)
	if err != nil {
		return nil, err
	}
	v := storage.GetDataByTs(req.MarketId, model.LineType(req.LineType), req.Ts)
	return v, nil
}

func SaveData(body string) (any, error) {
	req := model2.NewSaveDataReqFromJson(body)
	storage, err := getDb(req.DbName)
	if err != nil {
		return nil, err
	}

	if len(req.Datas) == 0 {
		return nil, fmt.Errorf("req data empty")
	}
	storage.SaveData(req.MarketId, model.LineType(req.LineType), req.Datas)
	return nil, nil
}

func getDb(dbName string) (*storage2.KVStorage, error) {
	var storage *storage2.KVStorage
	var err error
	if dbName == "main" {
		storage = storage2.MainStorage()
	} else if dbName == "eastmoney" {
		storage = storage2.EastMoneyStorage()
	} else if dbName == "test" {
		storage = storage2.TestStorage()
	} else {
		err = fmt.Errorf("db %v not support", dbName)
	}
	return storage, err
}
