package indicator

import (
	"github.com/yuanyangen/trader1024/engine/data_feed"
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

func (k *KLineIndicator) GetByTsAndCount(ts int64, count int64) ([]*data_feed.KNode, error) {
	res, err := k.BaseLine.GetByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]*data_feed.KNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*data_feed.KNode)
	}

	return newRes, nil
}

func (k *KLineIndicator) AddData(ts int64, node *data_feed.KNode) {
	k.BaseLine.AddData(ts, node)
	for _, ind := range k.Indicators {
		ind.AddData(ts, node)
	}

}

func (k *KLineIndicator) AddIndicatorLine(line MarketIndicator) {
	k.Indicators = append(k.Indicators, line)
}

func (k *KLineIndicator) GetAllSortedData() []*data_feed.KNode {
	oldRes := k.BaseLine.GetAllSortedData()
	res := make([]*data_feed.KNode, len(oldRes))
	for i, v := range oldRes {
		res[i] = v.(*data_feed.KNode)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].TimeStamp < res[j].TimeStamp
	})
	return res
}
