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

type LineKAMAIndicator struct {
	*indicator_base.IndicatorCommon
	inline      model.MarketIndicator
	KAMALine    *indicator_base.Line
	erPeriod    int64
	shortPeriod int64
	longPeriod  int64
}

func NewLineKAMAIndicator(inline model.MarketIndicator, period, shortPeriod, longPeriod int64) *LineKAMAIndicator {
	kama := &LineKAMAIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		erPeriod:        period,
		shortPeriod:     shortPeriod,
		longPeriod:      longPeriod,
		KAMALine:        indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("line_kama_%v", period)),
		inline:          inline,
	}
	inline.AddChildrenIndicator(kama)
	return kama
}

func (kama *LineKAMAIndicator) Name() string {
	return fmt.Sprintf("%v_KAMA_%v", kama.inline.Name(), kama.erPeriod)
}

func (kama *LineKAMAIndicator) AddData(ts int64, node any) {
	dataI, err := kama.inline.GetLastByTsAndCount(ts, kama.erPeriod+1)
	if err != nil {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	data := utils.AnySliceToFloat(dataI)
	if len(data) == 0 {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	lastDatas, _ := kama.KAMALine.GetLastByTsAndCount(ts, 2)
	if len(lastDatas) < 2 {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	lastData := lastDatas[0].Value
	out := talib.CustomKama(data, int(kama.erPeriod), int(kama.shortPeriod), int(kama.longPeriod), lastData)
	avg := out[len(out)-1]
	kama.KAMALine.AddData(ts, avg)
}
func (kama *LineKAMAIndicator) GetAllSortedData() []any {
	return nil
}

func (kama *LineKAMAIndicator) GetCurrentValue(ts int64) any {
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

func (kama *LineKAMAIndicator) GetCurrentFloat(ts int64) float64 {
	v := kama.GetCurrentValue(ts)
	f, _ := v.(float64)
	return f
}
func (kama *LineKAMAIndicator) GetLastByTsAndCount(ts int64, period int64) ([]any, error) {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	dataI, err := kama.KAMALine.GetLastByTsAndCount(ts, period)
	if err != nil {
		return nil, err
	}
	res := make([]any, len(dataI))
	for i, v := range dataI {
		res[i] = v.Value
	}

	return res, nil
}
func (kama *LineKAMAIndicator) GetByTs(ts int64) any {
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
func (kama *LineKAMAIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
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
