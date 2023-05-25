package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/portfolio"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type MarketEngine struct {
	Market          *model.Market
	DataFeed        model.DataFeed
	DailyIndicators *model.DailyIndicators
	DataFeedChannel chan *model.Data

	Strategies []model.Strategy
}

func NewMarket(id string, df model.DataFeed, strategy []model.Strategy) *MarketEngine {
	market := markets.GetMarketById(id)
	if market == nil {
		panic("market not define")
	}
	ed := &MarketEngine{
		Market:          market,
		DataFeed:        df,
		DataFeedChannel: make(chan *model.Data, 1024),
		DailyIndicators: &model.DailyIndicators{
			Kline:          indicator.NewKLine(market.Name, model.LineType_Day),
			ReceiveChannel: make(chan *model.Data, 1024),
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
	position := account.GetBackTestBroker().GetCurrentLivePositions(m.Market.MarketId)
	position.Report()
	kline := charts.NewKLine()
	m.DailyIndicators.Kline.DoPlot(p, kline)
}

func (m *MarketEngine) getKline() model.MarketIndicator {
	return m.DailyIndicators.Kline
}

func (m *MarketEngine) startOnBarLoop() {
	utils.AsyncRun(func() {
		for v := range m.DataFeedChannel {
			m.eventHandler(v)
		}
	})
}

func (m *MarketEngine) eventHandler(data *model.Data) {
	m.refreshIndicators(data)

	ctx := &model.MarketStrategyContext{
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
				Reason:       st.Name(),
				Ts:           data.KData.TimeStamp,
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

func (m *MarketEngine) refreshIndicators(data *model.Data) {
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
			ctx := &model.MarketStrategyContext{DailyData: m.DailyIndicators}
			stra.Init(ctx)
		}
	}
}

func (m *MarketEngine) startOneDataFeed(df model.DataFeed) {
	ch := make(chan *model.Data, 1024)
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

	return kline
}
