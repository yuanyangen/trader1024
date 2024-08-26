package main

import (
	"context"
	"github.com/yuanyangen/trader1024/data/crawler/common"
	"time"
)

func asyncStartDataCrawler() {
	go func() {
		startDataCrawler()
	}()
}

func startDataCrawler() {
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
	ctx := context.Background()
	//common.CrawlAllSubject(ctx)
	//common.CrawlAllMainContractBySubjectName(ctx, "橡胶")
	common.CrawlAllMainContractBySubjectName(ctx, "玉米")
}
