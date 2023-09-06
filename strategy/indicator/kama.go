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

type KAMAIndicator struct {
	*indicator_base.IndicatorCommon
	kline       model.MarketIndicator
	KAMALine    *indicator_base.Line
	erPeriod    int64
	shortPeriod int64
	longPeriod  int64
}

func NewKAMAIndicator(kline model.MarketIndicator, period, shortPeriod, longPeriod int64) *KAMAIndicator {
	kama := &KAMAIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		erPeriod:        period,
		shortPeriod:     shortPeriod,
		longPeriod:      longPeriod,
		KAMALine:        indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("kama_%v", period)),
		kline:           kline,
	}
	kline.AddChildrenIndicator(kama)
	return kama
}

func (kama *KAMAIndicator) Name() string {
	return fmt.Sprintf("KAMA_%v", kama.erPeriod)
}

func (kama *KAMAIndicator) AddData(ts int64, node any) {
	dataI, err := kama.kline.GetLastByTsAndCount(ts, kama.erPeriod+1)
	if err != nil {
		kama.KAMALine.AddData(ts, 0)
		return
	}
	data := model.NewKnodesFromAny(dataI)
	in := make([]float64, len(data))
	for i, knode := range data {
		in[i] = (knode.Close + knode.Open) / 2
	}
	out := talib.CustomKama(in, int(kama.erPeriod), int(kama.shortPeriod), int(kama.longPeriod), kama.GetCurrentFloat(data[len(data)-2].TimeStamp))
	avg := out[len(out)-1]
	kama.KAMALine.AddData(ts, avg)
}
func (kama *KAMAIndicator) GetAllSortedData() []any {
	return nil
}

func (kama *KAMAIndicator) GetCurrentValue(ts int64) any {
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

func (kama *KAMAIndicator) GetCurrentFloat(ts int64) float64 {
	v := kama.GetCurrentValue(ts)
	f, _ := v.(float64)
	return f
}
func (kama *KAMAIndicator) GetLastByTsAndCount(ts int64, period int64) ([]any, error) {
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
func (kama *KAMAIndicator) GetByTs(ts int64) any {
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
func (kama *KAMAIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := kama.KAMALine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToString(v.TimeStamp)
		y[i] = v.Value
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: kama.Name()})
	line.AddXAxis(x).AddYAxis(kama.Name(), y, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}

func (kama *KAMAIndicator) kama(values []float64, period, fastEma, slowEma int) []float64 {
	if len(values) <= period {
		panic("data error")
	}
	fastAlpha := 2.0 / (float64(fastEma) + 1.0)
	slowAlpha := 2 / (float64(slowEma) + 1.0)
	kamaV := float64(0.0)
	result := []float64{}
	change := []float64{}
	for i := range values {
		if i == 0 {
			continue
		}
		change = append(change, math.Abs(values[i]-values[i-1]))
		if i-period < 0 {
			continue
		}
		mon := math.Abs(values[i] - values[i-period])
		vol := 0.0
		for j := len(change) - period - 1; j < len(change); j++ {
			vol += change[j]
		}
		er := 0.0
		if vol != 0 {
			er = mon / vol
		}
		alpha := math.Pow(er*(fastAlpha-slowAlpha)+slowAlpha, 2)
		kamaV = alpha*values[i] + (1-alpha)*kamaV

		result[i] = kamaV
	}
	return result
}
