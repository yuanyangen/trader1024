package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/talib"
)

type BOPIndicator struct {
}

func NewBOPIndicator() *BOPIndicator {
	BOP := &BOPIndicator{}
	return BOP
}

func (bop *BOPIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	if currentNode.MomentumData == nil {
		currentNode.MomentumData = map[string]float64{}
	}
	data := line.GetAllLastNodeByTs(ts)
	inClose := []float64{}
	inHigh := []float64{}
	inLow := []float64{}
	inVolumn := []float64{}
	for _, v := range data {
		inClose = append(inClose, v.Close)
		inHigh = append(inHigh, v.High)
		inLow = append(inLow, v.Low)
		inVolumn = append(inVolumn, v.Volume)
	}
	tmpResult := talib.Bop(inHigh, inLow, inClose, inVolumn)
	currentNode.MomentumData[bop.Name()] = tmpResult[len(tmpResult)-1]
}

func (bop *BOPIndicator) Name() string {
	return fmt.Sprintf("bop")
}

func (bop *BOPIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	//return dn.MomentumData[bop.Name()]
	return dn.MomentumData[bop.Name()]
}
