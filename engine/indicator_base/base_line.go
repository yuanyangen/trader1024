package indicator_base

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"sort"
	"sync"
)

type BaseLine struct {
	Name    string
	Type    model.LineType
	StartTs int64
	EndTs   int64
	Mu      sync.Mutex
	data    map[int64]model.DataNode
}

func NewBaseLine(name string, t model.LineType) *BaseLine {
	bl := &BaseLine{
		Name: name,
		Type: t,
		data: map[int64]model.DataNode{},
	}
	return bl
}
func (bl *BaseLine) offset() int64 {
	if bl.Type == model.LineType_Day {
		return 86400
	} else if bl.Type == model.LineType_Hour {
		return 1440
	} else if bl.Type == model.LineType_Minite {
		return 60
	}
	panic("not support")
}

func (bl *BaseLine) UnityTimeStamp(ts int64) int64 {
	offset := bl.offset()
	return utils.UnityTimeStamp(ts, offset)
}
func (bl *BaseLine) GetByTs(ts int64) (model.DataNode, error) {
	ts = bl.UnityTimeStamp(ts)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	node, ok := bl.data[ts]
	if ok {
		return node, nil
	} else {
		return nil, fmt.Errorf("no data for_%v", ts)
	}
}

func (bl *BaseLine) GetLastByTsAndCount(ts int64, count int64) ([]model.DataNode, error) {
	offset := bl.offset()
	ts = bl.UnityTimeStamp(ts)
	resp := make([]model.DataNode, count)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	found := int64(0)
	for i := int64(0); found < count; i++ {
		timeK := ts - i*offset
		if timeK < bl.StartTs {
			return nil, fmt.Errorf("no enough data for %v", timeK)
		}

		node, ok := bl.data[timeK]
		if ok {
			resp[count-found-1] = node
			found++
		}
	}
	return resp, nil
}

func (bl *BaseLine) GetForwardByTsAndCount(ts int64, count int64) ([]model.DataNode, error) {
	offset := bl.offset()
	ts = bl.UnityTimeStamp(ts)
	resp := make([]model.DataNode, count)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	found := int64(0)
	for i := int64(0); found < count; i++ {
		timeK := ts + i*offset
		if timeK > bl.EndTs {
			return nil, fmt.Errorf("no enough data for %v", timeK)
		}

		node, ok := bl.data[timeK]
		if ok {
			resp[found] = node
			found++
		}
	}
	return resp, nil
}

func (bl *BaseLine) AddData(ts int64, node model.DataNode) {
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

func (bl *BaseLine) GetAllData() []model.DataNode {
	if bl == nil {
		return nil
	}
	bl.Mu.Lock()

	res := make([]model.DataNode, len(bl.data))
	i := 0
	for _, v := range bl.data {
		res[i] = v
		i++
	}
	bl.Mu.Unlock()
	return res
}

func (bl *BaseLine) GetAllSortedData() []model.DataNode {
	ldRes := bl.GetAllData()
	sort.Slice(ldRes, func(i, j int) bool {
		return ldRes[i].GetTs() < ldRes[j].GetTs()
	})
	return ldRes
}
