package data_feed

import (
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/engine/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

type CsvKLineDataFeed struct {
	*BaseDataFeed
	fileAddr string
	option   *CsvDataFeedOption
	csvData  map[int64]*KNode
}

type CsvFieldIndex string

const CsvFieldIndex_Date CsvFieldIndex = "Date"
const CsvFieldIndex_Open CsvFieldIndex = "Open"
const CsvFieldIndex_Close CsvFieldIndex = "Close"
const CsvFieldIndex_High CsvFieldIndex = "High"
const CsvFieldIndex_Low CsvFieldIndex = "Low"

var DefaultCsvFieldIndex = []CsvFieldIndex{
	CsvFieldIndex_Date,
	CsvFieldIndex_Open,
	CsvFieldIndex_Close,
	CsvFieldIndex_High,
	CsvFieldIndex_Low,
}
var csvFieldHandler = map[CsvFieldIndex]func(*CsvKLineDataFeed, *KNode, string){
	CsvFieldIndex_Date: func(ckdf *CsvKLineDataFeed, node *KNode, value string) {
		t, err := time.Parse(ckdf.option.DateFormat, value)
		if err != nil {
			panic("time format error")
		}
		node.Date = value
		node.TimeStamp = t.Unix()
	},
	CsvFieldIndex_Open: func(feed *CsvKLineDataFeed, node *KNode, s string) {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic("open data error" + s)
		}
		node.Open = v
	},
	CsvFieldIndex_Close: func(feed *CsvKLineDataFeed, node *KNode, s string) {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic("close data error" + s)
		}
		node.Close = v
	},
	CsvFieldIndex_High: func(feed *CsvKLineDataFeed, node *KNode, s string) {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic("high data error" + s)
		}
		node.High = v
	},
	CsvFieldIndex_Low: func(feed *CsvKLineDataFeed, node *KNode, s string) {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic("low data error" + s)
		}
		node.Low = v
	},
}

type Option func(*CsvDataFeedOption)

type CsvDataFeedOption struct {
	index []CsvFieldIndex

	splitter   string
	DateFormat string
}

func WithCsvDataIndex(index []CsvFieldIndex) func(*CsvDataFeedOption) {
	return func(option *CsvDataFeedOption) {
		option.index = index
	}
}

func NewCsvKLineDataFeed(fileName string, options ...Option) *CsvKLineDataFeed {
	option := &CsvDataFeedOption{
		index:      DefaultCsvFieldIndex,
		splitter:   ",",
		DateFormat: "2006-01-02",
	}
	for _, o := range options {
		o(option)
	}

	cdf := &CsvKLineDataFeed{

		BaseDataFeed: &BaseDataFeed{
			eventTriggerChan: make(chan *event.EventMsg, 1024),
			Source:           SourceType_CSV,
		},
		fileAddr: fileName,
		option:   option,
		csvData:  map[int64]*KNode{},
	}
	cdf.loadData()
	cdf.startEventReceiver()

	return cdf
}

func (ckdf *CsvKLineDataFeed) startEventReceiver() chan *Data {
	respChan := make(chan *Data, 1024)
	if ckdf.eventTriggerChan == nil {
		panic("event chan error")
	}
	utils.AsyncRun(func() {
		for event := range ckdf.eventTriggerChan {
			ts := event.TimeStamp
			data, ok := ckdf.csvData[utils.UnityDailyTimeStamp(ts)]
			if ok {
				node := &Data{
					DataType: DataTypeKLine,
					KData:    data,
				}
				ckdf.SendData(node)
			}
		}
	})
	return respChan
}

func (ckdf *CsvKLineDataFeed) loadData() {
	content, err := os.ReadFile(ckdf.fileAddr)
	if err != nil {
		panic(err)
		//return
	}
	tmp := strings.Split(string(content), "\n")
	for i := 1; i < len(tmp); i++ {
		if tmp[i] == "" {
			continue
		}
		tmp2 := strings.Split(tmp[i], ckdf.option.splitter)
		if len(tmp2) != len(ckdf.option.index) {
			panic("csv data type error")
		}

		node := &KNode{}
		for i, v := range tmp2 {
			v := strings.TrimSpace(v)
			fieldName := ckdf.option.index[i]
			handler := csvFieldHandler[fieldName]
			handler(ckdf, node, v)
		}
		ckdf.csvData[utils.UnityDailyTimeStamp(node.TimeStamp)] = node
	}
}
