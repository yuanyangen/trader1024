package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/indicator/indicator_base"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CrossOverIndicator struct {
	kline            *KLineIndicator
	lineA            *indicator_base.Line
	lineB            *indicator_base.Line
	crossoverScatter *indicator_base.Scatter
}

func NewCrossOverIndicator(kline *KLineIndicator, lineA *indicator_base.Line, lineB *indicator_base.Line) *CrossOverIndicator {
	sma := &CrossOverIndicator{
		kline:            kline,
		lineA:            lineA,
		lineB:            lineB,
		crossoverScatter: indicator_base.NewScatter(kline.Type, fmt.Sprintf("crossover")),
	}
	kline.AddIndicatorLine(sma)
	return sma
}

func (coi *CrossOverIndicator) Name() string {
	return fmt.Sprintf("CrossOver")
}

func (coi *CrossOverIndicator) AddData(ts int64, node any) {
	dataA := coi.getLast3Data(ts, coi.lineA)
	dataB := coi.getLast3Data(ts, coi.lineB)
	v := 0.0
	if len(dataB) != 3 || len(dataA) != 3 {
		coi.crossoverScatter.AddData(ts, v)
		return
	}
	if dataB[0] == 0 || dataB[1] == 0 || dataB[2] == 0 || dataA[0] == 1 || dataA[1] == 0 {
		coi.crossoverScatter.AddData(ts, v)
		return
	}

	out := talib.Crossover(dataA, dataB)
	if out {
		knode, _ := coi.kline.GetKnodeByTs(ts)
		if knode == nil {
			panic("should not reach here")
		}
		v = knode.Close + 1000.0
	}
	coi.crossoverScatter.AddData(ts, v)
}

func (coi *CrossOverIndicator) getLast3Data(ts int64, line *indicator_base.Line) []float64 {
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
func (coi *CrossOverIndicator) GetAllSortedData() []any {
	return nil
}

func (coi *CrossOverIndicator) GetCurrentValue(ts int64) bool {
	if coi.crossoverScatter == nil {
		panic("crossunderScatter error")
	}
	r, err := coi.crossoverScatter.GetByTs(ts)
	if err != nil || r == nil {
		return false
	}
	return r.Value != 0
}

func (coi *CrossOverIndicator) DoPlot(kline *charts.Kline) {
	allData := coi.crossoverScatter.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))

	for i, v := range allData {
		y[i] = v.Value
		x[i] = utils.TsToString(v.TimeStamp)
	}
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.TitleOpts{Title: coi.Name()})

	scatter.AddXAxis(x).AddYAxis(coi.Name(), y)
	kline.Overlap(scatter)
}
