package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/indicator/indicator_base"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type EMAIndicator struct {
	kline   *KLineIndicator
	EMALine *indicator_base.Line
	period  int64
}

func NewEMAIndicator(kline *KLineIndicator, period int64) *EMAIndicator {
	ema := &EMAIndicator{
		period:  period,
		EMALine: indicator_base.NewLine(kline.Type, fmt.Sprintf("ema_%v", period)),
		kline:   kline,
	}
	kline.AddIndicatorLine(ema)
	return ema
}

func (ema *EMAIndicator) Name() string {
	return fmt.Sprintf("EMA_%v", ema.period)
}

func (ema *EMAIndicator) AddData(ts int64, node any) {
	data, err := ema.kline.GetByTsAndCount(ts, ema.period)
	if err != nil {
		ema.EMALine.AddData(ts, 0)
		return
	}
	in := make([]float64, len(data))
	for i, v := range data {
		in[i] = (v.Close + v.Open) / 2
	}
	out := talib.Ema(in, int(ema.period))
	avg := out[len(out)-1]
	ema.EMALine.AddData(ts, avg)
}
func (ema *EMAIndicator) GetAllSortedData() []any {
	return nil
}

func (ema *EMAIndicator) GetCurrentValue(ts int64) any {
	if ema.EMALine == nil {
		panic("EMALine error")
	}
	if ema.period == 0 {
		panic("erPeriod empty")
	}
	data, err := ema.EMALine.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
}

func (ema *EMAIndicator) GetCurrentFloat(ts int64) float64 {
	v := ema.GetCurrentValue(ts)
	f, _ := v.(float64)
	return f
}

func (ema *EMAIndicator) DoPlot(kline *charts.Kline) {
	allData := ema.EMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: ema.Name()})
	line.AddXAxis(x).AddYAxis(ema.Name(), y, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}
