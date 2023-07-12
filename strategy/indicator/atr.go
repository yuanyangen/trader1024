package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type ATRIndicator struct {
	*indicator_base.IndicatorCommon
	kline   model.MarketIndicator
	SMALine *indicator_base.Line
	period  int64
}

func NewATRIndicator(kline model.MarketIndicator, period int64) *ATRIndicator {
	sma := &ATRIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          period,
		SMALine:         indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("atr_%v", period)),
		kline:           kline,
	}
	kline.AddChildrenIndicator(sma)
	return sma
}

func (sma *ATRIndicator) Name() string {
	return fmt.Sprintf("atr_%v", sma.period)
}

func (sma *ATRIndicator) AddData(ts int64, node any) {
	dataI, err := sma.kline.GetByTsAndCount(ts, sma.period)
	if err != nil {
		sma.SMALine.AddData(ts, 0)
		return
	}
	data := model.NewKnodesFromAny(dataI)

	hight := make([]float64, len(data))
	low := make([]float64, len(data))
	clo := make([]float64, len(data))

	for i, v := range data {
		hight[i] = v.High
		low[i] = v.Low
		clo[i] = v.Close
	}
	out := talib.Atr(hight, low, clo, int(sma.period)-1)
	avg := out[len(out)-1] / data[0].Close
	sma.SMALine.AddData(ts, avg)
}
func (sma *ATRIndicator) GetAllSortedData() []any {
	return nil
}

func (sma *ATRIndicator) GetByTs(ts int64) any {
	if sma.SMALine == nil {
		panic("SMALine error")
	}
	if sma.period == 0 {
		panic("erPeriod empty")
	}
	data, err := sma.SMALine.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
}
func (sma *ATRIndicator) GetByTsAndCount(ts, period int64) ([]any, error) {
	return nil, nil
}

func (sma *ATRIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := sma.SMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value * 100
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: sma.Name()})
	line.AddXAxis(x).AddYAxis(sma.Name(), y, charts.LineOpts{ConnectNulls: false})
	ratioLine.Overlap(line)
}
