package model

type EventMsg struct {
	Type      int
	TimeStamp int64
}

type EventTrigger interface {
	Start()
	RegisterEventReceiver(chan *EventMsg)
}
