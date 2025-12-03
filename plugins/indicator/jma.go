package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"math"
)

type JMAIndicator struct {
	power     float64
	indicator model.Indicator
}

func NewJMAIndicator(power float64, indicator model.Indicator) *JMAIndicator {
	sma := &JMAIndicator{
		power:     power,
		indicator: indicator,
	}

	return sma
}

func (jma *JMAIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	data := line.GetAllLastNodeByTs(ts)
	in := []float64{}
	for _, v := range data {
		in = append(in, jma.indicator.GetValueFromKNode(v))
	}
	if currentNode.JMAData == nil {
		currentNode.JMAData = map[string]float64{}
	}
	out := CalculateJMA(in, 5, 0, jma.power)
	currentNode.JMAData[jma.Name()] = out[len(out)-1]
}

func (jma *JMAIndicator) Name() string {
	return fmt.Sprintf("jma_(%v)_%v", jma.indicator.Name(), jma.power)
}

func (jma *JMAIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil || len(dn.JMAData) == 0 {
		return 0
	}
	return dn.JMAData[jma.Name()]
}

// period [0-x]越大平滑效果越差
// phase  [-100,100]越大平滑效果越差
// power [0,1]越小，延时越高，越平滑,
func CalculateJMA(values []float64, period int, phase int, power float64) []float64 {
	if len(values) == 0 || len(values) < period {
		return []float64{0}
	}

	beta := 0.45 * (float64(period) - 1) / (0.45*(float64(period)-1) + 2)
	alpha := math.Pow(beta, power)
	var ph float64
	if phase < -100 {
		ph = 0.5
	} else if phase > 100 {
		ph = 2.5
	} else {
		ph = float64(phase)/100 + 1.5
	}
	ma1, det0, det1, ma2, jma := 0.0, 0.0, 0.0, 0.0, 0.0
	var returnValues []float64
	for _, value := range values {
		if math.IsNaN(value) {
			panic(fmt.Sprintf("[Calculate] invalid value: %v", value))
		}
		if jma == 0 {
			jma = value
			ma1 = value
			ma2 = value
		} else {
			ma1 = (1-alpha)*value + alpha*ma1
			det0 = (value-ma1)*(1-beta) + beta*det0
			ma2 = ma1 + ph*det0
			det1 = math.Pow(1-alpha, 2)*(ma2-jma) + math.Pow(alpha, 2)*det1
			jma = jma + det1
		}

		returnValues = append(returnValues, jma)
	}
	return returnValues
}
