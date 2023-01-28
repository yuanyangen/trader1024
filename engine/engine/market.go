package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator/market"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type Market struct {
	Name       string
	DailyData  *model.DailyData
	Broker     model.Broker
	Account    *model.Account
	Strategys  []model.Strategy
	globalChan chan *model.GlobalMsg
}

func NewMarket(name string, df model.DataFeed, strategy []model.Strategy) *Market {
	ed := &Market{
		Name: name,
		DailyData: &model.DailyData{
			DataFeed:       df,
			Line:           market.NewKLine(model.LineType_Day),
			ReceiveChannel: make(chan *model.Data, 1024),
		},
		Strategys: strategy,
	}
	if ed.DailyData != nil && ed.DailyData.ReceiveChannel != nil {
		df.RegisterChan(ed.DailyData.ReceiveChannel)
	}
	return ed
}

func (m *Market) Start() {
	m.startDataFeed()
	m.startOnBarLoop()
}

func (m *Market) DoPlot(p *charts.Page) {
	kline := m.plotKline()
	p.Add(kline)
	if m.DailyData != nil {
		for _, w := range m.DailyData.Indicators {
			w.DoPlot(kline)
		}
	}

}

func (m *Market) getKline() *market.KLineIndicator {
	return m.DailyData.Line
}

func (m *Market) startOnBarLoop() {
	utils.AsyncRun(func() {
		for v := range m.DailyData.ReceiveChannel {
			m.eventHandler(v)
		}
	})
}

func (m *Market) eventHandler(data *model.Data) {
	ctx := &model.MarketStrategyContext{
		DailyData: m.DailyData,
		Broker:    m.Broker,
		Account:   m.Account,
	}
	for _, st := range m.Strategys {
		st.OnBar(ctx)
	}
	if m.globalChan != nil {
		msg := &model.GlobalMsg{}
		m.globalChan <- msg
	}
}

func (m *Market) startDataFeed() {
	if m.DailyData != nil {
		m.startOneDataFeed(m.DailyData.DataFeed)
	}
}

func (m *Market) startOneDataFeed(df model.DataFeed) {
	ch := make(chan *model.Data, 1024)
	df.RegisterChan(ch)
	utils.AsyncRun(func() {
		for v := range ch {
			m.DailyData.Line.AddData(v.KData.TimeStamp, v.KData)
			m.DailyData.ReceiveChannel <- v
		}
	})

	df.StartFeed()
}

func (m *Market) plotKline() *charts.Kline {
	kline := charts.NewKLine()
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: m.Name},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	x, y := m.convertData(m.DailyData.Line)
	kline.AddXAxis(x).AddYAxis("日K", y)
	return kline
}

func (m *Market) convertData(kline *market.KLineIndicator) ([]string, [][4]float32) {
	kDatas := kline.GetAllSortedData()
	x := make([]string, len(kDatas))
	y := make([][4]float32, len(kDatas))
	for i, kn := range kDatas {
		x[i] = kn.Date
		y[i] = [4]float32{
			float32(kn.Open),
			float32(kn.Close),
			float32(kn.Low),
			float32(kn.High),
		}
	}
	return x, y
}
