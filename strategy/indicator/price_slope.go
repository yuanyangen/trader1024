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
	inLine   model.MarketIndicator
	slopLine *indicator_base.Line
	period   int64
}

func NewSlopIndicator(rawLine model.MarketIndicator, period int64) model.MarketIndicator {
	slop := &SlopIndicator{
		inLine:          rawLine,
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          period,
		slopLine:        indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("slop")),
	}
	rawLine.AddChildrenIndicator(slop)
	return slop
}

func (si *SlopIndicator) Name() string {
	return fmt.Sprintf("Slop_%v", si.period)
}

func (si *SlopIndicator) AddData(ts int64, node any) {
	dataI, err := si.inLine.GetByTsAndCount(ts, si.period)
	if err != nil {
		si.slopLine.AddData(ts, 0)
		return
	}
	data := utils.AnySliceToFloat(dataI)
	if data == nil {
		si.slopLine.AddData(ts, 0)
		return
	}

	out := (data[1] - data[0]) / data[1]
	si.slopLine.AddData(ts, out)
}
func (si *SlopIndicator) GetAllSortedData() []any {
	return nil
}

func (si *SlopIndicator) GetByTs(ts int64) any {
	if si.slopLine == nil {
		panic("continousLine error")
	}
	if si.period == 0 {
		panic("erPeriod empty")
	}
	data, err := si.slopLine.GetByTs(ts)
	if err != nil {
		return 0.0
	} else {
		return data.Value
	}
}
func (si *SlopIndicator) GetByTsAndCount(ts, period int64) ([]any, error) {
	rawData, err := si.slopLine.GetByTsAndCount(ts, period)
	if err != nil {
		return nil, err
	}
	res := []any{}
	for _, v := range rawData {
		res = append(res, v.Value)
	}
	return res, nil
}

func (si *SlopIndicator) GetCurrentFloat(ts int64) float64 {
	v := si.GetByTs(ts)
	f, _ := v.(float64)
	return f
}

func (si *SlopIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := si.slopLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = (v.Value * 100)
	}
	if len(y) > 3 {
		y[0] = 0
		y[1] = 0
		y[2] = 0
	}
	//ratioLine := charts.NewLine()
	ratioLine.SetGlobalOptions(
		//charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		//charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		//charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	ratioLine.SetGlobalOptions(charts.TitleOpts{Title: si.Name()}, charts.ToolboxOpts{Show: true})
	ratioLine.AddXAxis(x).AddYAxis(si.Name(), y)
	//ratioLine.Overlap(bar)
	//line := charts.NewLine()
	//line.SetGlobalOptions(charts.TitleOpts{Title: si.Name()})
	//line.AddXAxis(x).AddYAxis(si.Name(), y, charts.LineOpts{ConnectNulls: false})
	//p.Add(line)
}
