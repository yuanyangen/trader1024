package common

import (
	"context"
	"fmt"
	"github.com/yuanyangen/trader1024/dal/mongo"
	"github.com/yuanyangen/trader1024/data/crawler/plugins/sina"
	"github.com/yuanyangen/trader1024/data/datasource"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
	"time"
)

type Crawler interface {
	CrawlContractDaily(ctx context.Context, market *model.Contract, startDate, endDate time.Time) ([]*model.KLineNode, error)
	CrawlContractMinute(ctx context.Context, market *model.Contract, startDate, endDate time.Time) ([]*model.KLineNode, error)
	CrawlContractWeekly(ctx context.Context, market *model.Contract, startDate, endDate time.Time) ([]*model.KLineNode, error)
}

func CrawlAllSubject(ctx context.Context) {
	subjects, err := mongo.QueryAllSubject(ctx)
	if err != nil {
		logs.Info("query all subject error ", err)
		return
	}
	for _, v := range subjects {
		CrawlAllMainContractBySubject(ctx, v, &sina.Sina{})
	}
}
func CrawlAllMainContractBySubjectName(ctx context.Context, subjectName string) {
	subject, err := mongo.QuerySubject(ctx, subjectName)
	if err != nil {
		logs.Info("query all subject error ", err)
		return
	}
	CrawlAllMainContractBySubject(ctx, subject, &sina.Sina{})
}

func CrawlAllMainContractBySubject(ctx context.Context, subject *model.SubjectDO, handler Crawler) {
	logs.Info("start to crawl crawlAllMainContractBySubjectName %v", subject.CNName)
	if subject == nil {
		panic("should not reach here")
	}
	if subject.OfflineTime != "" {
		return
	}
	// 从2010.1.1 开始计算数据。
	for t := time.Unix(1262275200, 0); t.Unix() < time.Now().Add(time.Hour*24*365).Unix(); t = t.Add(30 * 24 * time.Hour) {
		month := int(t.Month())
		for _, v := range subject.ContractMonth {
			contractDate := t.Format("200601")
			if v == month && contractDate > subject.FirstContract {
				if contractDate < "202001" {
					continue
				}
				logs.Info("start to crawl %v %v", subject.CNName, contractDate)
				contractDo, err := mongo.QueryContract(ctx, subject.CNName, contractDate)
				if err != nil {
					logs.Info("QueryContract error " + err.Error())
					continue
				}
				if contractDo == nil {
					contractDo = &model.ContractDO{
						ContractCnName: subject.CNName,
						ContractDate:   contractDate,
					}
				}
				contract := &model.Contract{SubjectDO: subject, ContractDO: contractDo}
				err = CrawlContractByContract(ctx, contract, handler)
				if err != nil {
					logs.Info("SaveContractInfo error " + err.Error())
				}
			}
		}
	}
	logs.Info("finish to crawl %v", subject.CNName)
}

func CrawlContractByContract(ctx context.Context, contract *model.Contract, handler Crawler) error {
	logs.Info("start to crawl crawlAllMainContractBySubjectName %v", contract.CNName)
	if contract == nil {
		panic("should not reach here")
	}
	if contract.OfflineTime != "" {
		return fmt.Errorf("already offline subject")
	}
	// 从2010.1.1 开始计算数据。
	//if contractDo.LastCrawlTime > time.Now().Add(-24*3*time.Hour).Unix() {
	if contract.ContractEndTime != 0 || contract.LastCrawlTime > time.Now().Add(-24*7*time.Hour).Unix() {
		logs.Info("finish: no crawl already crawled %v %v", contract.CNName, contract.ContractDate)
		//return fmt.Errorf("finish: no crawl already crawled %v %v", contract.CNName, contract.ContractDate)
	}

	dataNodes, err := handler.CrawlContractDaily(ctx, contract, time.Unix(0, 0), time.Now())
	if err != nil {
		logs.Info("finish: crawl daily data_crawler failed %v", err)
		return err
	}
	if len(dataNodes) == 0 {
		logs.Info("finish: crawl no data_crawler %v %v", contract.CNName, contract.ContractDate)
		return fmt.Errorf("finish: crawl no data_crawler %v %v", contract.CNName, contract.ContractDate)
	}
	contract.ContractStartTime = dataNodes[0].TimeStamp
	lastDataNode := dataNodes[len(dataNodes)-1]
	if lastDataNode.TimeStamp+7*86400 < time.Now().Unix() {
		contract.ContractEndTime = lastDataNode.TimeStamp
	}
	contract.LastCrawlTime = time.Now().Unix()

	for _, node := range dataNodes {
		err = mongo.SaveDataNode(ctx, node)
		if err != nil {
			panic("err " + err.Error())
		}
	}
	err = mongo.SaveContractInfo(ctx, contract.ContractDO)
	if err != nil {
		logs.Info("SaveContractInfo error " + err.Error())
	}
	logs.Info("finish: success crawl %v %v", contract.CNName, contract.ContractDate)

	logs.Info("finish to crawl %v", contract.CNName)
	return nil
}

func CrawlAllSubjectMainDailyData(ctx context.Context, handler Crawler) {
	subjects, err := mongo.QueryAllSubject(ctx)
	if err != nil {
		logs.Info("query all subject error " + err.Error())
		return
	}
	for _, v := range subjects {
		CrawlSubjectDailyData(ctx, v.CNName, handler)
	}
}

func CrawlSubjectDailyData(ctx context.Context, subjectName string, handler Crawler) {
	allContracts := datasource.GetAllContractFromDb(ctx, subjectName)
	for _, v := range allContracts {
		allNodes, err := handler.CrawlContractDaily(ctx, v, time.Unix(0, 0), time.Now())
		if err != nil {
			logs.Info("crawl data_crawler error %v", err)
			continue
		}
		for _, node := range allNodes {
			err = mongo.SaveDataNode(ctx, node)
			if err != nil {
				panic("err " + err.Error())
			}
		}
		logs.Info("finished %v %v\n", v.Id(), err)
	}
}

//
//func crawlHistoryMinuteData(handler Crawler) {
//	csvStorage := handler.StorageClient()
//	allContracts := handler.CrawlAllMainMarket()
//	for _, contract := range allContracts {
//		t := model.LineType_Minite
//		allNodes, err := handler.CrawlMinute(contract, time.Unix(0, 0), time.Now())
//		if err != nil {
//			panic("fadsfa")
//		}
//
//		csvStorage.SaveData(contract.Id(), t, allNodes)
//	}
//}
//func crawlHistoryWeekData(handler Crawler) {
//	allContracts := handler.CrawlAllMainMarket()
//	csvStorage := handler.StorageClient()
//	for _, contract := range allContracts {
//		t := model.LineType_Week
//		allNodes, err := handler.CrawlMinute(contract, time.Unix(0, 0), time.Now())
//		if err != nil {
//			panic("fadsfa")
//		}
//
//		csvStorage.SaveData(contract.Id(), t, allNodes)
//	}
//}
