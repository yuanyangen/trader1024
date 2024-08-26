package indicator

import (
	"github.com/markcheno/go-talib"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
)

type SMAIndicator struct {
	period int64
}

func NewSMAIndicator(period int64) *SMAIndicator {
	sma := &SMAIndicator{
		period: period,
	}
	return sma
}

func (sma *SMAIndicator) FillIndicatorToLine(line *model.KLine, ts int64) {
	data, err := line.GetLastByTsAndCount(ts, sma.period)
	if err != nil {
		return
	}

	in := make([]float64, len(data))
	for i, v := range data {
		in[i] = v.DailyNode.Close
	}
	out := talib.Sma(in, int(sma.period))
	avg := out[len(out)-1]
	currentNode, err := line.GetByTs(ts)
	if err != nil {
		logs.Info("get data_crawler from line error %v", ts)
		return
	}
	if currentNode.DailyNode.SmaData == nil {
		currentNode.DailyNode.SmaData = make(map[int64]float64)
	}
	currentNode.DailyNode.SmaData[sma.period] = avg
}
