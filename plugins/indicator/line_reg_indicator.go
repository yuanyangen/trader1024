package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
)

type LRIndicator struct {
	period    int64
	name      string
	valueFrom model.Indicator
}

func NewLRIndicator(period int64, source model.Indicator) *LRIndicator {
	LR := &LRIndicator{
		period:    period,
		valueFrom: source,
	}
	return LR
}

func (LR *LRIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {

	if currentNode.LRData == nil {
		currentNode.LRData = map[string]float64{}
	}
	data, _ := line.GetLastNodeByTsAndCount(ts, LR.period)
	period := int(LR.period)
	if period > len(data) {
		period = len(data)
	}
	in := make([]float64, len(data))
	for i, v := range data {
		in[i] = LR.valueFrom.GetValueFromKNode(v)
	}

	if len(in) == 1 || period == 1 || period == 0 {
		currentNode.LRData[LR.Name()] = LR.valueFrom.GetValueFromKNode(currentNode)
		return
	}

	out := LinearReg(in, int(period))
	currentNode.LRData[LR.Name()] = out[len(out)-1]
}

func (LR *LRIndicator) Name() string {
	return fmt.Sprintf("LR_%v_%v", LR.valueFrom.Name(), LR.period)
}

func (LR *LRIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	return dn.LRData[LR.Name()]
}

// LinearReg - Linear Regression
func LinearReg(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx - 1
	today := startIdx - 1
	sumX := inTimePeriodF * (inTimePeriodF - 1) * 0.5
	sumXSqr := inTimePeriodF * (inTimePeriodF - 1) * (2*inTimePeriodF - 1) / 6
	divisor := sumX*sumX - inTimePeriodF*sumXSqr
	//initialize values of sumY and sumXY over first (inTimePeriod) input values
	sumXY := 0.0
	sumY := 0.0
	i := inTimePeriod
	for i != 0 {
		i--
		tempValue1 := inReal[today-i]
		sumY += tempValue1
		sumXY += float64(i) * tempValue1
	}
	for today < len(inReal) {
		//sumX and sumXY are already available for first output value
		if today > startIdx-1 {
			tempValue2 := inReal[today-inTimePeriod]
			sumXY += sumY - inTimePeriodF*tempValue2
			sumY += inReal[today] - tempValue2
		}
		m := (inTimePeriodF*sumXY - sumX*sumY) / divisor
		b := (sumY - m*sumX) / inTimePeriodF
		outReal[outIdx] = b + m*(inTimePeriodF-1)
		outIdx++
		today++
	}
	return outReal
}
