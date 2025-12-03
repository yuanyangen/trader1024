package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
)

type SlopIndicator struct {
	period          int64
	sourceIndicator model.Indicator
}

func NewSlopIndicator(period int64, sourceIndicator model.Indicator) *SlopIndicator {
	Slop := &SlopIndicator{
		period:          period,
		sourceIndicator: sourceIndicator,
	}
	return Slop
}

func (Slop *SlopIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	data, err := line.GetLastNodeByTsAndCount(ts, Slop.period)
	if err != nil {
		return
	}

	in := make([]float64, len(data))
	for i, v := range data {
		if Slop.sourceIndicator.GetValueFromKNode(v) == 0 {
			return
		}
		in[i] = Slop.sourceIndicator.GetValueFromKNode(v)
	}
	if len(in) == 0 {
		return
	}
	out := (in[len(in)-1] - in[0]) / in[len(in)-1]

	if currentNode.SlopData == nil {
		currentNode.SlopData = map[string]float64{}
	}
	currentNode.SlopData[Slop.sourceIndicator.Name()] = out
}

func (Slop *SlopIndicator) Name() string {
	return fmt.Sprintf("slop_%v", Slop.sourceIndicator.Name())
}

func (Slop *SlopIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	return dn.SlopData[Slop.sourceIndicator.Name()]
}
