package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type RSIIndicator struct {
	*indicator_base.IndicatorCommon
	kline   model.MarketIndicator
	SMALine *indicator_base.Line
	period  int64
}

func NewRSIIndicator(kline model.MarketIndicator, period int64) *RSIIndicator {
	sma := &RSIIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          period,
		SMALine:         indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("rsi_%v", period)),
		kline:           kline,
	}
	kline.AddChildrenIndicator(sma)
	return sma
}

func (ri *RSIIndicator) Name() string {
	return fmt.Sprintf("RSI_%v", ri.period)
}

func (ri *RSIIndicator) AddData(ts int64, node any) {
	dataI, err := ri.kline.GetLastByTsAndCount(ts, ri.period+1)
	if err != nil {
		ri.SMALine.AddData(ts, 0)
		return
	}
	data := model.NewKnodesFromAny(dataI)

	in := make([]float64, len(data))
	for i, v := range data {
		in[i] = (v.Close + v.Open) / 2
	}
	out := talib.Rsi(in, int(ri.period))
	avg := out[len(out)-1]
	ri.SMALine.AddData(ts, avg)
}
func (ri *RSIIndicator) GetAllSortedData() []any {
	return nil
}

func (ri *RSIIndicator) GetByTs(ts int64) any {
	if ri.SMALine == nil {
		panic("SMALine error")
	}
	if ri.period == 0 {
		panic("erPeriod empty")
	}
	data, err := ri.SMALine.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
}
func (ri *RSIIndicator) GetLastByTsAndCount(ts, period int64) ([]any, error) {
	return nil, nil
}

func (ri *RSIIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := ri.SMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: ri.Name()})
	line.AddXAxis(x).AddYAxis(ri.Name(), y, charts.LineOpts{ConnectNulls: false})
	ratioLine.Overlap(line)
}
