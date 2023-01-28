package indicator

import (
	"sort"
)

type KLineIndicator struct {
	*BaseLine
	Indicators []MarketIndicator
}

func NewKLine(t LineType) *KLineIndicator {
	return &KLineIndicator{
		BaseLine:   NewBaseLine(t),
		Indicators: []MarketIndicator{},
	}
}

func (k *KLineIndicator) GetByTsAndCount(ts int64, count int64) ([]*KNode, error) {
	res, err := k.BaseLine.GetByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]*KNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*KNode)
	}

	return newRes, nil
}

func (k *KLineIndicator) AddData(ts int64, node *KNode) {
	k.BaseLine.AddData(ts, node)
	for _, ind := range k.Indicators {
		ind.AddData(ts, node)
	}

}

func (k *KLineIndicator) AddIndicatorLine(line MarketIndicator) {
	k.Indicators = append(k.Indicators, line)
}

func (k *KLineIndicator) GetAllSortedData() []*KNode {
	oldRes := k.BaseLine.GetAllSortedData()
	res := make([]*KNode, len(oldRes))
	for i, v := range oldRes {
		res[i] = v.(*KNode)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].TimeStamp < res[j].TimeStamp
	})
	return res
}
