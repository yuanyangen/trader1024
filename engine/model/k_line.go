package model

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"sort"
)

type KLine struct {
	*BaseLine
}

type KLineNode struct {
	ContractCnName string
	ContractDate   string
	Type           LineType
	TimeStamp      int64
	TimeStampDesc  string
	High           float64
	Low            float64
	Open           float64
	Close          float64
	Volume         float64
	Turnover       float64           // 成交额
	Swing          float64           // 振幅
	Increase       float64           // 涨跌幅 ??
	IncreaseMount  float64           // 涨跌额 ??
	TurnoverRate   float64           //换手率
	SmaData        map[int64]float64 // sma, key 是sma的周期， value是对应的值
}

func NewDataNodeNewFromAny(val any) *KLineNode {
	if val == nil {
		return nil
	}
	r, _ := val.(*KLineNode)
	return r
}

func NewDataNodeFromJson(val []byte) *KLineNode {
	knode := &KLineNode{}
	err := json.Unmarshal(val, &knode)
	if err != nil {
		panic(fmt.Sprintf("data_crawler in db unmarshal error %v", err))
	}
	return knode
}

func (dn *KLineNode) GetTs() int64 {
	return dn.TimeStamp
}
func (dn *KLineNode) String() string {
	r, _ := sonic.MarshalString(dn)
	return r
}

func NewKLine(name string, t LineType) *KLine {
	bl := &KLine{
		BaseLine: NewBaseLine(name, t),
	}
	return bl
}

func (bl *KLine) GetNodeByTs(ts int64) (*KLineNode, error) {
	return bl.convertAnyToLineNode(bl.GetLastByTs(ts))
}

// get last one
func (bl *KLine) GetLastNodeByTs(ts int64) (*KLineNode, error) {
	return bl.convertAnyToLineNode(bl.GetLastByTs(ts))
}

func (bl *KLine) GetLastNodeByTsAndCount(ts int64, count int64) ([]*KLineNode, error) {
	return bl.convertAnyToLineNodes(bl.GetLastByTsAndCount(ts, count))
}

func (bl *KLine) GetForwardNodeByTsAndCount(ts int64, count int64) ([]*KLineNode, error) {
	return bl.convertAnyToLineNodes(bl.GetForwardByTsAndCount(ts, count))

}

func (bl *KLine) AddNodeData(ts int64, node *KLineNode) {
	bl.AddData(ts, &LineNode{DataNode: node, TimeStamp: ts})
}

func (bl *KLine) GetAllNodeData() []*KLineNode {
	allData := bl.GetAllData()
	out, _ := bl.convertAnyToLineNodes(allData, nil)
	return out
}

func (bl *KLine) GetAllSortedData() []*KLineNode {
	ldRes := bl.GetAllNodeData()
	sort.Slice(ldRes, func(i, j int) bool {
		return ldRes[i].TimeStamp < ldRes[j].TimeStamp
	})
	return ldRes
}

func (bl *KLine) convertAnyToLineNode(in *LineNode, err error) (*KLineNode, error) {
	if err != nil {
		return nil, err
	}
	if in != nil && in.DataNode != nil {
		node, ok := in.DataNode.(*KLineNode)
		if !ok {
			return nil, fmt.Errorf("data not lineNode")
		}
		return node, nil
	}
	return nil, fmt.Errorf("no line node")
}

func (bl *KLine) convertAnyToLineNodes(in []*LineNode, err error) ([]*KLineNode, error) {
	if err != nil {
		return nil, err
	}
	out := []*KLineNode{}
	for _, vv := range in {
		if vv != nil && vv.DataNode != nil {
			node, ok := vv.DataNode.(*KLineNode)
			if !ok {
				return nil, fmt.Errorf("data not lineNode")
			}
			out = append(out, node)
		}
	}
	return out, nil
}
