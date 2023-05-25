package data_feed

import (
	"github.com/yuanyangen/trader1024/data/storage"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CsvKLineDataFeed struct {
	*BaseDataFeed
	marketId string
	storage  *storage.HttpStorage
}

func NewCsvKLineDataFeed(marketId string) *CsvKLineDataFeed {
	cdf := &CsvKLineDataFeed{
		BaseDataFeed: &BaseDataFeed{
			eventTriggerChan: make(chan *model.EventMsg, 1024),
			Source:           model.SourceType_CSV,
		},
		marketId: marketId,
		storage:  storage.EastMoneyHttpStorage(),
	}
	cdf.startEventReceiver()

	return cdf
}

func (ckdf *CsvKLineDataFeed) startEventReceiver() chan *model.Data {
	respChan := make(chan *model.Data, 1024)
	if ckdf.eventTriggerChan == nil {
		panic("event chan error")
	}
	utils.AsyncRun(func() {
		for event := range ckdf.eventTriggerChan {
			ts := event.TimeStamp
			data := ckdf.storage.GetDataByTs(ckdf.marketId, model.LineType_Day, ts)
			if data != nil {
				node := &model.Data{
					DataType: model.DataTypeKLine,
					KData:    data,
				}
				ckdf.SendData(node)
			}
		}
	})
	return respChan
}
