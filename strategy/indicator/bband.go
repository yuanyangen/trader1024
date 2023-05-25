package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/talib"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type BBANDIndicator struct {
	*indicator_base.BaseLine
	*indicator_base.IndicatorCommon
	kline          model.MarketIndicator
	BBANDUpperLine *indicator_base.Line
	BBANDMidLine   *indicator_base.Line
	BBANDDownLine  *indicator_base.Line
	erPeriod       int64
}

func NewBBANDIndicator(kline model.MarketIndicator, period int64) *BBANDIndicator {
	bband := &BBANDIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		erPeriod:        period,
		BBANDUpperLine:  indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("bband_uper_%v", period)),
		BBANDMidLine:    indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("bband_mid_%v", period)),
		BBANDDownLine:   indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("bband_low_%v", period)),
		kline:           kline,
	}
	kline.AddChildrenIndicator(bband)
	return bband
}

func (bband *BBANDIndicator) Name() string {
	return fmt.Sprintf("BBAND_%v", bband.erPeriod)
}

func (bband *BBANDIndicator) AddData(ts int64, node any) {
	data, err := bband.kline.GetByTsAndCount(ts, bband.erPeriod+1)
	if err != nil {
		bband.BBANDUpperLine.AddData(ts, 0)
		bband.BBANDMidLine.AddData(ts, 0)
		bband.BBANDDownLine.AddData(ts, 0)
		return
	}
	in := make([]float64, len(data))
	for i, v := range data {
		kNode := v.(*model.KNode)
		in[i] = (kNode.Close + kNode.Open) / 2
	}
	up, mid, low := talib.BBands(in, int(bband.erPeriod), 1, 1, 0)
	bband.BBANDUpperLine.AddData(ts, up[len(up)-1])
	bband.BBANDMidLine.AddData(ts, mid[len(mid)-1])
	bband.BBANDDownLine.AddData(ts, low[len(low)-1])
	bband.TriggerChildren(ts, node)
}
func (bband *BBANDIndicator) GetAllSortedData() []any {
	return nil
}

func (bband *BBANDIndicator) GetByTs(ts int64) any {
	if bband.BBANDUpperLine == nil {
		panic("BBANDLine error")
	}
	if bband.erPeriod == 0 {
		panic("erPeriod empty")
	}
	data, err := bband.BBANDUpperLine.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
}

func (bband *BBANDIndicator) GetUpperFloat(ts int64) float64 {
	return bband.doGetFloat(ts, bband.BBANDUpperLine)
}

func (bband *BBANDIndicator) GetMidFloat(ts int64) float64 {
	return bband.doGetFloat(ts, bband.BBANDMidLine)
}
func (bband *BBANDIndicator) GetLowFloat(ts int64) float64 {
	return bband.doGetFloat(ts, bband.BBANDDownLine)
}

func (bband *BBANDIndicator) doGetFloat(ts int64, l *indicator_base.Line) float64 {
	v := bband.doGetValue(ts, l)
	f, _ := v.(float64)
	return f
}

func (bband *BBANDIndicator) doGetValue(ts int64, l *indicator_base.Line) any {
	if l == nil {
		panic("BBANDLine error")
	}
	if bband.erPeriod == 0 {
		panic("erPeriod empty")
	}
	data, err := l.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
}

func (bband *BBANDIndicator) DoPlot(p *charts.Page, kline *charts.Kline) {
	bband.doPlotOneLine(kline, bband.BBANDUpperLine)
	bband.doPlotOneLine(kline, bband.BBANDMidLine)
	bband.doPlotOneLine(kline, bband.BBANDDownLine)
}

func (bband *BBANDIndicator) doPlotOneLine(kline *charts.Kline, l *indicator_base.Line) {
	allData := l.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: bband.Name()})
	line.AddXAxis(x).AddYAxis(bband.Name(), y, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}
