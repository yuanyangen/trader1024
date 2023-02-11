package data_feed

import "github.com/yuanyangen/trader1024/engine/event"

type DataType int64

const DataTypeKLine DataType = 1

type Data struct {
	DataType DataType
	KData    *KNode
}

type KNode struct {
	High      float64
	Low       float64
	Open      float64
	Close     float64
	Date      string
	TimeStamp int64
}

type SourceType int64

const SourceType_CSV = 1
const SourceType_Live = 2

type DataMeta struct {
}

type DataFeed interface {
	RegisterChan(out chan *Data)
	SetEventTrigger(event.EventTrigger)
}
