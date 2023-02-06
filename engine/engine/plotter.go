package engine

import (
	"github.com/go-echarts/go-echarts/charts"
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
	ps.Engine.Account.DoPlot(p)
	for _, v := range ps.Engine.Markets {
		v.DoPlot(p)
	}

	p.Render(w) // Render 可接收多个 io.Writer 接口
}

func (ps *WatcherBackend) Start() {
	http.HandleFunc("/", ps.httpHandler)
	http.ListenAndServe(":8081", nil)
}
