package data_feed

import (
	"github.com/yuanyangen/trader1024/engine/indicator"
)

type BaseDataFeed struct {
	dataFeedMeta  *indicator.DataMeta
	outerChannels []chan *indicator.Data
}

func (bdf *BaseDataFeed) RegisterChan(out chan *indicator.Data) {
	if bdf.outerChannels == nil {
		bdf.outerChannels = []chan *indicator.Data{out}
	} else {
		bdf.outerChannels = append(bdf.outerChannels, out)
	}
}

func (bdf *BaseDataFeed) SendData(out *indicator.Data) {
	out.DataMeta = bdf.dataFeedMeta
	for _, v := range bdf.outerChannels {
		v <- out
	}
}

func (bdf *BaseDataFeed) SetDataMeta(dm *indicator.DataMeta) {
	bdf.dataFeedMeta = dm
}

func (bdf *BaseDataFeed) GetMeta() *indicator.DataMeta {
	return bdf.dataFeedMeta
}
