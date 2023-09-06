package main

import (
	"fmt"
	model2 "github.com/yuanyangen/trader1024/data/model"
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

func getDb(dbName string) (*KVStorage, error) {
	var err error
	db := GetStorageByDbName(dbName)
	if db == nil {
		err = fmt.Errorf("db %v not support", dbName)
	}
	return db, err
}
