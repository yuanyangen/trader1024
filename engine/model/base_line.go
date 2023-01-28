package model

import (
	"fmt"
	"sync"
)

type BaseLine struct {
	Type    LineType
	StartTs int64
	EndTs   int64
	Mu      sync.Mutex
	data    map[int64]any
}

func (bl *BaseLine) UnityTimeStamp(ts int64) int64 {
	offset := int64(0)
	if bl.Type == LineType_Day {
		offset = 86400
	}
	ts = (ts / offset) * offset
	return ts
}

func (bl *BaseLine) GetByTsAndCount(ts int64, count int64) ([]any, error) {
	var offset int64

	ts = bl.UnityTimeStamp(ts)
	resp := make([]any, count)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	for i := int64(0); i < count; i++ {
		timeK := ts - i*offset
		node, ok := bl.data[timeK]
		if ok {
			resp[count-i-1] = node
		} else {
			return nil, fmt.Errorf("no data for %v", timeK)
		}
	}
	return resp, nil

}

func (bl *BaseLine) AddData(ts int64, node any) {
	ts = bl.UnityTimeStamp(ts)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	bl.data[ts] = node
	if bl.EndTs < ts {
		bl.EndTs = ts
	}
	if bl.StartTs > ts {
		bl.StartTs = ts
	}
}

func (bl *BaseLine) GetAllSortedData() []any {
	if bl == nil {
		return nil
	}
	bl.Mu.Lock()

	res := make([]any, len(bl.data))
	i := 0
	for _, v := range bl.data {
		res[i] = v
		i++
	}
	bl.Mu.Unlock()
	return res
}
