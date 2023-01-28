package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type SimpleMovingAverageIndicator struct {
	kline   *KLineIndicator
	smaLine *Line
	period  int64
}

func NewSMAIndicator(kline *KLineIndicator, period int64) *SimpleMovingAverageIndicator {
	sma := &SimpleMovingAverageIndicator{
		period:  period,
		smaLine: NewLine(kline.Type, fmt.Sprintf("sma_%v", period)),
		kline:   kline,
	}
	kline.AddIndicatorLine(sma)
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
	}
	sma.smaLine.AddData(ts, avg)
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
	data, err := sma.smaLine.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
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
	line.SetGlobalOptions(charts.TitleOpts{Title: sma.Name()})
	line.AddXAxis(x).AddYAxis(sma.Name(), y, charts.LineOpts{Smooth: true, ConnectNulls: true})
	kline.Overlap(line)
}
