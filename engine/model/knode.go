package model

type LineType int64

const LineType_Day LineType = 1
const LineType_Minite LineType = 2
const LineType_5Minite LineType = 3
const LineType_Hour LineType = 4

type KNode struct {
	Date          string
	High          float64
	Low           float64
	Open          float64
	Close         float64
	Volume        float64
	Turnover      float64 // 成交额
	Swing         float64 //振幅
	Increase      float64 // 涨跌幅 ??
	IncreaseMount float64 // 涨跌额 ??
	TurnoverRate  float64 //换手率
	TimeStamp     int64
}
