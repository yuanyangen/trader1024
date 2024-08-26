package main

import (
	"context"
	"github.com/yuanyangen/trader1024/data/crawler/common"
	"github.com/yuanyangen/trader1024/data/crawler/plugins/sina"
	"github.com/yuanyangen/trader1024/data/datasource"
	"testing"
)

func TestCrawl(t *testing.T) {
	common.CrawlAllSubject(context.Background())
}

func TestCrawlContractDaily(t *testing.T) {
	common.CrawlAllSubjectMainDailyData(context.Background(), &sina.Sina{})
}
func TestCrawlContractInfo(t *testing.T) {
	contract, err := datasource.GetContractByCnName(context.Background(), "玉米", "202407")
	if err != nil {
		panic(err)
	}
	common.CrawlContractByContract(context.Background(), contract, &sina.Sina{})
}
