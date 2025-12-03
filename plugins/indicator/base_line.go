package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
)

type BLIndicator struct {
	num       float64
	indicator model.Indicator
}

func NewBLIndicator(num float64) *BLIndicator {
	bl := &BLIndicator{
		num: num,
	}

	return bl
}

func (bl *BLIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
}

func (bl *BLIndicator) Name() string {
	return fmt.Sprintf("bl_%v", bl.num)
}

func (bl *BLIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	return bl.num
}
