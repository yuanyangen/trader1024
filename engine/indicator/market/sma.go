package market

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type SimpleMovingAverageIndicator struct {
	kline   *KLineIndicator
	smaLine *model.Line
	period  int64
}

func NewSMAIndicator(kline *model.KLineIndicator, period int64) *SimpleMovingAverageIndicator {
	sma := &SimpleMovingAverageIndicator{
		period:  period,
		smaLine: model.NewLine(kline.Type, fmt.Sprintf("sma_%v", period)),
	}
	var tmp any
	tmp = kline
	k := tmp.(KLineIndicator)
	k.AddIndicatorLine(sma)
	return sma
}

func (sma *SimpleMovingAverageIndicator) Name() string {
	return fmt.Sprintf("SimpleMovingAverage_%v", sma.period)
}

func (sma *SimpleMovingAverageIndicator) AddData(ts int64, node any) {
	data, err := sma.kline.GetByTsAndCount(ts, sma.period)
	avg := 0.0
	if err == nil {
		sum := 0.0
		for _, node := range data {
			sum += node.Close
		}
		avg = sum / float64(sma.period)
		sma.smaLine.AddData(ts, avg)
	}
}
func (sma *SimpleMovingAverageIndicator) GetAllSortedData() []any {
	return nil
}

func (sma *SimpleMovingAverageIndicator) GetCurrentValue(ts int64) float64 {
	if sma.smaLine == nil {
		panic("smaLine error")
	}
	if sma.period == 0 {
		panic("period empty")
	}
	data, err := sma.kline.GetByTsAndCount(ts, sma.period)
	avg := 0.0
	if err != nil {
		return 0
	} else {
		sum := 0.0
		for _, node := range data {
			sum += node.Close
		}
		avg = sum / float64(sma.period)
		sma.smaLine.AddData(ts, avg)
	}

	return avg
}

func (sma *SimpleMovingAverageIndicator) DoPlot(kline *charts.Kline) {
	allData := sma.smaLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "现金"})
	line.AddXAxis(x).AddYAxis(sma.Name(), y)
	kline.Overlap(line)
}
