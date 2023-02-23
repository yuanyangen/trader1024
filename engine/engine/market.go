package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/data_feed"
	"github.com/yuanyangen/trader1024/engine/indicator"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/portfolio"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy"
)

type MarketEngine struct {
	Market          *model.Market
	DataFeed        data_feed.DataFeed
	DailyIndicators *indicator.DailyIndicators
	DataFeedChannel chan *data_feed.Data

	Strategies []strategy.Strategy
}

func NewMarket(id string, df data_feed.DataFeed, strategy []strategy.Strategy) *MarketEngine {
	market := markets.GetMarketById(id)
	if market == nil {
		panic("market not define")
	}
	ed := &MarketEngine{
		Market:          market,
		DataFeed:        df,
		DataFeedChannel: make(chan *data_feed.Data, 1024),
		DailyIndicators: &indicator.DailyIndicators{
			Kline:          indicator.NewKLine(market.Name, model.LineType_Day),
			ReceiveChannel: make(chan *data_feed.Data, 1024),
		},
		Strategies: strategy,
	}
	df.RegisterChan(ed.DataFeedChannel)
	return ed
}

func (m *MarketEngine) Start() {
	m.startDataFeed()
	m.initStrategy()
	m.startOnBarLoop()
}

func (m *MarketEngine) DoPlot(p *charts.Page) {
	kline := m.plotKline()
	if m.DailyIndicators != nil && m.DailyIndicators.Kline != nil {
		for _, w := range m.DailyIndicators.Kline.Indicators {
			w.DoPlot(kline)
		}
	}
	p.Add(kline)
}

func (m *MarketEngine) getKline() *indicator.KLineIndicator {
	return m.DailyIndicators.Kline
}

func (m *MarketEngine) startOnBarLoop() {
	utils.AsyncRun(func() {
		for v := range m.DataFeedChannel {
			m.eventHandler(v)
		}
	})
}

func (m *MarketEngine) eventHandler(data *data_feed.Data) {
	m.refreshIndicators(data)

	ctx := &strategy.MarketStrategyContext{
		DailyData: m.DailyIndicators,
		Market:    m.Market,
	}
	req := &model.MarketPortfolioReq{
		Market: m.Market,
	}

	for _, st := range m.Strategies {

		stResult := st.OnBar(ctx, data.KData.TimeStamp)
		if len(stResult) > 0 {
			str := &model.StrategyReq{
				StrategyName: st.Name(),
				Cmds:         stResult,
			}
			req.Strategies = append(req.Strategies, str)
		}
	}
	portfolio.Portfolio(req)
	account.GetAccount().EventTrigger(data.KData.TimeStamp)
}

func (m *MarketEngine) BackTestClearALl() {
	req := &model.MarketPortfolioReq{
		Market: m.Market,
		Strategies: []*model.StrategyReq{
			{
				StrategyName: "BackTestClearAll",
				Cmds:         []*model.StrategyResult{{Cmd: model.StrategyCmdClean}},
			},
		},
	}
	portfolio.Portfolio(req)
}

func (m *MarketEngine) refreshIndicators(data *data_feed.Data) {
	kline := m.getKline()
	kline.AddData(data.KData.TimeStamp, data.KData)
}

func (m *MarketEngine) startDataFeed() {
	if m.DataFeed != nil {
		m.startOneDataFeed(m.DataFeed)
	}
}

func (m *MarketEngine) initStrategy() {
	if m.Strategies != nil {
		for _, stra := range m.Strategies {
			ctx := &strategy.MarketStrategyContext{DailyData: m.DailyIndicators}
			stra.Init(ctx)
		}
	}
}

func (m *MarketEngine) startOneDataFeed(df data_feed.DataFeed) {
	ch := make(chan *data_feed.Data, 1024)
	df.RegisterChan(ch)
	utils.AsyncRun(func() {
		for v := range ch {
			m.DailyIndicators.Kline.AddData(v.KData.TimeStamp, v.KData)
			m.DailyIndicators.ReceiveChannel <- v
		}
	})

}

func (m *MarketEngine) plotKline() *charts.Kline {
	kline := charts.NewKLine()
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: m.Market.MarketId},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	x, y := m.convertData(m.DailyIndicators.Kline)
	kline.AddXAxis(x).AddYAxis("日K", y)
	return kline
}

func (m *MarketEngine) convertData(kline *indicator.KLineIndicator) ([]string, [][4]float32) {
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
