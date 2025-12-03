package indicator

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
)

type RawKlineIndicator struct {
	name      string
	valueFrom func(in *model.KLineNode) float64
}

func NewRawKlineIndicator(name string) *RawKlineIndicator {
	RawKline := &RawKlineIndicator{
		name: name,
	}
	return RawKline
}

func (RawKline *RawKlineIndicator) FillIndicatorToLine(line *model.KLine, ts int64, currentNode *model.KLineNode) {
	if currentNode.RawKlineData == nil {
		currentNode.RawKlineData = make(map[string]float64)
	}
	currentNode.RawKlineData[RawKline.Name()] = RawKline.valueFrom(currentNode)
}

func (RawKline *RawKlineIndicator) Name() string {
	return fmt.Sprintf("%v", RawKline.name)
}

const DataSourceClose = "close"
const DataSourceOpen = "open"
const DataSourceHigh = "high"
const DataSourceLow = "low"
const DataSourceAVG = "avg"
const DataSourceFullAVG = "full_avg"
const DataSourceCloseSubOpen = "close_s_open"

func (RawKline *RawKlineIndicator) GetValueFromKNode(dn *model.KLineNode) float64 {
	if dn == nil {
		return 0
	}
	switch RawKline.name {
	case DataSourceClose:
		return dn.Close
	case DataSourceHigh:
		return dn.High
	case DataSourceLow:
		return dn.Low
	case DataSourceOpen:
		return dn.Open
	case DataSourceAVG:
		return (dn.Open + dn.Close) / 2
	case DataSourceFullAVG:
		return (dn.Open + dn.Close + dn.High + dn.Low) / 4

	case DataSourceCloseSubOpen:
		return dn.Close - dn.Open
	default:
		return 0
	}
}
