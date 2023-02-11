package data_feed

import (
	"testing"
	"time"
)

func TestCsvKLineDataFeed(t *testing.T) {
	df := NewCsvKLineDataFeed("/home/yuanyangen/HomeData/go/trader1024/data/data/daily/IC主力合约.csv")

	outChannel := make(chan *Data, 10)
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
