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

func (sma *SMAIndicator) AddData(ts int64, node model.DataNode) {
	dataI, err := sma.kline.GetLastByTsAndCount(ts, sma.period)
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
func (sma *SMAIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (sma *SMAIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if sma.SMALine == nil {
		panic("SMALine error")
	}
	if sma.period == 0 {
		panic("erPeriod empty")
	}
	return sma.SMALine.GetByTs(ts)
}
func (sma *SMAIndicator) GetLastByTsAndCount(ts, period int64) ([]model.DataNode, error) {
	return nil, nil
}

func (sma *SMAIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := sma.SMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.GetTs())
		y[i] = v.GetValue()
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: sma.Name()})
	line.AddXAxis(x).AddYAxis(sma.Name(), y, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}
