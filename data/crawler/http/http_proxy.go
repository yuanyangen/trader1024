package http

import (
	"context"
	"io"
	"net/http"

	"github.com/yuanyangen/trader1024/engine/logs"
)

// 小幻代理  https://ip.ihuan.me/
// 89免费代理  https://www.89ip.cn/
// 高可用全球免费代理IP库 https://ip.jiangxianli.com/
// 66免费代理  http://www.66ip.cn/
// 开心代理 http://www.kxdaili.com/dailiip.html
// 西拉代理  http://www.xiladaili.com/
// 云代理  ~~http://www.ip3366.net/(不可用)~~
// 企业级高速HTTP代理平台  ~~http://www.data5u.com/(不可用)~~
// 极速代理  ~~https://superfastip.com/#/freeip~~
// 全网代理IP  ~~http://www.goubanjia.com/~~

func Get(ctx context.Context, reqUrl string, queries, headers map[string]string, useProxy bool) (string, error) {
	reqUrl = reqUrl + "?"
	for k, v := range queries {
		reqUrl = reqUrl + k + "=" + v + "&"
	}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	httpClient := &http.Client{}
	if useProxy {
		proxyUri := GetProxy(ctx)
		if proxyUri == nil {
			logs.Info("get proxy error no proxy found %v", err)
			return "", err
		}
		httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUri)}
	}

	resp, err := httpClient.Do(req)
	logs.Info("do http req with proxy %v error:%v", nil, err)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyB, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyB), nil
}

func GetOld(ctx context.Context, reqUrl string, queries, headers map[string]string) (string, error) {
	reqUrl = reqUrl + "?"
	for k, v := range queries {
		reqUrl = reqUrl + k + "=" + v
	}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	proxyUri := GetProxy(ctx)
	if proxyUri == nil {
		logs.Info("get proxy error %v", err)
		return "", err
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUri),
		},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyB, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyB), nil
}
