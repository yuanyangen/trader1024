package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	model2 "github.com/yuanyangen/trader1024/data/model"
	stroage_server "github.com/yuanyangen/trader1024/data/storage_server"
)

func router(h *server.Hertz) {
	h.POST("/get_all_data", httpHandlerWrapper(stroage_server.GetAllData))
	h.POST("/get_data_by_ts", httpHandlerWrapper(stroage_server.GetDataByTs))
	h.POST("/save_data", httpHandlerWrapper(stroage_server.SaveData))
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
