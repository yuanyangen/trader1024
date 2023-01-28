package model

type LineType int64

const LineType_Day = 1
const LineType_Minite = 2
const LineType_Hour = 3

type KNode struct {
	High      float64
	Low       float64
	Open      float64
	Close     float64
	Date      string
	TimeStamp int64
}
