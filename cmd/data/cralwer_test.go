package main

import (
	"github.com/yuanyangen/trader1024/data/crawler/common"
	"github.com/yuanyangen/trader1024/data/crawler/plugins/sina"
	"testing"
)

func TestCrawl(t *testing.T) {
	common.CrawlMarketMainDailyData("橡胶", &sina.Sina{})
}
