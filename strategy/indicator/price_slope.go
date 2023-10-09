package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"math"
)

type SlopIndicator struct {
	*indicator_base.IndicatorCommon
	inLine   model.ContractIndicator
	slopLine *indicator_base.Line
	period   int64
}

func NewSlopIndicator(rawLine model.ContractIndicator, period int64) *SlopIndicator {
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

func (si *SlopIndicator) AddData(ts int64, node model.DataNode) {
	dataI, err := si.inLine.GetLastByTsAndCount(ts, si.period)
	if err != nil {
		si.slopLine.AddData(ts, 0)
		return
	}
	data := utils.DataNodeSliceToFloat(dataI)
	if data == nil {
		si.slopLine.AddData(ts, 0)
		return
	}

	out := (data[1] - data[0]) / data[1]
	si.slopLine.AddData(ts, out)
}
func (si *SlopIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (si *SlopIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if si.slopLine == nil {
		panic("continousLine error")
	}
	if si.period == 0 {
		panic("erPeriod empty")
	}
	return si.slopLine.GetByTs(ts)
}
func (si *SlopIndicator) GetLastByTsAndCount(ts, period int64) ([]model.DataNode, error) {
	return si.slopLine.GetLastByTsAndCount(ts, period)
}

func (si *SlopIndicator) GetCurrentFloat(ts int64) float64 {
	v, _ := si.GetByTs(ts)
	if v == nil {
		return 0
	}
	return v.GetValue()
}

func (si *SlopIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := si.slopLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.GetTs())
		if v.GetValue() == math.NaN() {
			y[i] = 0
		} else {
			y[i] = v.GetValue() * 100
		}
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
	//line.SetGlobalOptions(charts.TitleOpts{Title: si.CNName()})
	//line.AddXAxis(x).AddYAxis(si.CNName(), y, charts.LineOpts{ConnectNulls: false})
	//p.Add(line)
}
