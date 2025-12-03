package indicator

import (
	"fmt"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/model"
)

type EMAIndicator struct {
	period    int64
	indicator model.Indicator
}

func NewEMAIndicator(period int64, indicator model.Indicator) *EMAIndicator {
	EMA := &EMAIndicator{
		period:    period,
		indicator: indicator,
	}
	return EMA
}

func (EMA *EMAIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	data := line.GetAllSortedData()

	in := make([]float64, len(data))
	for i, v := range data {
		in[i] = EMA.indicator.GetValueFromKNode(v)
	}
	if len(in) < int(EMA.period) {
		return
	}
	out := talib.Ema(in, int(EMA.period))
	avg := out[len(out)-1]
	if currentNode.EMAData == nil {
		currentNode.EMAData = map[string]float64{}
	}
	currentNode.EMAData[EMA.Name()] = avg
}

func (EMA *EMAIndicator) Name() string {
	return fmt.Sprintf("ema_(%v)_%v", EMA.indicator.Name(), EMA.period)
}

func (EMA *EMAIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil || len(dn.EMAData) == 0 {
		return 0
	}
	return dn.EMAData[EMA.Name()]
}
