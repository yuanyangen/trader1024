package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
)

type ZDMAIndicator struct {
	period    int64
	indicator model.Indicator
}

func NewZDMAIndicator(period int64, indicator model.Indicator) *ZDMAIndicator {
	ZDMA := &ZDMAIndicator{
		period:    period,
		indicator: indicator,
	}
	return ZDMA
}

func (zdma *ZDMAIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	valueFrom := zdma.indicator.GetValueFromKNode

	if currentNode.ZDMAData == nil {
		currentNode.ZDMAData = map[string]float64{}
	}
	data, err := line.GetLastNodeByTsAndCount(ts, 6)
	if err != nil {
		if len(data) > 1 {
			currentNode.ZDMAData[zdma.Name()] = (valueFrom(data[len(data)-1]) + valueFrom(data[len(data)-2])) / 2
		} else if len(data) == 1 {
			currentNode.ZDMAData[zdma.Name()] = valueFrom(data[len(data)-1])
		}
		return
	}

	alpha := 2.0 / float64(zdma.period+1)
	x_0 := valueFrom(data[len(data)-1])
	x_1 := valueFrom(data[len(data)-2])
	x_2 := valueFrom(data[len(data)-3])
	y_1 := zdma.GetValueFromKNode(data[len(data)-2])
	y_2 := zdma.GetValueFromKNode(data[len(data)-3])

	// https://mp.weixin.qq.com/s?__biz=MzIxNzUyNTI4MA==&mid=2247484438&idx=1&sn=3f2c8d47efcc30ed734dbd6c62e2efe9&chksm=97f93959a08eb04ff37ed759372ab10d6f24931ab05aeb125397e10e7c92036fac6b6829ff4f&scene=21#wechat_redirect
	out := (alpha-(alpha*alpha)/4)*x_0 + ((alpha*alpha)/2)*x_1 - (alpha-(3*alpha*alpha)/4)*x_2 +
		2*(1-alpha)*y_1 - (1-alpha)*(1-alpha)*y_2

	currentNode.ZDMAData[zdma.Name()] = out
}

func (zdma *ZDMAIndicator) Name() string {
	return fmt.Sprintf("zdma_(%v)_%v", zdma.indicator.Name(), zdma.period)
}

func (zdma *ZDMAIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	return dn.ZDMAData[zdma.Name()]
}
