package event

import (
	"github.com/yuanyangen/trader1024/engine/utils"
	"time"
)

type EventTrigger interface {
	Start()
	RegisterEventReceiver(chan *EventMsg)
}

type baseEventTrigger struct {
	eventReceiver []chan *EventMsg
}

func (bet *baseEventTrigger) RegisterEventReceiver(ch chan *EventMsg) {
	bet.eventReceiver = append(bet.eventReceiver, ch)
}
func (bet *baseEventTrigger) sendEvent(msg *EventMsg) {
	for _, ch := range bet.eventReceiver {
		ch <- msg
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
			msg := &EventMsg{
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

func NewBackTestDailyEventTrigger(startTs, endTs int64) EventTrigger {
	if startTs >= endTs {
		panic("ts error")
	}
	return &BackTestDailyEventTrigger{
		StartTimeStamp:   startTs,
		EndTimeStamp:     endTs,
		baseEventTrigger: &baseEventTrigger{},
	}
}

func (det *BackTestDailyEventTrigger) Start() {
	utils.AsyncRun(func() {
		for ts := det.StartTimeStamp; ts <= det.EndTimeStamp; {
			msg := &EventMsg{
				TimeStamp: utils.UnityDailyTimeStamp(ts),
			}
			det.baseEventTrigger.sendEvent(msg)
			ts += 86400
		}
	})
}
