package account

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CashIndicator struct {
	x []string
	y []float64
}

func NewCashIndicator() *CashIndicator {
	cw := &CashIndicator{
		x: []string{},
		y: []float64{},
	}
	return cw
}

func (cw *CashIndicator) DoPlot(p *charts.Page) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "现金"}, charts.YAxisOpts{Scale: true})
	line.AddXAxis(cw.x).AddYAxis("现金", cw.y, charts.LineOpts{Step: false})
	line.SetGlobalOptions(
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	p.Add(line)
}

func (cw *CashIndicator) AddData(ts int64, account float64) {
	t := utils.TsToString(ts)
	cw.x = append(cw.x, t)
	cw.y = append(cw.y, account)
}
