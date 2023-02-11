package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/indicator/indicator_base"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CrossUnderIndicator struct {
	kline             *KLineIndicator
	lineA             *indicator_base.Line
	lineB             *indicator_base.Line
	crossunderScatter *indicator_base.Scatter
}

func NewCrossUnderIndicator(kline *KLineIndicator, lineA *indicator_base.Line, lineB *indicator_base.Line) *CrossUnderIndicator {
	sma := &CrossUnderIndicator{
		kline:             kline,
		lineA:             lineA,
		lineB:             lineB,
		crossunderScatter: indicator_base.NewScatter(kline.Type, fmt.Sprintf("crossunder")),
	}
	kline.AddIndicatorLine(sma)
	return sma
}

func (cou *CrossUnderIndicator) Name() string {
	return fmt.Sprintf("CrossUnder")
}

func (cou *CrossUnderIndicator) AddData(ts int64, node any) {
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
		knode, _ := cou.kline.GetKnodeByTs(ts)
		if knode == nil {
			panic("should not reach here")
		}
		v = knode.Close + 1000
	}
	cou.crossunderScatter.AddData(ts, v)
}

func (cou *CrossUnderIndicator) getLast3Data(ts int64, line *indicator_base.Line) []float64 {
	dataA, err := line.GetByTsAndCount(ts, 3)
	if err != nil || len(dataA) != 3 {
		return nil
	}
	in := make([]float64, len(dataA))
	for i, v := range dataA {
		in[i] = v.Value
	}
	return in
}
func (cou *CrossUnderIndicator) GetAllSortedData() []any {
	return nil
}

func (cou *CrossUnderIndicator) GetCurrentValue(ts int64) bool {
	if cou.crossunderScatter == nil {
		panic("crossunderScatter error")
	}
	r, err := cou.crossunderScatter.GetByTs(ts)
	if err != nil || r == nil {
		panic("should not reach herer")
	}
	return r.Value != 0
}

func (cou *CrossUnderIndicator) DoPlot(kline *charts.Kline) {
	allData := cou.crossunderScatter.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.TitleOpts{Title: cou.Name()})

	scatter.AddXAxis(x).AddYAxis(cou.Name(), y)
	kline.Overlap(scatter)
}
