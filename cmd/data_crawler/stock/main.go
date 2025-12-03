package main

import (
	"context"

	"github.com/yuanyangen/trader1024/dal/mongo"
	"github.com/yuanyangen/trader1024/data/crawler/stock"
)

func main() {
	// saveSubjectToMongo()
	// return
	ctx := context.Background()
	//common.CrawlAllSubject(ctx)
	//common.CrawlAllMainContractBySubjectName(ctx, "橡胶")
	//common.CrawlAllMainContractBySubjectName(ctx, "玉米")
	stock.CrawlBySubjectName(ctx, "比亚迪", &stock.Sina{})
}

func saveSubjectToMongo() {
	for _, v := range AllSubjects {
		mongo.SaveSubjectInfo(context.Background(), v)
	}
}
