package utils

import "time"

const DailyFormat = "2006-01-02"

func TsToString(ts int64) string {
	t := time.Unix(ts, 0).Format(DailyFormat)
	return t
}
