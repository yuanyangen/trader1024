package model

import (
	"fmt"
)

const DataTypeKLine DataType = 1

type Data struct {
	DataType DataType
	KData    *KNode
	DataMeta *DataMeta
}

type SourceType int64

const SourceType_CSV = 1
const SourceType_Live = 2

type DataType int64

const DataType_STOCK = 1
const DataType_FUTURE = 2

type DataMeta struct {
	Name   string
	Type   DataType
	Source SourceType
}

func (dm *DataMeta) String() string {
	return fmt.Sprintf("%v_%v_%v", dm.Name, dm.Type, dm.Source)
}

func NewDataMeta(name string, ty DataType, source SourceType) *DataMeta {
	return &DataMeta{
		Name:   name,
		Type:   ty,
		Source: source,
	}
}

type DataFeed interface {
	GetMeta() *DataMeta
	StartFeed() chan *Data // 每次data feed生成的数据, 比如一个时间点的k线中的一个点
	RegisterChan(out chan *Data)
}
