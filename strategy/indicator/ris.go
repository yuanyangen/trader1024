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
	kline   model.ContractIndicator
	SMALine *indicator_base.Line
	period  int64
}

func NewRSIIndicator(kline model.ContractIndicator, period int64) *RSIIndicator {
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

func (ri *RSIIndicator) AddData(ts int64, node model.DataNode) {
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
func (ri *RSIIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (ri *RSIIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if ri.SMALine == nil {
		panic("SMALine error")
	}
	if ri.period == 0 {
		panic("erPeriod empty")
	}
	return ri.SMALine.GetByTs(ts)
}
func (ri *RSIIndicator) GetLastByTsAndCount(ts, period int64) ([]model.DataNode, error) {
	return nil, nil
}

func (ri *RSIIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := ri.SMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.GetTs())
		y[i] = v.GetValue()
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: ri.Name()})
	line.AddXAxis(x).AddYAxis(ri.Name(), y, charts.LineOpts{ConnectNulls: false})
	ratioLine.Overlap(line)
}
