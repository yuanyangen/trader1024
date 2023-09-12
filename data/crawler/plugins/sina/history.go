package sina

import (
	"encoding/json"
	"fmt"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/data/storage_client"
	"github.com/yuanyangen/trader1024/engine/model"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Sina struct {
}

func (em *Sina) StorageClient() *storage_client.HttpStorageClient {
	return storage_client.SinaHttpStorage()
}

func (em *Sina) CrawlAllMainMarket() []*model.Contract {
	allSubject := markets.GetAllFutureSubjects()
	allMarkets := make([]*model.Contract, len(allSubject))
	for i, v := range allSubject {
		allMarkets[i] = &model.Contract{
			Subject:    v,
			ContractId: v.CNName + "",
		}
	}
	sort.Slice(allMarkets, func(i, j int) bool {
		return allMarkets[i].Exchange+allMarkets[i].ContractId < allMarkets[j].Exchange+allMarkets[j].ContractId
	})
	return allMarkets
}

func (em *Sina) CrawlAllAvailableMainMarket() []*model.Contract {
	allSubject := markets.GetAllFutureSubjects()
	allMarkets := make([]*model.Contract, len(allSubject))
	for i, v := range allSubject {
		for _, d := range getCurrentAvailable() {
			allMarkets[i] = &model.Contract{
				Subject:    v,
				ContractId: v.CNName + d,
			}
		}
	}
	return allMarkets
}
func getCurrentAvailable() []string {
	return nil
}

func (em *Sina) CrawlWeekly(market *model.Contract, startDate, endDate time.Time) ([]*model.KNode, error) {
	return nil, nil
}
func (em *Sina) CrawlMinute(market *model.Contract, startDate, endDate time.Time) ([]*model.KNode, error) {
	return nil, nil
}

func (em *Sina) CrawlDaily(market *model.Contract, startDate, endDate time.Time) ([]*model.KNode, error) {
	crawlDateStr := time.Now().Format("2006_1_2")
	reqUrl := fmt.Sprintf(`https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%%20_%v%v=/InnerFuturesNewService.getDailyKLine?symbol=%v&_=%v`, market.ContractId, crawlDateStr, market.ContractId, crawlDateStr)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authority", "stock2.finance.sina.com.cn")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	//req.Header.Set("Cookie", "UOR=cn.bing.com,finance.sina.com.cn,; SINAGLOBAL=101.38.217.243_1685114877.365135; close_leftanswer=1; U_TRS1=00000026.d78f163c.64d57694.0a143a09; rotatecount=1; Apache=115.47.210.1_1691971932.526467; U_TRS2=00000001.403bc7a4.64d9715e.cb447859; FIN_ALL_VISITED=V0%2COI2309%2CSA2309%2CSA0%2CEB2312%2CUR2401%2CUR0; ULV=1691971961537:6:4:2:115.47.210.1_1691971932.526467:1691971932726; NEWESTVISITED_FUTURE=%7B%22code%22%3A%22V0%22%2C%22hqcode%22%3A%22nf_V0%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22OI2309%22%2C%22hqcode%22%3A%22nf_OI2309%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22SA2309%22%2C%22hqcode%22%3A%22nf_SA2309%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22SA0%22%2C%22hqcode%22%3A%22nf_SA0%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22EB2312%22%2C%22hqcode%22%3A%22nf_EB2312%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22UR2401%22%2C%22hqcode%22%3A%22nf_UR2401%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22UR0%22%2C%22hqcode%22%3A%22nf_UR0%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22EB0%22%2C%22hqcode%22%3A%22nf_EB0%22%2C%22type%22%3A1%7D%7C%7B%22code%22%3A%22EB2304%22%2C%22hqcode%22%3A%22nf_EB2304%22%2C%22type%22%3A1%7D")
	req.Header.Set("Referer", "https://finance.sina.com.cn/futures/quotes/V0.shtml")
	req.Header.Set("Sec-Ch-Ua", "\"Not_A Brand\";v=\"99\", \"Microsoft Edge\";v=\"109\", \"Chromium\";v=\"109\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	req.Header.Set("Sec-Fetch-Dest", "script")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.49")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tmp := strings.Split(string(bodyB), "(")
	if len(tmp) < 2 {
		return nil, fmt.Errorf("sina resp data error")
	}
	tmp2 := tmp[1]
	body := tmp2[:len(tmp2)-2]
	allKnodes := []*SinaKnode{}
	err = json.Unmarshal([]byte(body), &allKnodes)
	if err != nil {
		return nil, fmt.Errorf("sina resp data error")
	}
	return convertSinaKnodeToKnode(allKnodes), nil
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

func convertSinaKnodeToKnode(in []*SinaKnode) []*model.KNode {
	out := make([]*model.KNode, len(in))
	for i, sinaKnode := range in {
		out[i] = &model.KNode{
			Date:   sinaKnode.Date,
			Open:   StrToFloat(sinaKnode.Open),
			Close:  StrToFloat(sinaKnode.Close),
			Low:    StrToFloat(sinaKnode.Low),
			High:   StrToFloat(sinaKnode.High),
			Volume: StrToFloat(sinaKnode.Volume),
		}
		t, _ := time.Parse("2006-01-02", sinaKnode.Date)
		out[i].TimeStamp = t.Unix()
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
