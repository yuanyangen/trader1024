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

func (em *Sina) StorageClient() *storage_client.HttpStorageClient {
	return storage_client.SinaHttpStorage()
}

func (em *Sina) CrawlAllMainMarket() []*model.Market {
	allSubject := markets.GetAllFutureSubjects()
	allMarkets := make([]*model.Market, len(allSubject))
	for i, v := range allSubject {
		allMarkets[i] = &model.Market{
			Subject:  v,
			VendorId: buildVendorIdByDate(allVendorPrefix[v.Name], ""),
		}
	}
	sort.Slice(allMarkets, func(i, j int) bool {
		return allMarkets[i].Exchange+allMarkets[i].VendorId < allMarkets[j].Exchange+allMarkets[j].VendorId
	})
	return allMarkets
}

func (em *Sina) CrawlAllAvailableMainMarket() []*model.Market {
	allSubject := markets.GetAllFutureSubjects()
	allMarkets := make([]*model.Market, len(allSubject))
	for i, v := range allSubject {
		for _, d := range getCurrentAvailable() {
			allMarkets[i] = &model.Market{
				Subject:  v,
				VendorId: buildVendorIdByDate(allVendorPrefix[v.Name], d),
			}
		}
	}
	return allMarkets
}
func getCurrentAvailable() []string {
	return nil
}

// date
func buildVendorIdByDate(vendorIdPrefix, date string) string {
	if date == "" {
		return vendorIdPrefix + "0"
	}
	t, err := time.Parse(date, "2006-01-02")
	if err != nil {
		panic(err)
	}
	return vendorIdPrefix + t.Format("0102")

}

func (em *Sina) CrawlWeekly(market *model.Market, startDate, endDate time.Time) ([]*model.KNode, error) {
	return nil, nil
}
func (em *Sina) CrawlMinute(market *model.Market, startDate, endDate time.Time) ([]*model.KNode, error) {
	return nil, nil
}

func (em *Sina) CrawlDaily(market *model.Market, startDate, endDate time.Time) ([]*model.KNode, error) {
	crawlDateStr := time.Now().Format("2006_1_2")
	reqUrl := fmt.Sprintf(`https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%%20_%v%v=/InnerFuturesNewService.getDailyKLine?symbol=%v&_=%v`, market.VendorId, crawlDateStr, market.VendorId, crawlDateStr)
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
