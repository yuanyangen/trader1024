package model

import "time"

type MarKetType int64

const MarKetType_STOCK MarKetType = 1
const MarKetType_FUTURE MarKetType = 2

type Market struct {
	*Subject
	VendorId string
}

type Exchange struct {
	Name string
}

type ExchangeTime struct {
	CollectionBiddingDeclaration string
	CollectionBiddingMatchmaking string
	TradeTimes1                  []string
	TradeTimes2                  []string
}

type Subject struct {
	Name          string //CN name
	Type          MarKetType
	Exchange      string
	OnlineDay     string
	FirstContract string

	OnlineTime        time.Time //第一次上线交易的时间
	OfflineTime       time.Time //最后一次上线交易的时间
	DailyExchangeTime []string
}

func (s *Subject) StartDate() string {
	if s == nil || s.OnlineTime.Unix() == 0 {
		panic("online data error")
	}
	return s.OnlineTime.Format("060102")
}

func (s *Subject) EndDate() string {
	if s == nil || s.OfflineTime.Unix() == 0 {
		panic("offline data error")
	}
	return s.OfflineTime.Format("060102")
}

func (s *Subject) AllDates() []string {
	end := s.OfflineTime
	if end.Unix() == 0 {
		end = time.Now().Add(8 * 30 * 24 * time.Hour)
	}
	out := []string{}
	for st := s.OnlineTime; st.Before(end); st.Add(time.Hour * 24) {
		out = append(out, st.Format("060102"))
	}
	return out
}

type MarketPortfolioReq struct {
	Market     *Market
	Strategies []*StrategyReq
	Ts         int64
}
type StrategyReq struct {
	StrategyName string
	Cmd          *StrategyResult
	Reason       string
}
