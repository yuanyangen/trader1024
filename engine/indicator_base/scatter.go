package indicator_base

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"sort"
)

type Scatter struct {
	*BaseLine
}

type ScatterNode struct {
	Value     float64
	TimeStamp int64
}

func NewScatter(t model.LineType, name string) *Scatter {
	return &Scatter{
		BaseLine: NewBaseLine(name, t),
	}
}

func (k *Scatter) GetByTs(ts int64) (*ScatterNode, error) {
	res, err := k.BaseLine.GetByTs(ts)
	if err != nil {
		return nil, err
	}
	vv := res.(*ScatterNode)
	return vv, nil
}

func (k *Scatter) GetLastByTsAndCount(ts int64, count int64) ([]*ScatterNode, error) {
	res, err := k.BaseLine.GetLastByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]*ScatterNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*ScatterNode)
	}

	return newRes, nil
}

func (k *Scatter) AddData(ts int64, value float64) {
	node := &ScatterNode{
		TimeStamp: ts, Value: value,
	}
	k.BaseLine.AddData(ts, node)
}

func (k *Scatter) GetAllSortedData() []*ScatterNode {
	res := k.BaseLine.GetAllData()
	newRes := make([]*ScatterNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*ScatterNode)
	}

	sort.Slice(newRes, func(i, j int) bool {
		return newRes[i].TimeStamp < newRes[j].TimeStamp
	})
	return newRes
}
