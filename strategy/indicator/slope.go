package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type SlopIndicator struct {
	*indicator_base.IndicatorCommon
	inLine   *indicator_base.Line
	slopLine *indicator_base.Line
	period   int64
}

func NewSlopIndicator(kline model.MarketIndicator) model.MarketIndicator {
	sma := &SlopIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          2,
		slopLine:        indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("slop")),
	}
	kline.AddChildrenIndicator(sma)
	return sma
}

func (si *SlopIndicator) Name() string {
	return fmt.Sprintf("Slop_%v", si.period)
}

func (si *SlopIndicator) AddData(ts int64, node any) {
	data, err := si.inLine.GetByTsAndCount(ts, si.period)
	if err != nil {
		si.slopLine.AddData(ts, 0)
		return
	}
	out := data[1].Value - data[0].Value
	si.slopLine.AddData(ts, out)
}
func (si *SlopIndicator) GetAllSortedData() []any {
	return nil
}

func (si *SlopIndicator) GetByTs(ts int64) any {
	if si.slopLine == nil {
		panic("slopLine error")
	}
	if si.period == 0 {
		panic("erPeriod empty")
	}
	data, err := si.slopLine.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
}
func (si *SlopIndicator) GetByTsAndCount(ts, period int64) ([]any, error) {
	return nil, nil
}

func (si *SlopIndicator) GetCurrentFloat(ts int64) float64 {
	v := si.GetByTs(ts)
	f, _ := v.(float64)
	return f
}

func (si *SlopIndicator) DoPlot(p *charts.Page, kline *charts.Kline) {
	allData := si.slopLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: si.Name()})
	line.AddXAxis(x).AddYAxis(si.Name(), y, charts.LineOpts{ConnectNulls: false})
	p.Add(line)
}
