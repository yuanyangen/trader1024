package main

import (
	"github.com/yuanyangen/trader1024/data/crawler"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/data/storage"
	"github.com/yuanyangen/trader1024/engine/model"
)

func main() {
	crawlMinuteData()

}

func crawlDailyData() {
	eastMoneyCrawler := &crawler.EastMoney{}
	marketIds := markets.GetAllFutureMarketIds()
	for _, marketId := range marketIds {
		t := model.LineType_Day
		allNodes, err := eastMoneyCrawler.CrawlDaily(marketId, "19900101", "20230307")
		if err != nil {
			panic("fadsfa")
		}

		csvStorage := storage.NewCsvStorage()
		csvStorage.SaveData(marketId, t, allNodes)
	}
}

func crawlMinuteData() {
	eastMoneyCrawler := &crawler.EastMoney{}
	marketIds := markets.GetAllFutureMarketIds()
	for _, marketId := range marketIds {
		t := model.LineType_5Minite

		allNodes, err := eastMoneyCrawler.CrawlMinute5(marketId)
		if err != nil {
			panic("fadsfa")
		}

		csvStorage := storage.NewCsvStorage()
		csvStorage.SaveData(marketId, t, allNodes)
	}
}
