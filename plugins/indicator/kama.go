package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/talib"
)

type KAMAIndicator struct {
	period    int64
	name      string
	valueFrom func(in *model.KLineNode) float64
	toData    func(node *model.KLineNode, float642 float64, period int64)
}

func NewKAMAIndicator(period int64, name string, valueFrom func(in *model.KLineNode) float64) *KAMAIndicator {
	KAMA := &KAMAIndicator{
		period:    period,
		name:      name,
		valueFrom: valueFrom,
	}
	return KAMA
}

func (kama *KAMAIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	lastNode, _ := line.GetLastNodeByTs(ts)
	if lastNode == nil {
		return
	}
	data, _ := line.GetLastNodeByTsAndCount(ts, kama.period+1)
	in := make([]float64, len(data))
	for i, v := range data {
		in[i] = kama.valueFrom(v)
	}
	if len(in) < int(kama.period)+1 {
		return
	}

	lastData := kama.GetValueFromKNode(lastNode)
	out := talib.CustomKama(in, int(kama.period), 2, 30, lastData)
	avg := out[len(out)-1]
	if currentNode.KAMAData == nil {
		currentNode.KAMAData = map[string]float64{}
	}
	currentNode.KAMAData[kama.Name()] = avg
}

func (kama *KAMAIndicator) Name() string {
	return fmt.Sprintf("KAMA_%v_%v", kama.name, kama.period)
}

func (kama *KAMAIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil || len(dn.KAMAData) == 0 {
		return 0
	}
	return dn.KAMAData[kama.Name()]
}
