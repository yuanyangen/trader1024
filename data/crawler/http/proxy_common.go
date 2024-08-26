package http

import (
	"context"
	"github.com/yuanyangen/trader1024/dal/mongo"
	"github.com/yuanyangen/trader1024/engine/model"
	"net/url"
	"time"
)

var allProxyCrawler = []func() ([]*model.HttpProxy, error){
	getProxyAddrFrom51Proxy,
}

func CrawlProxy(ctx context.Context) {
	for _, cralwer := range allProxyCrawler {
		proxys, _ := cralwer()
		for _, p := range proxys {
			mongo.SaveProxy(ctx, p)
		}
	}
}

func GetProxy(ctx context.Context) *url.URL {
	u := GetProxyFromDb(ctx)
	if u == nil {
		CrawlProxy(ctx)
	}
	u = GetProxyFromDb(ctx)
	return u
}

func GetProxyFromDb(ctx context.Context) *url.URL {
	proxyUrl, err := mongo.QueryProxyWithExpireTime(ctx, time.Now().Unix())
	if err != nil || proxyUrl == nil {
		return nil
	}
	if proxyUrl.ProxyType == "http" {
		proxyUri, _ := url.Parse("http://" + proxyUrl.Addr)
		if proxyUri != nil && proxyUrl.UserName != "" {
			proxyUri.User = url.UserPassword(proxyUrl.UserName, proxyUrl.Password)
		}
		return proxyUri
	} else if proxyUrl.ProxyType == "socks5" {
		proxyUri, _ := url.Parse("socks5://" + proxyUrl.Addr)
		if proxyUri != nil && proxyUrl.UserName != "" {
			proxyUri.User = url.UserPassword(proxyUrl.UserName, proxyUrl.Password)
		}
		return proxyUri
	}
	return nil
}
