package indicator

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"sort"
)

type KLineIndicator struct {
	*indicator_base.BaseLine
	Indicators []MarketIndicator
}

func NewKLine(name string, t model.LineType) *KLineIndicator {
	return &KLineIndicator{
		BaseLine:   indicator_base.NewBaseLine(name, t),
		Indicators: []MarketIndicator{},
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

func (k *KLineIndicator) GetKnodeByTs(ts int64) (*model.KNode, error) {
	vI, err := k.BaseLine.GetByTs(ts)
	if err != nil || vI == nil {
		return nil, err
	}
	node, ok := vI.(*model.KNode)
	if !ok {
		panic("should not reach here")
	}
	return node, nil
}

func (k *KLineIndicator) AddIndicatorLine(line MarketIndicator) {
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

func (k *KLineIndicator) plotKline() *charts.Kline {
	kline := charts.NewKLine()
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: k.Name},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	x, y := k.convertData()
	kline.AddXAxis(x).AddYAxis("日K", y)
	return kline
}

func (k *KLineIndicator) convertData() ([]string, [][4]float32) {
	kDatas := k.GetAllSortedData()
	x := make([]string, len(kDatas))
	y := make([][4]float32, len(kDatas))
	for i, kn := range kDatas {
		x[i] = kn.Date
		y[i] = [4]float32{
			float32(kn.Open),
			float32(kn.Close),
			float32(kn.Low),
			float32(kn.High),
		}
	}
	return x, y
}
