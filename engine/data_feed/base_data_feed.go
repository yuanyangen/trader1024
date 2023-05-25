package data_feed

import (
	"github.com/yuanyangen/trader1024/engine/model"
)

type BaseDataFeed struct {
	Source           model.SourceType
	outerChannels    []chan *model.Data
	eventTriggerChan chan *model.EventMsg
	et               model.EventTrigger
}

func (bdf *BaseDataFeed) RegisterChan(out chan *model.Data) {
	if bdf.outerChannels == nil {
		bdf.outerChannels = []chan *model.Data{out}
	} else {
		bdf.outerChannels = append(bdf.outerChannels, out)
	}
}

func (bdf *BaseDataFeed) SendData(out *model.Data) {
	for _, v := range bdf.outerChannels {
		v <- out
	}
}

func (bdf *BaseDataFeed) SetEventTrigger(et model.EventTrigger) {
	bdf.et = et
	et.RegisterEventReceiver(bdf.eventTriggerChan)
}
