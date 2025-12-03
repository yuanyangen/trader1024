package future

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/yuanyangen/trader1024/data/crawler/http"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
)

var allVendorPrefix = map[string]string{
	"不锈钢": "SS",
	"橡胶":  "RU",
	"沥青":  "BU",
	"沪金":  "AU",
	"沪铅":  "PB",
	"沪铜":  "CU",
	"沪铝":  "AL",
	"沪银":  "AG",
	"沪锌":  "ZN",
	"沪锡":  "SN",
	"沪镍":  "NI",
	"热卷":  "HC",
	"燃油":  "FU",
	"纸浆":  "SP",
	"线材":  "WR",
	"螺纹钢": "RB",

	"IC合约": "IC",
	"IF合约": "IF",
	"IH合约": "IH",
	"IM合约": "IM",
	"TF合约": "TF",
	"TS合约": "TS",
	"T合约":  "T",

	"LPG": "PG",
	"PVC": "V",
	"乙二醇": "EG",
	"塑料":  "L",
	"棕榈油": "P",
	"淀粉":  "CS",
	"焦炭":  "J",
	"焦煤":  "JM",
	"玉米":  "C",
	"生猪":  "LH",
	"粳米":  "RR",
	"纤维板": "FB",
	"聚丙烯": "PP",
	"胶合板": "BB",
	"苯乙烯": "EB",
	"豆一":  "A",
	"豆二":  "B",
	"豆油":  "Y",
	"豆粕":  "M",
	"铁矿石": "I",
	"鸡蛋":  "JD",

	"PTA": "TA",
	"动力煤": "ZC",
	"尿素":  "UR",
	"强麦":  "WH",
	"晚籼稻": "LR",
	"普麦":  "", // 没了
	"棉纱":  "CY",
	"棉花":  "CF",
	"玻璃":  "FG",
	"甲醇":  "MA",
	"白糖":  "SR",
	"短纤":  "PF",
	"硅铁":  "SF",
	"粳稻":  "JR",
	"红枣":  "CJ",
	"纯碱":  "SA",
	"花生":  "PK",
	"苹果":  "AP",
	"菜油":  "OI",
	"菜籽":  "RS",
	"菜粕":  "RM",
	"锰硅":  "SM",
}

type Sina struct {
}

//func (em *Sina) CrawlAllMainMarket(ctx context.Context) []*model.Contract {
//	allSubject, _ := markets.GetAllFutureSubjects(ctx)
//	allMarkets := make([]*model.Contract, len(allSubject))
//	for i, v := range allSubject {
//		allMarkets[i] = &model.Contract{
//			SubjectDO: v,
//		}
//	}
//	sort.Slice(allMarkets, func(i, j int) bool {
//		return allMarkets[i].Exchange+allMarkets[i].Id() < allMarkets[j].Exchange+allMarkets[j].Id()
//	})
//	return allMarkets
//}

// date
func buildVendorIdByDate(vendorIdPrefix, date string) string {
	if vendorIdPrefix == "" {
		return ""
	}
	if date == model.MainContinuous {
		return vendorIdPrefix + "0"
	}
	t, err := time.Parse("200601", date)
	if err != nil {
		panic(err)
	}
	return vendorIdPrefix + t.Format("0601")

}

func (em *Sina) CrawlContractMinute(ctx context.Context, market *model.TradeObject, startDate, endDate time.Time) ([]*model.KLineNode, error) {
	return nil, nil
}
func (em *Sina) CrawlContractWeekly(ctx context.Context, market *model.TradeObject, startDate, endDate time.Time) ([]*model.KLineNode, error) {
	return nil, nil
}

func (em *Sina) CrawlContractDaily(ctx context.Context, market *model.TradeObject, startDate, endDate time.Time) ([]*model.KLineNode, error) {
	time.Sleep(time.Second * 2)
	return em.crawlDailyOld(ctx, market, startDate, endDate)
}

//func (em *Sina) crawlDailyNew(ctx context.Context, market *model.Contract, startDate, endDate time.Time) ([]*model.KLineNode, error) {
//	contractId := buildVendorIdByDate(allVendorPrefix[market.CNName], market.ContractDate)
//	body, err := http.Get(ctx, "https://stock2.finance.sina.com.cn/futures/api/json.php/IndexService.getInnerFuturesDailyKLine", map[string]string{
//		"symbol": contractId,
//	}, map[string]string{
//		"Authority":          "stock2.finance.sina.com.cn",
//		"Accept":             "*/*",
//		"Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
//		"Referer":            "https://finance.sina.com.cn/futures/quotes/V0.shtml",
//		"Sec-Ch-Ua":          "\"Not_A Brand\";v=\"99\", \"Microsoft Edge\";v=\"109\", \"Chromium\";v=\"109\"",
//		"Sec-Ch-Ua-Mobile":   "?0",
//		"Sec-Ch-Ua-Platform": "\"Linux\"",
//		"Sec-Fetch-Dest":     "script",
//		"Sec-Fetch-Mode":     "no-cors",
//		"Sec-Fetch-Site":     "same-site",
//		"User-Agent":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.49",
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	allKnodes := [][]string{}
//	err = sonic.UnmarshalString(body, &allKnodes)
//	if err != nil {
//		logs.Info("sina response error")
//		return nil, fmt.Errorf("sina resp data_crawler error")
//	}
//	return convertSinaKnodeToKnode(ctx, market, allKnodes), nil
//}

func (em *Sina) crawlDailyOld(ctx context.Context, contract *model.TradeObject, startDate, endDate time.Time) ([]*model.KLineNode, error) {
	//crawlDateStr := time.Now().Format("2006_1_2")
	contractId := buildVendorIdByDate(allVendorPrefix[contract.CNName], contract.ContractDate)
	if contractId == "" {
		logs.Info("%v %v not support", contract.ContractCnName, contract.ContractDate)
		return nil, nil
	}
	bodyB, err := http.Get(ctx, "https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var=/InnerFuturesNewService.getDailyKLine", map[string]string{
		"symbol": contractId,
	}, map[string]string{
		"Authority":          "stock2.finance.sina.com.cn",
		"Accept":             "*/*",
		"Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"Referer":            "https://finance.sina.com.cn/futures/quotes/V0.shtml",
		"Sec-Ch-Ua":          "\"Not_A Brand\";v=\"99\", \"Microsoft Edge\";v=\"109\", \"Chromium\";v=\"109\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "\"Linux\"",
		"Sec-Fetch-Dest":     "script",
		"Sec-Fetch-Mode":     "no-cors",
		"Sec-Fetch-Site":     "same-site",
		"User-Agent":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.49",
	}, false)
	if err != nil {
		logs.Info("sina resp data_crawler error, no this contract data_crawler  %v", err)
		return nil, fmt.Errorf("do http error %v", err)
	}

	tmp := strings.Split(string(bodyB), "(")
	if len(tmp) < 2 {
		logs.Info("sina resp data_crawler error, no this contract data_crawler  %v %v", err, string(bodyB))
		return nil, fmt.Errorf("sina resp data_crawler error, no this contract data_crawler  %v %v", err, string(bodyB))
	}
	tmp2 := tmp[1]
	body := tmp2[:len(tmp2)-2]
	allKnodes := []*SinaKnode{}
	err = json.Unmarshal([]byte(body), &allKnodes)
	if err != nil {
		logs.Info("sina resp data_crawler error, no this contract data_crawler  %v", err)
		return nil, fmt.Errorf("sina resp data_crawler error, no this contract data_crawler %v", err)
	}
	return convertSinaKnodeToKnodeOld(ctx, contract, allKnodes), nil
}

func convertSinaKnodeToKnode(ctx context.Context, market *model.TradeObject, in [][]string) []*model.KLineNode {
	out := make([]*model.KLineNode, len(in))
	for i, sinaKnode := range in {
		out[i] = &model.KLineNode{
			ContractCnName: market.ContractCnName,
			ContractDate:   market.ContractDate,
			Type:           model.LineType_Day,
			Open:           StrToFloat(sinaKnode[1]),
			High:           StrToFloat(sinaKnode[2]),
			Low:            StrToFloat(sinaKnode[3]),
			Close:          StrToFloat(sinaKnode[4]),
			Volume:         StrToFloat(sinaKnode[5]),
		}
		t, _ := time.Parse("2006-01-02", sinaKnode[0])
		out[i].TimeStamp = t.Unix()
		out[i].TimeStampDesc = sinaKnode[0]
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].TimeStamp < out[j].TimeStamp
	})
	return out
}

func StrToFloat(in string) float64 {
	o, _ := strconv.ParseFloat(in, 10)
	return o
}

type SinaKnode struct {
	Date   string `json:"d"`
	Open   string `json:"o"`
	High   string `json:"h"`
	Low    string `json:"l"`
	Close  string `json:"c"`
	Volume string `json:"v"`
	//P      string `json:"p"`
	//S      string `json:"s"`
}

func convertSinaKnodeToKnodeOld(ctx context.Context, market *model.TradeObject, in []*SinaKnode) []*model.KLineNode {
	out := make([]*model.KLineNode, len(in))
	for i, sinaKnode := range in {
		out[i] = &model.KLineNode{
			ContractCnName: market.ContractCnName,
			ContractDate:   market.ContractDate,
			Type:           model.LineType_Day,
			TimeStampDesc:  sinaKnode.Date,
			Open:           StrToFloat(sinaKnode.Open),
			Close:          StrToFloat(sinaKnode.Close),
			Low:            StrToFloat(sinaKnode.Low),
			High:           StrToFloat(sinaKnode.High),
			Volume:         StrToFloat(sinaKnode.Volume),
		}
		t, _ := time.Parse("2006-01-02", sinaKnode.Date)
		out[i].TimeStamp = t.Unix()
		out[i].TimeStampDesc = sinaKnode.Date

	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].TimeStamp < out[j].TimeStamp
	})
	return out
}
