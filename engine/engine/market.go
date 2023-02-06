package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type Market struct {
	Name            string
	DataFeed        data_feed.DataFeed
	DailyIndicators *indicator.DailyIndicators
	Strategies      []model.Strategy
}

func NewMarket(name string, df data_feed.DataFeed, strategy []model.Strategy) *Market {
	ed := &Market{
		Name:     name,
		DataFeed: df,
		DailyIndicators: &indicator.DailyIndicators{
			Line:           indicator.NewKLine(indicator.LineType_Day),
			ReceiveChannel: make(chan *data_feed.Data, 1024),
		},
		Strategies: strategy,
	}
	if ed.DailyIndicators != nil && ed.DailyIndicators.ReceiveChannel != nil {
		df.RegisterChan(ed.DailyIndicators.ReceiveChannel)
	}
	return ed
}

func (m *Market) Start() {
	m.startDataFeed()
	m.initStrategy()
	m.startOnBarLoop()
}

func (m *Market) DoPlot(p *charts.Page) {
	kline := m.plotKline()
	if m.DailyIndicators != nil && m.DailyIndicators.Line != nil {
		for _, w := range m.DailyIndicators.Line.Indicators {
			w.DoPlot(kline)
		}
	}
	p.Add(kline)
}

func (m *Market) getKline() *indicator.KLineIndicator {
	return m.DailyIndicators.Line
}

func (m *Market) startOnBarLoop() {
	utils.AsyncRun(func() {
		for v := range m.DailyIndicators.ReceiveChannel {
			m.eventHandler(v)
		}
	})
}

func (m *Market) eventHandler(data *data_feed.Data) {
	ctx := &model.MarketStrategyContext{
		DailyData: m.DailyIndicators,
	}
	for _, st := range m.Strategies {
		st.OnBar(ctx, data.KData.TimeStamp)
	}
}

func (m *Market) startDataFeed() {
	if m.DataFeed != nil {
		m.startOneDataFeed(m.DataFeed)
	}
}

func (m *Market) initStrategy() {
	if m.Strategies != nil {
		for _, strategy := range m.Strategies {
			ctx := &model.MarketStrategyContext{DailyData: m.DailyIndicators}
			strategy.Init(ctx)
		}
	}
}

func (m *Market) startOneDataFeed(df data_feed.DataFeed) {
	ch := make(chan *data_feed.Data, 1024)
	df.RegisterChan(ch)
	utils.AsyncRun(func() {
		for v := range ch {
			m.DailyIndicators.Line.AddData(v.KData.TimeStamp, v.KData)
			m.DailyIndicators.ReceiveChannel <- v
		}
	})

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
	x, y := m.convertData(m.DailyIndicators.Line)
	kline.AddXAxis(x).AddYAxis("日K", y)
	return kline
}

func (m *Market) convertData(kline *indicator.KLineIndicator) ([]string, [][4]float32) {
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
