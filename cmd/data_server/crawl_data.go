package main

import (
    "github.com/yuanyangen/trader1024/data/crawler"
    "github.com/yuanyangen/trader1024/data/markets"
    "github.com/yuanyangen/trader1024/data/storage"
    "github.com/yuanyangen/trader1024/engine/model"
    "time"
)

var crawledDate = map[int64]bool{}

// 每天抓一次数据,这里暂时不用考虑时区的问题
func startHistoryCrawlDataServer() {
    for {
        now := time.Now().Unix()
        idx := now/86400
        _, ok := crawledDate[idx]
        if !ok {
            oneCrawlLoop()
        }
        time.Sleep(time.Minute)
    }
}


func oneCrawlLoop() {
    crawlEastMoneyAllFutureData()
}


func crawlEastMoneyAllFutureData() {
    crawlDailyData()
    crawlMinuteData()
}

func crawlDailyData() {
	eastMoneyCrawler := &crawler.EastMoney{}
	marketIds := markets.GetAllFutureMarketIds()
    csvStorage := storage.EastMoneyStorage()
	for _, marketId := range marketIds {
		t := model.LineType_Day
        startDate := time.Now().Add(time.Hour * -24 * 2).Format("20060102")
        endDate := time.Now().Add(time.Hour * 24).Format("20060102")
        allNodes, err := eastMoneyCrawler.CrawlDaily(marketId, startDate, endDate)
		if err != nil {
			panic("fadsfa")
		}
		csvStorage.SaveData(marketId, t, allNodes)
	}
}

func crawlMinuteData() {
	eastMoneyCrawler := &crawler.EastMoney{}
	marketIds := markets.GetAllFutureMarketIds()
    csvStorage := storage.EastMoneyStorage()

	for _, marketId := range marketIds {
		t := model.LineType_Minite

		allNodes, err := eastMoneyCrawler.CrawlMinute(marketId)
		if err != nil {
			panic("fadsfa")
		}

		csvStorage.SaveData(marketId, t, allNodes)
	}
}

