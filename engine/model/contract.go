package model

import (
	"github.com/bytedance/sonic"
)

type MarKetType int64

const MarKetType_STOCK MarKetType = 1
const MarKetType_FUTURE MarKetType = 2

type Contract struct {
	*SubjectDO
	*ContractDO
}
type ContractDO struct {
	ContractCnName       string
	ContractDate         string
	ContractStartTime    int64
	ContractEndTime      int64
	ContractTradeEndTime int64
	ContractDeliveryTime int64
	LastCrawlTime        int64
}

func (s *ContractDO) String() string {
	r, _ := sonic.MarshalString(s)
	return r
}

func (c *Contract) Id() string {
	return c.CNName + c.ContractDate
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

type SubjectDO struct {
	CNName        string //CN name
	Type          MarKetType
	Exchange      string
	OnlineDay     string
	FirstContract string

	OnlineTime               string   //第一次上线交易的时间 20060102
	OfflineTime              string   //最后一次上线交易的时间 20060102
	DailyExchangeTime        []string // 每天交易的时间
	ContractMonth            []int
	ContractLastTradeDay     int // 合约的最后一个交易日
	RealContractLastTradeDay int // 散户实际的最后一个交易日
}

func (s *SubjectDO) StartDate() string {
	return s.OnlineTime
}

func (s *SubjectDO) EndDate() string {
	return s.OfflineTime
}

func (s *SubjectDO) AllDates() []string {
	//end := s.OfflineTime
	//if end.Unix() == 0 {
	//	end = time.Now().Add(8 * 30 * 24 * time.Hour)
	//}
	var out []string
	//for st := s.OnlineTime; st.Before(end); st.Add(time.Hour * 24) {
	//	out = append(out, st.Format("060102"))
	//}
	return out
}

func (s *SubjectDO) String() string {
	r, _ := sonic.MarshalString(s)
	return r
}
