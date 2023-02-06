package account

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CashIndicator struct {
	x []string
	y []int64
}

func NewCashIndicator() *CashIndicator {
	cw := &CashIndicator{
		x: []string{},
		y: []int64{},
	}
	return cw
}

func (cw *CashIndicator) DoPlot(p *charts.Page) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "现金"})
	line.AddXAxis(cw.x).AddYAxis("现金", cw.y)
	p.Add(line)
}

func (cw *CashIndicator) AddData(ts int64, account int64) {
	t := utils.TsToString(ts)
	cw.x = append(cw.x, t)
	cw.y = append(cw.y, account)
}
