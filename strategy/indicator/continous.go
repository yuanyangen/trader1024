package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type ContinousIndicator struct {
	*indicator_base.IndicatorCommon
	kline         model.MarketIndicator
	continousLine *indicator_base.Line
	period        int64
}

func NewContinousIndicator(rawLine model.MarketIndicator) model.MarketIndicator {
	slop := &ContinousIndicator{
		kline:           rawLine,
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          2,
		continousLine:   indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("continous")),
	}
	rawLine.AddChildrenIndicator(slop)
	return slop
}

func (si *ContinousIndicator) Name() string {
	return fmt.Sprintf("continus")
}

func (si *ContinousIndicator) AddData(ts int64, node any) {
	dataI, err := si.kline.GetLastByTsAndCount(ts, 2)
	if err != nil {
		si.continousLine.AddData(ts, 0)
		return
	}
	knodes := model.NewKnodesFromAny(dataI)
	if knodes == nil {
		si.continousLine.AddData(ts, 0)
		return
	}

	out := (knodes[1].Open - knodes[0].Close) / (knodes[1].Open + knodes[0].Close)
	si.continousLine.AddData(ts, out)
}
func (si *ContinousIndicator) GetAllSortedData() []any {
	return nil
}

func (si *ContinousIndicator) GetByTs(ts int64) any {
	if si.continousLine == nil {
		panic("continous error")
	}
	if si.period == 0 {
		panic("continous empty")
	}
	data, err := si.continousLine.GetByTs(ts)
	if err != nil {
		return 0.0
	} else {
		return data.Value
	}
}
func (si *ContinousIndicator) GetLastByTsAndCount(ts, period int64) ([]any, error) {
	rawData, err := si.continousLine.GetLastByTsAndCount(ts, period)
	if err != nil {
		return nil, err
	}
	res := []any{}
	for _, v := range rawData {
		res = append(res, v.Value)
	}
	return res, nil
}

func (si *ContinousIndicator) GetCurrentFloat(ts int64) float64 {
	v := si.GetByTs(ts)
	f, _ := v.(float64)
	return f
}

func (si *ContinousIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := si.continousLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value * 100
	}
	if len(y) > 3 {
		y[0] = 0
		y[1] = 0
		y[2] = 0
	}
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.TitleOpts{Title: si.Name()}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(x).AddYAxis(si.Name(), y)
	ratioLine.Overlap(bar)
	//line := charts.NewLine()
	//line.SetGlobalOptions(charts.TitleOpts{Title: si.CNName()})
	//line.AddXAxis(x).AddYAxis(si.CNName(), y, charts.LineOpts{ConnectNulls: false})
	//p.Add(line)
}
