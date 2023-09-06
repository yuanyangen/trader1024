package model

import (
	"encoding/json"
	"fmt"
)

type LineType int64

const LineType_Day LineType = 1
const LineType_Minite LineType = 2
const LineType_5Minite LineType = 3
const LineType_Hour LineType = 4
const LineType_Week LineType = 5

type KNode struct {
	Date          string
	High          float64
	Low           float64
	Open          float64
	Close         float64
	Volume        float64
	Turnover      float64 // 成交额
	Swing         float64 //振幅
	Increase      float64 // 涨跌幅 ??
	IncreaseMount float64 // 涨跌额 ??
	TurnoverRate  float64 //换手率
	TimeStamp     int64
}

// 获取knode的当前价格
func (k *KNode) GetValue() float64 {
	return (k.Open + k.Close) / 2
}

func NewKnodeFromAny(val any) *KNode {
	if val == nil {
		return nil
	}
	r, _ := val.(*KNode)
	return r
}

func NewKnodesFromAny(val []any) []*KNode {
	if val == nil {
		return nil
	}
	res := make([]*KNode, len(val))
	for i, v := range val {
		r, _ := v.(*KNode)
		if r == nil {
			panic("knod nil")
		}
		res[i] = r

	}
	return res
}

func NewKnodeFromJson(val []byte) *KNode {
	knode := &KNode{}
	err := json.Unmarshal(val, &knode)
	if err != nil {
		panic(fmt.Sprintf("data in db unmarshal error %v", err))
	}
	return knode
}
