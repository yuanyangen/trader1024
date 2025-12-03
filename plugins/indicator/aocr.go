package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"math"
)

type AOCRIndicator struct {
	period int64
}

// 计算开盘价与收盘价之间的波动
func NewAOCRIndicator(period int64) *AOCRIndicator {
	sma := &AOCRIndicator{
		period: period,
	}
	return sma
}

func (aocr *AOCRIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	data_crawler, err := line.GetLastNodeByTsAndCount(ts, aocr.period)

	if err != nil {
		return
	}

	sum := 0.0

	for _, v := range data_crawler {
		sum += math.Abs(v.Close - v.Open)
	}
	avg := sum / float64(aocr.period)
	aocr.AOCRDataSaver(currentNode, avg, aocr.period)
}

func (aocr *AOCRIndicator) AOCRDataSaver(currentNode *model.KLineNode, avg float64, period int64) {
	if currentNode.AOCRData == nil {
		currentNode.AOCRData = make(map[int64]float64)
	}
	currentNode.AOCRData[period] = avg
}

func (aocr *AOCRIndicator) Name() string {
	return fmt.Sprintf("AOCR_%v", aocr.period)
}

func (aocr *AOCRIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil || dn.AOCRData == nil {
		return 0
	}

	return dn.AOCRData[aocr.period]
}
