package model

import (
	"encoding/json"
	"github.com/yuanyangen/trader1024/engine/model"
)

type GetAllDataReq struct {
	MarketId string
	DbName   string
	LineType int
}

type CommonHttpResp struct {
	Code    int
	CodeMsg string
	Data    any
}

type GetAllDataResp struct {
	Code    int
	CodeMsg string
	Data    []*model.KNode
}

func NewGetAllDataReqFromJson(in string) *GetAllDataReq {
	r := &GetAllDataReq{}
	json.Unmarshal([]byte(in), &r)
	return r
}

func NewGetAllDataRespFromJson(in string) *GetAllDataResp {
	r := &GetAllDataResp{}
	json.Unmarshal([]byte(in), &r)
	return r
}

type GetDataByTsReq struct {
	MarketId string
	DbName   string
	LineType int
	Ts       int64
}
type GetDataByTsResp struct {
	Code    int
	CodeMsg string
	Data    *model.KNode
}

func NewGetDataByTsReqFromJson(in string) *GetDataByTsReq {
	r := &GetDataByTsReq{}
	json.Unmarshal([]byte(in), &r)
	return r
}

func NewGetDataByTsRespFromJson(in string) *GetDataByTsResp {
	r := &GetDataByTsResp{}
	json.Unmarshal([]byte(in), &r)
	return r
}

type SaveDataReq struct {
	MarketId string
	DbName   string
	LineType int
	Datas    []*model.KNode
}

func NewSaveDataReqFromJson(in string) *SaveDataReq {
	r := &SaveDataReq{}
	json.Unmarshal([]byte(in), &r)
	return r
}
