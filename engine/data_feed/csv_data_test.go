package data_feed

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"testing"
	"time"
)

func TestCsvKLineDataFeed(t *testing.T) {
	df := NewCsvKLineDataFeed("/home/yuanyangen/HomeData/go/trader1024/datas/datas/daily/IC主力合约.csv")

	outChannel := make(chan *model.Data, 10)
	df.RegisterChan(outChannel)
	df.StartFeed()

	go func() {
		for v := range outChannel {
			if v == nil {

			}
			//fmt.Println(*v)
		}
	}()
	time.Sleep(time.Second * 1)
}
