package account

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type CashIndicator struct {
	cashLine *indicator_base.Line
}

func NewCashIndicator() *CashIndicator {
	cw := &CashIndicator{cashLine: indicator_base.NewLine(model.LineType_Day, "账户金额")}
	return cw
}

func (cw *CashIndicator) DoPlot(p *charts.Page) {
	line := charts.NewLine()

	allData := cw.cashLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.GetTs())
		y[i] = v.GetValue()
	}
	line.SetGlobalOptions(charts.TitleOpts{Title: "现金"}, charts.YAxisOpts{Scale: true})
	line.AddXAxis(x).AddYAxis("现金", y, charts.LineOpts{Step: false})
	line.SetGlobalOptions(
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	p.Add(line)
}

func (cw *CashIndicator) AddData(ts int64, account float64) {
	cw.cashLine.AddData(ts, account)
}
