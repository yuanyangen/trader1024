package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"math"
)

type VolumnSlopIndicator struct {
	*indicator_base.IndicatorCommon
	inLine   model.ContractIndicator
	slopLine *indicator_base.Line
	period   int64
}

func NewVolumnSlopIndicator(rawLine model.ContractIndicator) model.ContractIndicator {
	slop := &VolumnSlopIndicator{
		inLine:          rawLine,
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          2,
		slopLine:        indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("volumn_slop")),
	}
	rawLine.AddChildrenIndicator(slop)
	return slop
}

func (si *VolumnSlopIndicator) Name() string {
	return fmt.Sprintf("volumn_slop_%v", si.period)
}

func (si *VolumnSlopIndicator) AddData(ts int64, node model.DataNode) {
	dataI, err := si.inLine.GetLastByTsAndCount(ts, si.period)
	if err != nil {
		si.slopLine.AddData(ts, 0)
		return
	}
	knodes := model.NewKnodesFromAny(dataI)
	if len(knodes) == 0 {
		si.slopLine.AddData(ts, 0)
		return
	}

	out := 2 * (knodes[1].Volume - knodes[0].Volume) / (knodes[1].Volume + knodes[0].Volume)
	si.slopLine.AddData(ts, out)
}
func (si *VolumnSlopIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (si *VolumnSlopIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if si.slopLine == nil {
		panic("continousLine error")
	}
	if si.period == 0 {
		panic("erPeriod empty")
	}
	return si.slopLine.GetByTs(ts)
}
func (si *VolumnSlopIndicator) GetLastByTsAndCount(ts, period int64) ([]model.DataNode, error) {
	return si.slopLine.GetLastByTsAndCount(ts, period)
}

func (si *VolumnSlopIndicator) GetCurrentFloat(ts int64) float64 {
	v, _ := si.GetByTs(ts)
	if v == nil {
		return 0
	}
	return v.GetValue()
}

func (si *VolumnSlopIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := si.slopLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.GetTs())
		y[i] = math.Abs(v.GetValue())
	}
	if len(y) > 3 {
		y[0] = 0
		y[1] = 0
		y[2] = 0
	}
	slopLine := charts.NewLine()
	slopLine.SetGlobalOptions(
		//charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		//charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		//charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	slopLine.SetGlobalOptions(charts.TitleOpts{Title: si.Name()}, charts.ToolboxOpts{Show: true})
	slopLine.AddXAxis(x).AddYAxis(si.Name(), y)
	ratioLine.Overlap(slopLine)
	//line := charts.NewLine()
	//line.SetGlobalOptions(charts.TitleOpts{Title: si.CNName()})
	//line.AddXAxis(x).AddYAxis(si.CNName(), y, charts.LineOpts{ConnectNulls: false})
	//p.Add(line)
}
