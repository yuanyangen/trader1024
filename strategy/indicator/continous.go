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
	kline         model.ContractIndicator
	continousLine *indicator_base.Line
	period        int64
}

func NewContinousIndicator(rawLine model.ContractIndicator) model.ContractIndicator {
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

func (si *ContinousIndicator) AddData(ts int64, node model.DataNode) {
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
func (si *ContinousIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (si *ContinousIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if si.continousLine == nil {
		panic("continous error")
	}
	if si.period == 0 {
		panic("continous empty")
	}
	return si.continousLine.GetByTs(ts)
}
func (si *ContinousIndicator) GetLastByTsAndCount(ts, period int64) ([]model.DataNode, error) {
	rawData, err := si.continousLine.GetLastByTsAndCount(ts, period)
	if err != nil {
		return nil, err
	}
	return rawData, nil
}

func (si *ContinousIndicator) GetCurrentFloat(ts int64) float64 {
	v, _ := si.GetByTs(ts)
	if v == nil {
		return 0
	}
	return v.GetValue()
}

func (si *ContinousIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := si.continousLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.GetTs())
		y[i] = v.GetValue() * 100
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
