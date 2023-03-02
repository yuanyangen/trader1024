package data_feed

import (
	"github.com/yuanyangen/trader1024/data/storage"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CsvKLineDataFeed struct {
	*BaseDataFeed
	marketId string
	storage  *storage.KVStorage
}

func NewCsvKLineDataFeed(marketId string) *CsvKLineDataFeed {
	cdf := &CsvKLineDataFeed{
		BaseDataFeed: &BaseDataFeed{
			eventTriggerChan: make(chan *event.EventMsg, 1024),
			Source:           SourceType_CSV,
		},
		marketId: marketId,
		storage:  storage.EastMoneyStorage(),
	}
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
			data := ckdf.storage.GetDataByTs(ckdf.marketId, model.LineType_Day, ts)
			if data != nil {
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
