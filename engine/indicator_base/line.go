package indicator_base

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"sort"
)

type Line struct {
	*BaseLine
}

type LineNode struct {
	Value     float64
	TimeStamp int64
}

func NewLine(t model.LineType, name string) *Line {
	return &Line{
		BaseLine: NewBaseLine(name, t),
	}
}

func (k *Line) GetByTs(ts int64) (*LineNode, error) {
	res, err := k.BaseLine.GetByTs(ts)
	if err != nil {
		return nil, err
	}
	vv := res.(*LineNode)
	return vv, nil
}

func (k *Line) GetLastByTsAndCount(ts int64, count int64) ([]*LineNode, error) {
	res, err := k.BaseLine.GetLastByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]*LineNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*LineNode)
	}

	return newRes, nil
}

func (k *Line) AddData(ts int64, value float64) {
	node := &LineNode{
		TimeStamp: ts, Value: value,
	}
	k.BaseLine.AddData(ts, node)
}

func (k *Line) GetAllSortedData() []*LineNode {
	res := k.BaseLine.GetAllData()
	newRes := make([]*LineNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*LineNode)
	}

	sort.Slice(newRes, func(i, j int) bool {
		return newRes[i].TimeStamp < newRes[j].TimeStamp
	})
	return newRes
}
