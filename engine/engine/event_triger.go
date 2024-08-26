package engine

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"time"
)

type baseEventTrigger struct {
	eventReceiver []model.EventReceiver
}

func (bet *baseEventTrigger) RegisterEventReceiver(ch model.EventReceiver) {
	bet.eventReceiver = append(bet.eventReceiver, ch)
}
func (bet *baseEventTrigger) sendEvent(msg *model.EventMsg) {
	for _, ch := range bet.eventReceiver {
		ch.DealEvent(msg)
	}
}

type DailyEventTrigger struct {
	*baseEventTrigger
}

func NewDailyEventTrigger() *DailyEventTrigger {
	return &DailyEventTrigger{
		baseEventTrigger: &baseEventTrigger{},
	}
}

func (det *DailyEventTrigger) Start() {
	utils.AsyncRun(func() {
		for {
			msg := &model.EventMsg{
				TimeStamp: time.Now().Unix(),
			}
			det.baseEventTrigger.sendEvent(msg)
			time.Sleep(86400)
		}
	})
}

type BackTestDailyEventTrigger struct {
	StartTimeStamp int64
	EndTimeStamp   int64
	*baseEventTrigger
}

func NewBackTestDailyEventTrigger(startDate, endDate, format string) model.EventTrigger {
	if startDate >= endDate {
		panic("ts error")
	}
	startTs := utils.StrToTs(startDate, format)
	endTs := utils.StrToTs(endDate, format)
	return &BackTestDailyEventTrigger{
		StartTimeStamp:   startTs,
		EndTimeStamp:     endTs,
		baseEventTrigger: &baseEventTrigger{},
	}
}

func (det *BackTestDailyEventTrigger) Start() {
	utils.AsyncRun(func() {
		for ts := det.StartTimeStamp; ts <= det.EndTimeStamp; {
			msg := &model.EventMsg{
				TimeStamp: utils.UnityDailyTimeStamp(ts),
			}
			det.baseEventTrigger.sendEvent(msg)
			ts += 86400
		}
	})
}
