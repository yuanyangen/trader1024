package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/account/local_account"
	"github.com/yuanyangen/trader1024/engine/model"
	"net/http"
	"sort"
)

type WatcherBackend struct {
	plotterInfos []model.Plotter
}

func NewPlotterServers() *WatcherBackend {
	ps := &WatcherBackend{}
	return ps
}

func (ps *WatcherBackend) AddPlotter(p model.Plotter) {
	ps.plotterInfos = append(ps.plotterInfos, p)
}

func (ps *WatcherBackend) httpHandler(w http.ResponseWriter, _ *http.Request) {
	p := charts.NewPage()
	sort.Slice(ps.plotterInfos, func(i, j int) bool {
		return ps.plotterInfos[i].Name() > ps.plotterInfos[j].Name()
	})
	local_account.DefaultLocalAccount.DoPlot(p)
	for _, v := range ps.plotterInfos {
		v.DoPlot(p)
	}
	p.Render(w) // Render 可接收多个 io.Writer 接口
}

func (ps *WatcherBackend) Start() {
	http.HandleFunc("/", ps.httpHandler)
	http.ListenAndServe(":8081", nil)
}
