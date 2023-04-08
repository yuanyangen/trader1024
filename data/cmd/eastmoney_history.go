package main

import (
	"fmt"
	"github.com/yuanyangen/trader1024/data/plugins"
	"github.com/yuanyangen/trader1024/data/storage"
	"github.com/yuanyangen/trader1024/engine/model"
)

func crawlEastMoneyHostoryFutureData() {
	crawlHistoryDailyData()
}

func crawlHistoryDailyData() {
	eastMoneyCrawler := &plugins.EastMoney{}
	//marketIds := markets.GetAllFutureMarketIds()
	csvStorage := storage.EastMoneyHttpStorage()
	allMarkets := eastMoneyCrawler.CrawlAllMarket()
	for _, v := range allMarkets {
		t := model.LineType_Day
		v.Count = int64(len(allMarkets))
		allNodes, err := eastMoneyCrawler.CrawlDaily(v.SecId, "19900101", "20500101")
		//fmt.Println(v, len(allNodes))
		fmt.Printf("\"%v\":    {Type: model.MarKetType_FUTURE, MarketId: \"%v\", Code: \"%v\", Name: \"%v\", SecId: \"%v\", Exchange: \"%v\", Count:%v},\n",
			v.Code, v.MarketId, v.Code, v.Name, v.SecId, v.Exchange, len(allNodes))

		if err != nil {
			panic("fadsfa")
		}
		csvStorage.SaveData(v.MarketId, t, allNodes)
	}
}
