package http

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"io"
	"net/http"
	"time"
)

const proxyAddr51 = "http://bapi.51daili.com/getapi2?linePoolIndex=-1&packid=2&time=5&qty=1&port=2&format=json&field=ipport,expiretime&dt=3&ct=1&usertype=17&uid=46309&accessName=yuanyangen&accessPassword=A6656E662086F719CA2A6E4E4A5CFCBA&skey=autoaddwhiteip"

func getProxyAddrFrom51Proxy() ([]*model.HttpProxy, error) {
	resp, err := http.Get(proxyAddr51)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respS := &Proxy51Resp{}
	sonic.Unmarshal(bodyB, &respS)
	if respS != nil && len(respS.Data) > 0 {
		out := []*model.HttpProxy{}
		for _, v := range respS.Data {
			expireTs := utils.StrToTs(v.ExpireTime, "2006-01-02 15:04:05")
			if expireTs > time.Now().Unix() {
				out = append(out, &model.HttpProxy{
					Addr:       v.IP,
					ExpireTime: expireTs,
					ProxyType:  "socks5",
					UserName:   "yuanyangen",
					Password:   "A6656E662086F719CA2A6E4E4A5CFCBA",
				})
			}
		}
		return out, nil
	}

	return nil, fmt.Errorf("no proxy addr found ")
}

type Proxy51Resp struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    []struct {
		IP         string `json:"IP"`
		ExpireTime string `json:"ExpireTime"`
		IPAddress  string `json:"IpAddress"`
		Isp        string `json:"ISP"`
	} `json:"data_crawler"`
}
