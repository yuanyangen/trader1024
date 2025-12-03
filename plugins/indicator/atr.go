package indicator

import (
	"fmt"
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/model"
)

type ATRIndicator struct {
	period int64
}

func NewATRIndicator(period int64) *ATRIndicator {
	sma := &ATRIndicator{
		period: period,
	}
	return sma
}

func (sma *ATRIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	data_crawler, err := line.GetLastNodeByTsAndCount(ts, sma.period)

	if err != nil {
		return
	}

	hight := make([]float64, len(data_crawler))
	low := make([]float64, len(data_crawler))
	clo := make([]float64, len(data_crawler))

	for i, v := range data_crawler {
		hight[i] = v.High
		low[i] = v.Low
		clo[i] = v.Close
	}
	out := talib.Atr(hight, low, clo, int(sma.period)-1)
	avg := out[len(out)-1]
	sma.atrDataSaver(currentNode, avg, sma.period)
}

func (sma *ATRIndicator) atrDataSaver(currentNode *model.KLineNode, avg float64, period int64) {
	if currentNode.AtrData == nil {
		currentNode.AtrData = make(map[int64]float64)
	}
	currentNode.AtrData[period] = avg
}

func (sma *ATRIndicator) Name() string {
	return fmt.Sprintf("atr_%v", sma.period)
}

func (sma *ATRIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil || dn.AtrData == nil {
		return 0
	}

	return dn.AtrData[sma.period]
}
