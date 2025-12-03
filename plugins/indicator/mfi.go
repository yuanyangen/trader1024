package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/talib"
)

type MFIIndicator struct {
	period int64
}

func NewMFIIndicator(period int64) *MFIIndicator {
	MFI := &MFIIndicator{
		period: period,
	}
	return MFI
}

func (mfi *MFIIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
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
	if len(data) >= int(mfi.period)+1 {
		tmpResult := talib.Mfi(inHigh, inLow, inClose, inVolumn, int(mfi.period))
		currentNode.MomentumData[mfi.Name()] = tmpResult[len(tmpResult)-1]
	} else {
		currentNode.MomentumData[mfi.Name()] = 0

	}

}

func (mfi *MFIIndicator) Name() string {
	return fmt.Sprintf("mfi_%v", mfi.period)
}

func (mfi *MFIIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}

	//return dn.MomentumData[mfi.Name()]
	return dn.Close - dn.MomentumData[mfi.Name()]*4
}
