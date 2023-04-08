package storage

import (
	"bytes"
	"encoding/json"
	model2 "github.com/yuanyangen/trader1024/data/model"
	"github.com/yuanyangen/trader1024/engine/model"
	"io"
	"net/http"
)

// const httpAddr = "http://127.0.0.1:8888"
const httpAddr = "http://192.168.1.106:8888"

//
//第一层是：数据的来源，包括： main, eastMoney等，
//第二层是具体的DB文件，分成128个文件，文件名字是通过marketID hash得到的。
// bucket的名字是 market + dataType ，数据类型是枚举值： daily_k, minute5_k
// 每个bucket的k 是对应数据unix 时间戳， value是数据的json编码的结果。

//表示 存放 某Market的某个类型的在某个时间范围范围内的量
// 每一个文件de key的

type HttpStorage struct {
	Name string
}

var allHppStorage = map[string]*HttpStorage{}

func init() {
	InitHttpStorage("eastmoney")
	InitHttpStorage("main")
}

func EastMoneyHttpStorage() *HttpStorage {
	return allHppStorage["eastmoney"]
}

func MainHttpStorage() *HttpStorage {
	return allHppStorage["main"]
}

func InitHttpStorage(name string) *HttpStorage {
	cs := &HttpStorage{
		Name: name,
	}
	allHppStorage[name] = cs
	return cs
}

func (cs *HttpStorage) SaveData(marketId string, t model.LineType, kdatas []*model.KNode) error {
	param := model2.SaveDataReq{
		MarketId: marketId,
		LineType: int(t),
		DbName:   cs.Name,

		Datas: kdatas,
	}
	_, err := cs.httpPost("/save_data", param)
	return err
}

func (cs *HttpStorage) GetAllData(marketId string, t model.LineType) []*model.KNode {
	param := model2.GetDataByTsReq{
		MarketId: marketId,
		LineType: int(t),
		DbName:   cs.Name,
	}
	r, err := cs.httpPost("/get_all_data", param)
	if err != nil {
		return nil
	}
	resp := model2.NewGetAllDataRespFromJson(r)
	return resp.Data
}

func (cs *HttpStorage) GetDataByTs(marketId string, t model.LineType, ts int64) *model.KNode {
	param := model2.GetDataByTsReq{
		MarketId: marketId,
		LineType: int(t),
		Ts:       ts,
		DbName:   cs.Name,
	}

	r, err := cs.httpPost("/get_data_by_ts", param)
	if err != nil {
		return nil
	}
	resp := model2.NewGetDataByTsRespFromJson(r)
	return resp.Data
}

func (cs *HttpStorage) httpPost(uri string, param any) (string, error) {
	url := httpAddr + uri
	body, _ := json.Marshal(param)
	resp, err := http.Post(url, "", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respB), nil
}
