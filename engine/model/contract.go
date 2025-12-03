package model

import (
	"github.com/bytedance/sonic"
)

type MarKetType int64

const MainContinuous = "main_continuous"

const MarKetType_STOCK MarKetType = 1
const MarKetType_FUTURE MarKetType = 2

type TradeObject struct {
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

func (c *TradeObject) Id() string {
	if c.Type == MarKetType_STOCK {
		return c.UniqueCode
	}
	return c.CNName + c.ContractDate
}

// type Exchange struct {
// 	Name string
// }

// type ExchangeTime struct {
// 	CollectionBiddingDeclaration string
// 	CollectionBiddingMatchmaking string
// 	TradeTimes1                  []string
// 	TradeTimes2                  []string
// }

type SubjectDO struct {
	CNName       string //CN name
	UniqueCode   string // 唯一编码， 格式是： 002594.SZ
	Type         MarKetType
	Exchange     string // 交易所的名字
	OnlineDay    string
	TypeLevel1   []string //按照 中国上市公司分类指引 ， BYD的一级分类，二级分类分别是什么？
	TypeLevel2   []string
	ContractInfo *ContractInfo
}
type ContractInfo struct {
	FirstContract            string
	OnlineTime               string   //第一次上线交易的时间 20060102
	OfflineTime              string   //最后一次上线交易的时间 20060102
	DailyExchangeTime        []string // 每天交易的时间
	ContractMonth            []int
	ContractLastTradeDay     int // 合约的最后一个交易日
	RealContractLastTradeDay int // 散户实际的最后一个交易日
}

func (s *SubjectDO) StartDate() string {
	return s.ContractInfo.OnlineTime
}

func (s *SubjectDO) EndDate() string {
	return s.ContractInfo.OfflineTime
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
