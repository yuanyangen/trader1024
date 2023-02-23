package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/model"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 数据是从 http://quote.eastmoney.com/zs000905.html#fullScreenChart 这个页面获取的。 真难找。。
type EastMoney struct {
}

var EastmoneyParamsMap = map[string]string{
	"f51": "Date",
	"f52": "Open",
	"f53": "Close",
	"f54": "High",
	"f55": "Low",
	"f56": "Volume",
	"f57": "成交额",
	"f58": "振幅",
	"f59": "涨跌幅",
	"f60": "涨跌额",
	"f61": "换手率",
}

func (em *EastMoney) CrawlDaily(marketId string, startDate string, endDate string) ([]*model.KNode, error) {
	market := markets.GetMarketById(marketId)
	if market == nil {
		return nil, fmt.Errorf("no market")
	}
	req := &EastMoneyReq{
		Fields1:    "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
		Fields2:    "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		StartDate:  startDate,
		EndDate:    endDate,
		KlineType:  101,
		ReturnType: 6,
		FuQuanType: 2,
		secId:      market.SecId,
	}
	return em.doCrawl(req, "2006-01-02")
}

func (em *EastMoney) CrawlMinute5(marketId string, startDate string, endDate string) ([]*model.KNode, error) {
	market := markets.GetMarketById(marketId)
	if market == nil {
		return nil, fmt.Errorf("no market")
	}
	req := &EastMoneyReq{
		Fields1:    "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
		Fields2:    "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		StartDate:  startDate,
		EndDate:    endDate,
		KlineType:  5,
		ReturnType: 6,
		FuQuanType: 2,
		secId:      market.SecId,
		Lmt:        120,
	}
	return em.doCrawl(req, "2006-01-02 15:04")

}

type EastMoneyReq struct {
	Fields1    string
	Fields2    string
	StartDate  string // 开始日期 例如 20200101
	EndDate    string //结束日期 例如 20200201
	ReturnType int
	KlineType  int //1: 分钟 5: 5分钟 101: 日 102: 周
	FuQuanType int // 复权方式  不复权 : 0 前复权 : 1 后复权 : 2
	secId      string
	Lmt        int
}

var eastmoneyHeaders = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko",
	"Accept":          "*/*",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Referer":         "http://quote.eastmoney.com/center/gridlist.html",
}

func (em *EastMoney) doCrawl(req *EastMoneyReq, dateformat string) ([]*model.KNode, error) {
	params := url.Values{}
	params.Set("fields1", req.Fields1)
	params.Set("fields2", req.Fields2)
	//params.Set("beg", req.StartDate)
	params.Set("end", req.EndDate)
	params.Set("rtntype", fmt.Sprint(req.ReturnType))
	params.Set("secid", fmt.Sprint(req.secId))
	params.Set("fqt", fmt.Sprint(req.FuQuanType))
	params.Set("klt", fmt.Sprint(req.KlineType))
	params.Set("lmt", fmt.Sprint(req.Lmt))
	eastMoneyUrl, err := url.Parse("https://push2his.eastmoney.com/api/qt/stock/kline/get")
	if err != nil {
		return nil, err
	}
	eastMoneyUrl.RawQuery = params.Encode()

	httpReq, _ := http.NewRequest("GET", eastMoneyUrl.String(), nil)
	for k, v := range eastmoneyHeaders {
		httpReq.Header.Add(k, v)
	}

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	bodyB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	//fmt.Println(string(bodyB))

	res := &EastMoneyResp{}
	err = json.Unmarshal(bodyB, &res)
	if err != nil {
		return nil, err
	}
	if res == nil || res.Data.Klines == nil {
		return nil, nil
	}

	return convertDataToStruct(res.Data.Klines, dateformat), nil
}
func convertDataToStruct(in []string, dateFormate string) []*model.KNode {
	res := make([]*model.KNode, len(in))

	for i, oneK := range in {
		tmpData := strings.Split(oneK, ",")
		knode := &model.KNode{
			Date:          tmpData[0],
			Open:          convertStrToFloat(tmpData[1]),
			Close:         convertStrToFloat(tmpData[2]),
			High:          convertStrToFloat(tmpData[3]),
			Low:           convertStrToFloat(tmpData[4]),
			Volume:        convertStrToFloat(tmpData[5]),
			Turnover:      convertStrToFloat(tmpData[6]),
			Swing:         convertStrToFloat(tmpData[7]),
			Increase:      convertStrToFloat(tmpData[8]),
			IncreaseMount: convertStrToFloat(tmpData[9]),
			TurnoverRate:  convertStrToFloat(tmpData[10]),
		}
		t, _ := time.Parse(dateFormate, tmpData[0])
		knode.TimeStamp = t.Unix()

		res[i] = knode
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].TimeStamp < res[j].TimeStamp
	})
	return res
}

func convertStrToFloat(in string) float64 {
	r, _ := strconv.ParseFloat(in, 64)
	return r
}

type EastMoneyResp struct {
	//Rc     int    `json:"rc"`
	//Rt     int    `json:"rt"`
	//Svr    int    `json:"svr"`
	//Lt     int    `json:"lt"`
	//Full   int    `json:"full"`
	//Dlmkts string `json:"dlmkts"`
	Data struct {
		//Code       string   `json:"code"`
		//Market     int      `json:"market"`
		//Name       string   `json:"name"`
		//Decimal    int      `json:"decimal"`
		//Dktotal    int      `json:"dktotal"`
		//PreKPrice  int      `json:"preKPrice"`
		//PrePrice   int      `json:"prePrice"`
		//QtMiscType int      `json:"qtMiscType"`
		Klines []string `json:"klines"`
	} `json:"data"`
}
