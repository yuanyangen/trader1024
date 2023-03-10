package utils

import "time"

const DailyFormat = "2006-01-02"

func TsToString(ts int64) string {
	t := time.Unix(ts, 0).Format(DailyFormat)
	return t
}

func UnityTimeStamp(ts int64, offset int64) int64 {
	ts = (ts / offset) * offset
	return ts
}

func UnityDailyTimeStamp(ts int64) int64 {
	return UnityTimeStamp(ts, 86400)
}
