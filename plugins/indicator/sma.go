package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
)

type SMAIndicator struct {
	period    int64
	indicator model.Indicator
}

func NewSMAIndicator(period int64, indicator model.Indicator) *SMAIndicator {
	sma := &SMAIndicator{
		period:    period,
		indicator: indicator,
	}

	return sma
}

func (sma *SMAIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	data, err := line.GetLastNodeByTsAndCount(ts, sma.period)
	if err != nil {
		return
	}

	sum := 0.0
	for _, v := range data {
		sum += sma.indicator.GetValueFromKNode(v)
	}
	avg := sum / float64(sma.period)
	if currentNode.SMAData == nil {
		currentNode.SMAData = map[string]float64{}
	}
	currentNode.SMAData[sma.Name()] = avg
}

func (sma *SMAIndicator) Name() string {
	return fmt.Sprintf("sma_(%v)_%v", sma.indicator.Name(), sma.period)
}

func (sma *SMAIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil || len(dn.SMAData) == 0 {
		return 0
	}
	return dn.SMAData[sma.Name()]
}
