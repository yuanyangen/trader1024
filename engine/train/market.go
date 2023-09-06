package train

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
)

type TrainMarketEngine struct {
	Market          *model.Market
	DataFeed        model.DataFeed
	DailyIndicators *model.DailyIndicators
	DataFeedChannel chan *model.Data
	Strategies      []model.Strategy
	Train           *Train
}

func NewMarket(id string, df model.DataFeed, strategy []model.Strategy) *TrainMarketEngine {
	market := markets.GetMarketById(id)
	if market == nil {
		panic("market not define")
	}
	kline := indicator.NewKLine(market.Name, model.LineType_Day)
	ed := &TrainMarketEngine{
		Market:          market,
		DataFeed:        df,
		DataFeedChannel: make(chan *model.Data, 1024),
		DailyIndicators: &model.DailyIndicators{
			Kline:          kline,
			ReceiveChannel: make(chan *model.Data, 1024),
		},
		Strategies: strategy,
		Train:      NewTrain(market, kline),
	}
	df.RegisterChan(ed.DataFeedChannel)
	return ed
}

func (m *TrainMarketEngine) Start() {
	m.startDataFeed()
	m.initStrategy()
	m.startOnBarLoop()
}

func (m *TrainMarketEngine) DoPlot(p *charts.Page) {
	position := account.GetBackTestBroker().GetCurrentLivePositions(m.Market.MarketId)
	position.Report()
	kline := charts.NewKLine()
	line := charts.NewLine()

	m.DailyIndicators.Kline.DoPlot(kline, line)
	p.Add(kline)
	p.Add(line)
}

func (m *TrainMarketEngine) startOnBarLoop() {
	utils.AsyncRun(func() {
		for v := range m.DataFeedChannel {
			m.eventHandler(v)
		}
	})
}

func (m *TrainMarketEngine) eventHandler(data *model.Data) {
	m.refreshIndicators(data)

	ctx := &model.MarketStrategyContext{
		DailyData: m.DailyIndicators,
		Market:    m.Market,
	}
	req := &model.MarketPortfolioReq{
		Market: m.Market,
		Ts:     data.KData.TimeStamp,
	}

	for _, st := range m.Strategies {
		stResult := st.OnBar(ctx, data.KData.TimeStamp)
		if stResult != nil {
			req.Strategies = append(req.Strategies, &model.StrategyReq{
				StrategyName: st.Name(),
				Cmd:          stResult,
				Reason:       st.Name(),
			})
		}
	}
	m.Train.TrainReq(req)
}

func (m *TrainMarketEngine) refreshIndicators(data *model.Data) {
	kline := m.DailyIndicators.Kline
	kline.AddData(data.KData.TimeStamp, data.KData)
}

func (m *TrainMarketEngine) startDataFeed() {
	if m.DataFeed != nil {
		m.startOneDataFeed(m.DataFeed)
	}
}

func (m *TrainMarketEngine) initStrategy() {
	if m.Strategies != nil {
		for _, stra := range m.Strategies {
			ctx := &model.MarketStrategyContext{DailyData: m.DailyIndicators}
			stra.Init(ctx)
		}
	}
}

func (m *TrainMarketEngine) startOneDataFeed(df model.DataFeed) {
	ch := make(chan *model.Data, 1024)
	df.RegisterChan(ch)
	utils.AsyncRun(func() {
		for v := range ch {
			m.DailyIndicators.Kline.AddData(v.KData.TimeStamp, v.KData)
			m.DailyIndicators.ReceiveChannel <- v
		}
	})
}

func (m *TrainMarketEngine) plotKline() *charts.Kline {
	kline := charts.NewKLine()
	return kline
}
