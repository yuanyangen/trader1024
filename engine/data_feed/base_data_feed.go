package data_feed

type BaseDataFeed struct {
	dataFeedMeta  *DataMeta
	outerChannels []chan *Data
}

func (bdf *BaseDataFeed) RegisterChan(out chan *Data) {
	if bdf.outerChannels == nil {
		bdf.outerChannels = []chan *Data{out}
	} else {
		bdf.outerChannels = append(bdf.outerChannels, out)
	}
}

func (bdf *BaseDataFeed) SendData(out *Data) {
	out.DataMeta = bdf.dataFeedMeta
	for _, v := range bdf.outerChannels {
		v <- out
	}
}

func (bdf *BaseDataFeed) SetDataMeta(dm *DataMeta) {
	bdf.dataFeedMeta = dm
}

func (bdf *BaseDataFeed) GetMeta() *DataMeta {
	return bdf.dataFeedMeta
}
