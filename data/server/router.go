package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func router(h *server.Hertz) {
	h.POST("/get_all_data", httpHandlerWrapper(GetAllData))
	h.POST("/get_data_by_ts", httpHandlerWrapper(GetDataByTs))
	h.POST("/save_data", httpHandlerWrapper(SaveData))
}
