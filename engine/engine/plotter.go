package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/account"
	"net/http"
)

type WatcherBackend struct {
	Engine *Engine
}

func NewPlotterServers(Engine *Engine) *WatcherBackend {
	ps := &WatcherBackend{
		Engine: Engine,
	}
	return ps
}

func (ps *WatcherBackend) httpHandler(w http.ResponseWriter, _ *http.Request) {
	p := charts.NewPage()
	account.GetAccount().DoPlot(p)
	for _, m := range ps.Engine.Contracts {
		//m.BackTestClearALl()
		m.DoPlot(p)
	}

	p.Render(w) // Render 可接收多个 io.Writer 接口
}

func (ps *WatcherBackend) Start() {
	http.HandleFunc("/", ps.httpHandler)
	http.ListenAndServe(":8081", nil)
}
