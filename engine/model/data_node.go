package model

type LineType int64

const LineType_Day LineType = 1
const LineType_Minite LineType = 2
const LineType_5Minite LineType = 3
const LineType_Hour LineType = 4
const LineType_Week LineType = 5

func (ln *LineNode) IsValid() bool {
	if ln == nil {
		return false
	}
	if ln.DataNode == nil || ln.TimeStamp == 0 {
		return false
	}
	return true
}
