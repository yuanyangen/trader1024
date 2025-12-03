package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"math"
)

type ZDKAMAIndicator struct {
	period     int64
	kamaPeriod int64
	fastEma    int64
	slowEma    int64
	name       string
	valueFrom  func(in *model.KLineNode) float64
}

func NewZDKAMAIndicator(period int64, name string, source func(in *model.KLineNode) float64) *ZDKAMAIndicator {
	ZDKAMA := &ZDKAMAIndicator{
		period:     period,
		valueFrom:  source,
		kamaPeriod: 5,
		fastEma:    2,
		slowEma:    30,
		name:       name,
	}
	return ZDKAMA
}

func (ZDKAMA *ZDKAMAIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {

	if currentNode.ZDKAMAData == nil {
		currentNode.ZDKAMAData = map[string]float64{}
	}
	data, err := line.GetLastNodeByTsAndCount(ts, 6)
	if err != nil {
		if len(data) > 1 {
			currentNode.ZDKAMAData[ZDKAMA.Name()] = (ZDKAMA.valueFrom(data[len(data)-1]) + ZDKAMA.valueFrom(data[len(data)-2])) / 2
		} else if len(data) == 1 {
			currentNode.ZDKAMAData[ZDKAMA.Name()] = ZDKAMA.valueFrom(data[len(data)-1])
		}
		return
	}

	alpha := 2.0 / float64(ZDKAMA.period+1)
	newAlpha := ZDKAMA.getAlpha(line, ts, ZDKAMA.kamaPeriod, ZDKAMA.fastEma, ZDKAMA.slowEma)
	if newAlpha != 0 {
		alpha = newAlpha
	}
	x_0 := ZDKAMA.valueFrom(data[len(data)-1])
	x_1 := ZDKAMA.valueFrom(data[len(data)-2])
	x_2 := ZDKAMA.valueFrom(data[len(data)-3])
	y_1 := ZDKAMA.GetValueFromKNode(data[len(data)-2])
	y_2 := ZDKAMA.GetValueFromKNode(data[len(data)-3])

	// https://mp.weixin.qq.com/s?__biz=MzIxNzUyNTI4MA==&mid=2247484438&idx=1&sn=3f2c8d47efcc30ed734dbd6c62e2efe9&chksm=97f93959a08eb04ff37ed759372ab10d6f24931ab05aeb125397e10e7c92036fac6b6829ff4f&scene=21#wechat_redirect
	out := (alpha-(alpha*alpha)/4)*x_0 + ((alpha*alpha)/2)*x_1 - (alpha-(3*alpha*alpha)/4)*x_2 +
		2*(1-alpha)*y_1 - (1-alpha)*(1-alpha)*y_2

	currentNode.ZDKAMAData[ZDKAMA.Name()] = out
}

func (ZDKAMA *ZDKAMAIndicator) getAlpha(line *model.KLine, ts int64, period, fastEma, slowEma int64) float64 {
	data, err := line.GetLastNodeByTsAndCount(ts, period+1)
	if err != nil {
		return 0
	}

	values := make([]float64, len(data))
	for i, v := range data {
		values[i] = ZDKAMA.valueFrom(v)
	}
	fastAlpha := 2.0 / (float64(fastEma) + 1.0)
	slowAlpha := 2 / (float64(slowEma) + 1.0)
	mon := math.Abs(values[len(values)-1] - values[0])
	vol := 0.0
	er := 0.0
	for i := 1; i < len(values); i++ {
		vol += math.Abs(values[i] - values[i-1])
	}
	if vol != 0 {
		er = mon / vol
	}
	alpha := math.Pow(er*(fastAlpha-slowAlpha)+slowAlpha, 2)
	return alpha
}

func (ZDKAMA *ZDKAMAIndicator) Name() string {
	return fmt.Sprintf("ZDKAMA_%v_%v", ZDKAMA.name, ZDKAMA.period)
}

func (ZDKAMA *ZDKAMAIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	return dn.ZDKAMAData[ZDKAMA.Name()]
}
