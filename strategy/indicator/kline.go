package indicator

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"sort"
)

type KLineIndicator struct {
	*indicator_base.BaseLine
	*indicator_base.IndicatorCommon
}

func NewKLine(name string, t model.LineType) model.MarketIndicator {
	return &KLineIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		BaseLine:        indicator_base.NewBaseLine(name, t),
	}
}

func (k *KLineIndicator) Name() string {
	return k.IndicatorCommon.Name()
}
func (k *KLineIndicator) GetByTsAndCount(ts int64, count int64) ([]any, error) {
	res, err := k.BaseLine.GetByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]any, len(res))
	for i, v := range res {
		newRes[i] = v.(*model.KNode)
	}

	return newRes, nil
}

func (k *KLineIndicator) AddData(ts int64, in any) {
	node := in.(*model.KNode)
	k.BaseLine.AddData(ts, node)
	k.IndicatorCommon.TriggerChildren(ts, node)
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
func (k *KLineIndicator) GetByTs(ts int64) any {
	vI, err := k.BaseLine.GetByTs(ts)
	if err != nil || vI == nil {
		return nil
	}
	node, ok := vI.(*model.KNode)
	if !ok {
		panic("should not reach here")
	}
	return node
}

func (k *KLineIndicator) GetAllSortedData() []any {
	oldRes := k.BaseLine.GetAllData()
	res := model.NewKnodesFromAny(oldRes)
	sort.Slice(res, func(i, j int) bool {
		return res[i].TimeStamp < res[j].TimeStamp
	})
	r := make([]any, len(res))
	for i, v := range res {
		r[i] = v
	}
	return r
}

func (k *KLineIndicator) DoPlot(p *charts.Page, kline *charts.Kline) {
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: k.IndicatorCommon.Name()},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	x, y := k.convertData()
	kline.AddXAxis(x).AddYAxis("日K", y)
	p.Add(kline)
	k.PlotChildren(p, kline)
}

func (k *KLineIndicator) convertData() ([]string, [][4]float32) {
	kDatas := k.GetAllSortedData()
	x := make([]string, len(kDatas))
	y := make([][4]float32, len(kDatas))
	for i, kni := range kDatas {
		kn := kni.(*model.KNode)
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
