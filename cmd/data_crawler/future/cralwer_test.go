package main

import (
	"context"
	"testing"

	"github.com/yuanyangen/trader1024/data/crawler/future"
	"github.com/yuanyangen/trader1024/data/crawler/plugins/sina"
	"github.com/yuanyangen/trader1024/data/datasource"
)

func TestCrawl(t *testing.T) {
	future.CrawlAllSubject(context.Background())
}

func TestCrawlContractDaily(t *testing.T) {
	future.CrawlAllSubjectMainDailyData(context.Background(), &sina.Sina{})
}
func TestCrawlContractInfo(t *testing.T) {
	contract, err := datasource.GetContractByCnName(context.Background(), "玉米", "202407")
	if err != nil {
		panic(err)
	}
	future.CrawlContractByContract(context.Background(), contract, &sina.Sina{})
}
