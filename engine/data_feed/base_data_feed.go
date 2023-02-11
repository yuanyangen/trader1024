package data_feed

import (
	"github.com/yuanyangen/trader1024/engine/event"
)

type BaseDataFeed struct {
	Source           SourceType
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
	for _, v := range bdf.outerChannels {
		v <- out
	}
}

func (bdf *BaseDataFeed) SetEventTrigger(et event.EventTrigger) {
	bdf.et = et
	et.RegisterEventReceiver(bdf.eventTriggerChan)
}
