package local_broker

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"sync"
)

var DefaultLocalAccount *LocalAccount

type LocalAccount struct {
	Total       decimal.Decimal // 单位是分
	GlobalEvent chan *model.EventMsg
	AccountLine *model.BaseLine
	mu          sync.Mutex
}

func InitLocalAccount(start int64) {
	DefaultLocalAccount = &LocalAccount{
		Total:       decimal.NewFromInt(start),
		GlobalEvent: make(chan *model.EventMsg, 1024),
		AccountLine: model.NewBaseLine("账户金额", model.LineType_Day),
	}
}

func (a *LocalAccount) DoPlot(p *charts.Page) {
	a.mu.Lock()
	defer a.mu.Unlock()

	line := charts.NewLine()

	allData := a.AccountLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.TimeStamp)
		value := v.DataNode.(float64)
		y[i] = value
	}
	line.SetGlobalOptions(charts.TitleOpts{Title: "现金"}, charts.YAxisOpts{Scale: true})
	line.AddXAxis(x).AddYAxis("现金", y, charts.LineOpts{Step: false})
	line.SetGlobalOptions(
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	p.Add(line)

	a.showFinalNum()
}
func (a *LocalAccount) showFinalNum() {
	total := a.Total
	logs.Info(total.String())
}

func (a *LocalAccount) DealEvent(event *model.EventMsg) {
	ts := event.TimeStamp
	currentVal, _ := a.Total.Float64()
	a.AccountLine.AddData(ts, &model.LineNode{TimeStamp: ts, DataNode: currentVal})
}

func (a *LocalAccount) ChangeValue(count decimal.Decimal) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Total = a.Total.Add(count)
}
