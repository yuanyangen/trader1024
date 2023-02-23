package data_feed

import (
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/engine/model"
)

type DataType int64

const DataTypeKLine DataType = 1

type Data struct {
	DataType DataType
	KData    *model.KNode
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
