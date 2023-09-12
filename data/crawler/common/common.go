package common

import (
	"fmt"
	"github.com/yuanyangen/trader1024/data/storage_client"
	"github.com/yuanyangen/trader1024/engine/model"
	"time"
)

type Crawler interface {
	StorageClient() *storage_client.HttpStorageClient
	CrawlAllMainMarket() []*model.Contract
	CrawlAllAvailableMainMarket() []*model.Contract
	CrawlDaily(market *model.Contract, startDate, endDate time.Time) ([]*model.KNode, error)
	CrawlMinute(market *model.Contract, startDate, endDate time.Time) ([]*model.KNode, error)
	CrawlWeekly(market *model.Contract, startDate, endDate time.Time) ([]*model.KNode, error)
}

func CrawlMainDailyData(handler Crawler) {
	csvStorage := handler.StorageClient()
	allMarkets := handler.CrawlAllMainMarket()
	for _, v := range allMarkets {
		t := model.LineType_Day
		allNodes, err := handler.CrawlDaily(v, time.Unix(0, 0), time.Now())
		if err != nil {
			panic("fadsfa")
		}
		csvStorage.SaveData(v.ContractId, t, allNodes)
		fmt.Println("finished " + v.ContractId)

	}
}

func crawlHistoryMinuteData(handler Crawler) {
	csvStorage := handler.StorageClient()
	allMarkets := handler.CrawlAllMainMarket()
	for _, market := range allMarkets {
		t := model.LineType_Minite
		allNodes, err := handler.CrawlMinute(market, time.Unix(0, 0), time.Now())
		if err != nil {
			panic("fadsfa")
		}

		csvStorage.SaveData(market.ContractId, t, allNodes)
	}
}
func crawlHistoryWeekData(handler Crawler) {
	allMarkets := handler.CrawlAllMainMarket()
	csvStorage := handler.StorageClient()
	for _, market := range allMarkets {
		t := model.LineType_Week
		allNodes, err := handler.CrawlMinute(market, time.Unix(0, 0), time.Now())
		if err != nil {
			panic("fadsfa")
		}

		csvStorage.SaveData(market.ContractId, t, allNodes)
	}
}
