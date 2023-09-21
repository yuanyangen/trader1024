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

func (kama *LineKAMAIndicator) AddData(ts int64, node model.DataNode) {
	dataI, err := kama.inline.GetLastByTsAndCount(ts, kama.erPeriod+1)
	if err != nil {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	data := utils.DataNodeSliceToFloat(dataI)
	if len(data) == 0 {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	lastDatas, _ := kama.KAMALine.GetLastByTsAndCount(ts, 2)
	if len(lastDatas) < 2 {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	lastData := lastDatas[0].GetValue()
	out := talib.CustomKama(data, int(kama.erPeriod), int(kama.shortPeriod), int(kama.longPeriod), lastData)
	avg := out[len(out)-1]
	kama.KAMALine.AddData(ts, avg)
}
func (kama *LineKAMAIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (kama *LineKAMAIndicator) GetCurrentValue(ts int64) model.DataNode {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	data, _ := kama.KAMALine.GetByTs(ts)
	return data
}

func (kama *LineKAMAIndicator) GetCurrentFloat(ts int64) float64 {
	v := kama.GetCurrentValue(ts)
	if v == nil {
		return 0
	}
	return v.GetValue()
}
func (kama *LineKAMAIndicator) GetLastByTsAndCount(ts int64, period int64) ([]model.DataNode, error) {
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
	return dataI, nil
}
func (kama *LineKAMAIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if kama.KAMALine == nil {
		panic("KAMALine error")
	}
	if kama.erPeriod == 0 {
		panic("erPeriod empty")
	}
	return kama.KAMALine.GetByTs(ts)
}
func (kama *LineKAMAIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := kama.KAMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.GetTs())
		y[i] = math.Abs(v.GetValue() * 100)
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: kama.Name()})
	line.AddXAxis(x).AddYAxis(kama.Name(), y, charts.LineOpts{ConnectNulls: false})
	ratioLine.Overlap(line)
}
