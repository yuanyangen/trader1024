package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	model2 "github.com/yuanyangen/trader1024/data/model"
	"github.com/yuanyangen/trader1024/data/storage"
)

func main() {
	Init()
	startDataServer()
}

func Init() {
    const storageDataPath = "/home/yuanyangen/HomeData/go/trader1024/data/datas"
    storage.InitAllStorage(storageDataPath)
}

func startDataServer() {
	h := server.Default()
	router(h)
	h.Spin()
}

func httpHandlerWrapper(handler func(querys string) (any, error)) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		body := ctx.Request.Body()
		code := 0
		codeMsg := ""
		var err error
		var rawResp any
		if err == nil {
			rawResp, err = handler(string(body))
		} else {
			code = 1
			codeMsg = err.Error()
		}
		resp := model2.CommonHttpResp{
			Code:    code,
			CodeMsg: codeMsg,
			Data:    rawResp,
		}
		ctx.JSON(200, resp)
	}
}
