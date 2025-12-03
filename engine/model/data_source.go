package model

import "context"

type DataType int64

const DataTypeKLine DataType = 1

type Data struct {
	DataType DataType
	Data     *KLineNode
}

type SourceType int64

const SourceType_CSV = 1
const SourceType_Live = 2

type DataMeta struct {
}

type DataFeed interface {
	RegisterChan(out chan *Data)
}

type DateSource interface {
	GetDataByTs(ctx context.Context, uniqueCode string, lineType LineType, ts int64) *KLineNode
	SaveDataByTs(ctx context.Context, data *KLineNode) error
	GetAllContractBySubjectName(ctx context.Context, subjectName string) []*TradeObject
}
