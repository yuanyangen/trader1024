package main

import (
	"github.com/yuanyangen/trader1024/data/crawler/common"
	"github.com/yuanyangen/trader1024/data/crawler/plugins/sina"
	"time"
)

func main() {
	for {
		now := time.Now().Unix()
		idx := now / 86400
		_, ok := crawledDate[idx]
		if !ok {
			oneCrawlLoop()
		}
		time.Sleep(time.Minute)
	}
}

// 每天抓一次数据,这里暂时不用考虑时区的问题
var crawledDate = map[int64]bool{}

func oneCrawlLoop() {
	common.CrawlMarketMainDailyData("橡胶", &sina.Sina{})
	//common.CrawlMainDailyData(&sina.Sina{})
	//common.CrawlMainDailyData(&eastmoney.EastMoney{})
}
