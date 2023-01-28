package indicator

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"sort"
)

type MarketIndicator interface {
	Name() string
	GetCurrentValue(int64) float64
	AddData(ts int64, node any)
	GetAllSortedData() []any
	DoPlot(page *charts.Kline)
}

type GlobalIndicator interface {
	AddData(ctx *GlobalMsg)
	DoPlot(page *charts.Page)
}

type Line struct {
	*BaseLine
	Name string
}

type LineNode struct {
	Value     float64
	TimeStamp int64
}

func NewLine(t LineType, name string) *Line {
	return &Line{
		BaseLine: NewBaseLine(t),
		Name:     name,
	}
}

func (k *Line) GetByTs(ts int64) (*LineNode, error) {
	res, err := k.BaseLine.GetByTs(ts)
	if err != nil {
		return nil, err
	}
	vv := res.(*LineNode)
	return vv, nil
}

func (k *Line) GetByTsAndCount(ts int64, count int64) ([]*LineNode, error) {
	res, err := k.BaseLine.GetByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]*LineNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*LineNode)
	}

	return newRes, nil
}

func (k *Line) AddData(ts int64, value float64) {
	node := &LineNode{
		TimeStamp: ts, Value: value,
	}
	k.BaseLine.AddData(ts, node)
}

func (k *Line) GetAllSortedData() []*LineNode {
	res := k.BaseLine.GetAllSortedData()
	newRes := make([]*LineNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*LineNode)
	}

	sort.Slice(newRes, func(i, j int) bool {
		return newRes[i].TimeStamp < newRes[j].TimeStamp
	})
	return newRes
}

type LineType int64

const LineType_Day = 1
const LineType_Minite = 2
const LineType_Hour = 3

type KNode struct {
	High      float64
	Low       float64
	Open      float64
	Close     float64
	Date      string
	TimeStamp int64
}

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
