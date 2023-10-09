package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/talib"
	"github.com/yuanyangen/trader1024/engine/utils"
	"math"
)

type LineSmaIndicator struct {
	*indicator_base.IndicatorCommon
	inline      model.ContractIndicator
	KAMALine    *indicator_base.Line
	erPeriod    int64
	shortPeriod int64
	longPeriod  int64
}

func NewLineSmaIndicator(inline model.ContractIndicator, period int64) *LineSmaIndicator {
	kama := &LineSmaIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		erPeriod:        period,
		KAMALine:        indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("line_sma_%v", period)),
		inline:          inline,
	}
	inline.AddChildrenIndicator(kama)
	return kama
}

func (kama *LineSmaIndicator) Name() string {
	return fmt.Sprintf("%v_sma_%v", kama.inline.Name(), kama.erPeriod)
}

func (kama *LineSmaIndicator) AddData(ts int64, node model.DataNode) {
	dataI, err := kama.inline.GetLastByTsAndCount(ts, kama.erPeriod)
	if err != nil {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	data := utils.DataNodeSliceToFloat(dataI)
	if len(data) == 0 {
		kama.KAMALine.AddData(ts, 0)
		return
	}

	out := talib.Sma(data, int(kama.erPeriod))
	avg := out[len(out)-1]
	kama.KAMALine.AddData(ts, avg)
}
func (kama *LineSmaIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (kama *LineSmaIndicator) GetCurrentValue(ts int64) model.DataNode {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	data, _ := kama.KAMALine.GetByTs(ts)
	return data
}

func (kama *LineSmaIndicator) GetCurrentFloat(ts int64) float64 {
	v := kama.GetCurrentValue(ts)
	if v == nil {
		return 0
	}
	return v.GetValue()
}
func (kama *LineSmaIndicator) GetLastByTsAndCount(ts int64, period int64) ([]model.DataNode, error) {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	return kama.KAMALine.GetLastByTsAndCount(ts, period)
}
func (kama *LineSmaIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	return kama.KAMALine.GetByTs(ts)
}
func (kama *LineSmaIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := kama.KAMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.GetTs())
		y[i] = math.Abs(v.GetValue() * 100)
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: kama.Name()})
	line.AddXAxis(x).AddYAxis(kama.Name(), y, charts.LineOpts{ConnectNulls: false})
	ratioLine.Overlap(line)
}
