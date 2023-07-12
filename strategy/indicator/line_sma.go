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
	inline      model.MarketIndicator
	KAMALine    *indicator_base.Line
	erPeriod    int64
	shortPeriod int64
	longPeriod  int64
}

func NewLineSmaIndicator(inline model.MarketIndicator, period int64) *LineSmaIndicator {
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

func (kama *LineSmaIndicator) AddData(ts int64, node any) {
	dataI, err := kama.inline.GetByTsAndCount(ts, kama.erPeriod)
	if err != nil {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	data := utils.AnySliceToFloat(dataI)
	if len(data) == 0 {
		kama.KAMALine.AddData(ts, 0)
		return
	}

	out := talib.Sma(data, int(kama.erPeriod))
	avg := out[len(out)-1]
	kama.KAMALine.AddData(ts, avg)
}
func (kama *LineSmaIndicator) GetAllSortedData() []any {
	return nil
}

func (kama *LineSmaIndicator) GetCurrentValue(ts int64) any {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	data, err := kama.KAMALine.GetByTs(ts)
	if err != nil {
		return 0
	} else {
		return data.Value
	}
}

func (kama *LineSmaIndicator) GetCurrentFloat(ts int64) float64 {
	v := kama.GetCurrentValue(ts)
	f, _ := v.(float64)
	return f
}
func (kama *LineSmaIndicator) GetByTsAndCount(ts int64, period int64) ([]any, error) {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	dataI, err := kama.KAMALine.GetByTsAndCount(ts, period)
	if err != nil {
		return nil, err
	}
	res := make([]any, len(dataI))
	for i, v := range dataI {
		res[i] = v.Value
	}

	return res, nil
}
func (kama *LineSmaIndicator) GetByTs(ts int64) any {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	data, err := kama.KAMALine.GetByTs(ts)
	if err != nil {
		return 0.0
	} else {
		return data.Value
	}
}
func (kama *LineSmaIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := kama.KAMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = math.Abs(v.Value * 100)
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: kama.Name()})
	line.AddXAxis(x).AddYAxis(kama.Name(), y, charts.LineOpts{ConnectNulls: false})
	ratioLine.Overlap(line)
}
