package market

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"sort"
)

type KLineIndicator struct {
	*model.KLineIndicator
}

func NewKLine(t model.LineType) *KLineIndicator {
	return &KLineIndicator{
		KLineIndicator: &model.KLineIndicator{

			BaseLine: &model.BaseLine{
				Type: t,
			},
			Indicators: []model.MarketIndicator{},
		},
	}
}

func (k *KLineIndicator) GetByTsAndCount(ts int64, count int64) ([]*model.KNode, error) {
	res, err := k.BaseLine.GetByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]*model.KNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*model.KNode)
	}

	return newRes, nil
}

func (k *KLineIndicator) AddData(ts int64, node *model.KNode) {
	k.BaseLine.AddData(ts, node)
	for _, ind := range k.Indicators {
		ind.AddData(ts, node)
	}

}

func (k *KLineIndicator) AddIndicatorLine(line model.MarketIndicator) {
	k.Indicators = append(k.Indicators, line)
}

func (k *KLineIndicator) GetAllSortedData() []*model.KNode {
	oldRes := k.BaseLine.GetAllSortedData()
	res := make([]*model.KNode, len(oldRes))
	for i, v := range oldRes {
		res[i] = v.(*model.KNode)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].TimeStamp < res[j].TimeStamp
	})
	return res
}
