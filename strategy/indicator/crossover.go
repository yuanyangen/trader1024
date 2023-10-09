package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	indicator_base2 "github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CrossOverIndicator struct {
	*indicator_base2.IndicatorCommon
	kline            model.ContractIndicator
	lineA            *indicator_base2.Line
	lineB            *indicator_base2.Line
	crossoverScatter *indicator_base2.Scatter
}

func NewCrossOverIndicator(kline model.ContractIndicator, lineA *indicator_base2.Line, lineB *indicator_base2.Line) *CrossOverIndicator {
	sma := &CrossOverIndicator{
		IndicatorCommon: indicator_base2.NewIndicatorCommon(),

		kline:            kline,
		lineA:            lineA,
		lineB:            lineB,
		crossoverScatter: indicator_base2.NewScatter(model.LineType_Day, fmt.Sprintf("crossover")),
	}
	kline.AddChildrenIndicator(sma)
	return sma
}

func (coi *CrossOverIndicator) Name() string {
	return fmt.Sprintf("CrossOver")
}

func (coi *CrossOverIndicator) AddData(ts int64, node model.DataNode) {
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
		knode, _ := coi.kline.GetByTs(ts)
		if knode == nil {
			panic("should not reach here")
		}
		v = knode.GetValue() + 1000.0
	}
	coi.crossoverScatter.AddData(ts, v)
}

func (coi *CrossOverIndicator) getLast3Data(ts int64, line *indicator_base2.Line) []float64 {
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
func (coi *CrossOverIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (coi *CrossOverIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if coi.crossoverScatter == nil {
		panic("crossunderScatter error")
	}
	return coi.crossoverScatter.GetByTs(ts)
}
func (coi *CrossOverIndicator) GetLastByTsAndCount(ts int64, period int64) ([]model.DataNode, error) {
	return nil, nil
}

func (coi *CrossOverIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := coi.crossoverScatter.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))

	for i, v := range allData {
		y[i] = v.GetValue()
		x[i] = utils.TsToDateString(v.GetTs())
	}
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.TitleOpts{Title: coi.Name()})

	scatter.AddXAxis(x).AddYAxis(coi.Name(), y)
	kline.Overlap(scatter)
}
