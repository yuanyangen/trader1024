package stock

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
	"github.com/yuanyangen/trader1024/data/crawler/http"
	"github.com/yuanyangen/trader1024/engine/model"
)

var codeMap = map[string]string{
	"002594.SZ": "sz002594",
}

type Sina struct {
}

func (em *Sina) CrawlMinute(ctx context.Context, market *model.TradeObject) ([]*model.KLineNode, error) {
	return nil, nil
}
func (em *Sina) CrawlWeekly(ctx context.Context, market *model.TradeObject) ([]*model.KLineNode, error) {
	return nil, nil
}
func (em *Sina) convertToCode(ctx context.Context, market *model.TradeObject) string {
	if market == nil || market.SubjectDO == nil || market.SubjectDO.UniqueCode == "" {
		panic("should not reach here")
	}
	v, ok := codeMap[market.UniqueCode]
	if ok {
		return v
	}
	return market.UniqueCode
}

func (em *Sina) CrawlDaily(ctx context.Context, market *model.TradeObject) ([]*model.KLineNode, error) {
	time.Sleep(time.Second * 2)

	// 比亚迪股票代码（深交所：002594.SZ）
	stockCode := em.convertToCode(ctx, market)
	// 新浪财经日K线数据接口（默认返回最近1个月的数据）
	url := "https://quotes.sina.cn/cn/api/json_v2.php/CN_MarketData.getKLineData"
	// 发送HTTP GET请求
	body, err := http.Get(ctx, url, map[string]string{
		"symbol":  stockCode,
		"scale":   "240",
		"ma":      "no",
		"datalen": "1000",
	}, nil, false)
	if err != nil {
		fmt.Printf("HTTP请求失败: %v\n", err)
		return nil, err
	}

	// 解析JSON数据
	var stockDataList []*StockData
	err = sonic.UnmarshalString(body, &stockDataList)
	if err != nil {
		fmt.Printf("JSON解析失败: %v\n", err)
		return nil, err
	}
	return convertSinaKnodeToKnode(ctx, market, stockDataList), nil
}

func StrToFloat(in string) float64 {
	o, _ := strconv.ParseFloat(in, 10)
	return o
}

// 定义新浪财经返回的日K线数据结构
type StockData struct {
	Day    string `json:"day"`    // 日期，格式：YYYY-MM-DD
	Open   string `json:"open"`   // 开盘价
	High   string `json:"high"`   // 最高价
	Low    string `json:"low"`    // 最低价
	Close  string `json:"close"`  // 收盘价
	Volume string `json:"volume"` // 成交量（手）
}

func convertSinaKnodeToKnode(ctx context.Context, market *model.TradeObject, in []*StockData) []*model.KLineNode {
	out := make([]*model.KLineNode, len(in))
	for i, sinaKnode := range in {
		out[i] = &model.KLineNode{
			UniqueCode:    market.Id(),
			Type:          model.LineType_Day,
			TimeStampDesc: sinaKnode.Day,
			Open:          StrToFloat(sinaKnode.Open),
			Close:         StrToFloat(sinaKnode.Close),
			Low:           StrToFloat(sinaKnode.Low),
			High:          StrToFloat(sinaKnode.High),
			Volume:        StrToFloat(sinaKnode.Volume),
		}
		t, _ := time.Parse("2006-01-02", sinaKnode.Day)
		out[i].TimeStamp = t.Unix()
		out[i].TimeStampDesc = sinaKnode.Day

	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].TimeStamp < out[j].TimeStamp
	})
	return out
}
