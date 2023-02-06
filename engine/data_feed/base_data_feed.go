package data_feed

import (
	"github.com/yuanyangen/trader1024/engine/event"
)

type BaseDataFeed struct {
	DataFeedMeta     *DataMeta
	outerChannels    []chan *Data
	eventTriggerChan chan *event.EventMsg
	et               event.EventTrigger
}

func (bdf *BaseDataFeed) RegisterChan(out chan *Data) {
	if bdf.outerChannels == nil {
		bdf.outerChannels = []chan *Data{out}
	} else {
		bdf.outerChannels = append(bdf.outerChannels, out)
	}
}

func (bdf *BaseDataFeed) SendData(out *Data) {
	out.DataMeta = bdf.DataFeedMeta
	for _, v := range bdf.outerChannels {
		v <- out
	}
}

func (bdf *BaseDataFeed) SetDataMeta(dm *DataMeta) {
	bdf.DataFeedMeta = dm
}
func (bdf *BaseDataFeed) SetEventTrigger(et event.EventTrigger) {
	bdf.et = et
	et.RegisterEventReceiver(bdf.eventTriggerChan)
}

func (bdf *BaseDataFeed) GetMeta() *DataMeta {
	return bdf.DataFeedMeta
}
