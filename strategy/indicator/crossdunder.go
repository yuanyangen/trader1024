package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	indicator_base2 "github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CrossUnderIndicator struct {
	*indicator_base2.IndicatorCommon
	kline             model.MarketIndicator
	lineA             *indicator_base2.Line
	lineB             *indicator_base2.Line
	crossunderScatter *indicator_base2.Scatter
}

func NewCrossUnderIndicator(kline model.MarketIndicator, lineA *indicator_base2.Line, lineB *indicator_base2.Line) *CrossUnderIndicator {
	sma := &CrossUnderIndicator{
		IndicatorCommon:   indicator_base2.NewIndicatorCommon(),
		kline:             kline,
		lineA:             lineA,
		lineB:             lineB,
		crossunderScatter: indicator_base2.NewScatter(model.LineType_Day, fmt.Sprintf("crossunder")),
	}
	kline.AddChildrenIndicator(sma)
	return sma
}

func (cou *CrossUnderIndicator) Name() string {
	return fmt.Sprintf("CrossUnder")
}

func (cou *CrossUnderIndicator) AddData(ts int64, node model.DataNode) {
	dataA := cou.getLast3Data(ts, cou.lineA)
	dataB := cou.getLast3Data(ts, cou.lineB)
	v := 0.0
	cou.crossunderScatter.AddData(ts, v)
	if len(dataB) != 3 || len(dataA) != 3 {
		cou.crossunderScatter.AddData(ts, v)
		return
	}
	if dataB[0] == 0 || dataB[1] == 0 || dataB[2] == 0 || dataA[0] == 0 || dataA[1] == 0 || dataA[2] == 0 {
		cou.crossunderScatter.AddData(ts, v)
		return
	}

	out := talib.Crossunder(dataA, dataB)
	if out {
		knode, _ := cou.kline.GetByTs(ts)
		if knode == nil {
			panic("should not reach here")
		}
		v = knode.GetValue() + 1000
	}
	cou.crossunderScatter.AddData(ts, v)
}

func (cou *CrossUnderIndicator) getLast3Data(ts int64, line *indicator_base2.Line) []float64 {
	dataA, err := line.GetLastByTsAndCount(ts, 3)
	if err != nil || len(dataA) != 3 {
		return nil
	}
	in := make([]float64, len(dataA))
	for i, v := range dataA {
		in[i] = v.GetValue()
	}
	return in
}
func (cou *CrossUnderIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (cou *CrossUnderIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if cou.crossunderScatter == nil {
		panic("crossunderScatter error")
	}
	return cou.crossunderScatter.GetByTs(ts)
}
func (cou *CrossUnderIndicator) GetLastByTsAndCount(ts int64, period int64) ([]model.DataNode, error) {
	return nil, nil
}

func (cou *CrossUnderIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := cou.crossunderScatter.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.GetTs())
		y[i] = v.GetValue()
	}
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.TitleOpts{Title: cou.Name()}, charts.YAxisOpts{Scale: true}, charts.TooltipOpts{
		Show:      true,
		Formatter: "(params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => {}",
	})

	scatter.AddXAxis(x).AddYAxis(cou.Name(), y)
	kline.Overlap(scatter)
}
