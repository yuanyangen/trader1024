package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/talib"
)

type CCIIndicator struct {
	period int64
}

func NewCCIIndicator(period int64) *CCIIndicator {
	CCI := &CCIIndicator{
		period: period,
	}
	return CCI
}

func (cci *CCIIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	if currentNode.MomentumData == nil {
		currentNode.MomentumData = map[string]float64{}
	}
	data := line.GetAllLastNodeByTs(ts)
	inClose := []float64{}
	inHigh := []float64{}
	inLow := []float64{}
	for _, v := range data {
		inClose = append(inClose, v.Close)
		inHigh = append(inHigh, v.High)
		inLow = append(inLow, v.Low)
	}
	if len(data) >= int(cci.period) {
		tmpResult := talib.Cci(inHigh, inLow, inClose, int(cci.period))
		currentNode.MomentumData[cci.Name()] = tmpResult[len(tmpResult)-1]
	} else {
		currentNode.MomentumData[cci.Name()] = 0

	}

}

func (cci *CCIIndicator) Name() string {
	return fmt.Sprintf("cci_%v", cci.period)
}

func (cci *CCIIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	//return dn.MomentumData[cci.Name()]
	return dn.MomentumData[cci.Name()] + 2000
}
