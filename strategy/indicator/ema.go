package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type EMAIndicator struct {
	*indicator_base.IndicatorCommon
	kline   *KLineIndicator
	EMALine *indicator_base.Line
	period  int64
}

func NewEMAIndicator(kline *KLineIndicator, period int64) *EMAIndicator {
	ema := &EMAIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          period,
		EMALine:         indicator_base.NewLine(kline.Type, fmt.Sprintf("ema_%v", period)),
		kline:           kline,
	}
	kline.AddChildrenIndicator(ema)
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
		knode := model.NewKnodeFromAny(v)
		if knode != nil {
			in[i] = (knode.Close + knode.Open) / 2
		} else {
			panic("knode nil")
		}
	}
	out := talib.Ema(in, int(ema.period))
	avg := out[len(out)-1]
	ema.EMALine.AddData(ts, avg)
}
func (ema *EMAIndicator) GetAllSortedData() []any {
	return nil
}
func (ema *EMAIndicator) GetByTsAndCount(ts int64, period int64) ([]any, error) {
	return nil, nil
}

func (ema *EMAIndicator) GetByTs(ts int64) any {
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
	v := ema.GetByTs(ts)
	f, _ := v.(float64)
	return f
}

func (ema *EMAIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
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
