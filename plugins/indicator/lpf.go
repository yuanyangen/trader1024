package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator/filter"
)

type LPFIndicator struct {
	period    float64
	name      string
	valueFrom func(in *model.KLineNode) float64
}

func NewLPFIndicator(period float64, name string, source func(in *model.KLineNode) float64) *LPFIndicator {
	LPF := &LPFIndicator{
		period:    period,
		valueFrom: source,
		name:      name,
	}
	return LPF
}

func (LPF *LPFIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {

	if currentNode.LPFData == nil {
		currentNode.LPFData = map[string]float64{}
	}
	data := line.GetAllLastNodeByTs(ts)
	if len(data) > 0 {
		currentNode.LPFData[LPF.Name()] = 0
	}
	in := []float64{}
	for _, v := range data {
		in = append(in, LPF.valueFrom(v))
	}
	out := filter.LowPassFilter(LPF.period, in)
	currentNode.LPFData[LPF.Name()] = out[len(out)-1]
}

func (LPF *LPFIndicator) Name() string {
	return fmt.Sprintf("LPF_%v_%v", LPF.name, LPF.period)
}

func (LPF *LPFIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	return dn.LPFData[LPF.Name()]
}
