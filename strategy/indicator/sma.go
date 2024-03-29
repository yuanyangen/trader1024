package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type SMAIndicator struct {
	*indicator_base.IndicatorCommon
	kline   model.MarketIndicator
	SMALine *indicator_base.Line
	period  int64
}

func NewSMAIndicator(kline model.MarketIndicator, period int64) *SMAIndicator {
	sma := &SMAIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          period,
		SMALine:         indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("sma_%v", period)),
		kline:           kline,
	}
	kline.AddChildrenIndicator(sma)
	return sma
}

func (sma *SMAIndicator) Name() string {
	return fmt.Sprintf("SimpleMovingAverage_%v", sma.period)
}

func (sma *SMAIndicator) AddData(ts int64, node any) {
	dataI, err := sma.kline.GetByTsAndCount(ts, sma.period)
	if err != nil {
		sma.SMALine.AddData(ts, 0)
		return
	}
	data := model.NewKnodesFromAny(dataI)

	in := make([]float64, len(data))
	for i, v := range data {
		in[i] = (v.Close + v.Open) / 2
	}
	out := talib.Sma(in, int(sma.period))
	avg := out[len(out)-1]
	sma.SMALine.AddData(ts, avg)
}
func (sma *SMAIndicator) GetAllSortedData() []any {
	return nil
}

func (sma *SMAIndicator) GetByTs(ts int64) any {
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
func (sma *SMAIndicator) GetByTsAndCount(ts, period int64) ([]any, error) {
	return nil, nil
}

func (sma *SMAIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := sma.SMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: sma.Name()})
	line.AddXAxis(x).AddYAxis(sma.Name(), y, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}
