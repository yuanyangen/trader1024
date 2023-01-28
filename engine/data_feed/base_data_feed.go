package data_feed

import "github.com/yuanyangen/trader1024/engine/model"

type BaseDataFeed struct {
	dataFeedMeta  *model.DataMeta
	outerChannels []chan *model.Data
}

func (bdf *BaseDataFeed) RegisterChan(out chan *model.Data) {
	if bdf.outerChannels == nil {
		bdf.outerChannels = []chan *model.Data{out}
	} else {
		bdf.outerChannels = append(bdf.outerChannels, out)
	}
}

func (bdf *BaseDataFeed) SendData(out *model.Data) {
	out.DataMeta = bdf.dataFeedMeta
	for _, v := range bdf.outerChannels {
		v <- out
	}
}

func (bdf *BaseDataFeed) SetDataMeta(dm *model.DataMeta) {
	bdf.dataFeedMeta = dm
}

func (bdf *BaseDataFeed) GetMeta() *model.DataMeta {
	return bdf.dataFeedMeta
}
