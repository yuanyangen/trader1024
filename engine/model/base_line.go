package model

import (
	"fmt"
	"github.com/yuanyangen/trader1024/engine/utils"
	"slices"
	"sort"
	"sync"
)

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

type BaseLine struct {
	Name    string
	Type    LineType
	StartTs int64
	EndTs   int64
	Mu      sync.Mutex
	data    map[int64]*LineNode
}

type LineNode struct {
	DataNode  any
	TimeStamp int64
}

func NewBaseLine(name string, t LineType) *BaseLine {
	bl := &BaseLine{
		Name: name,
		data: map[int64]*LineNode{},
		Type: t,
	}
	return bl
}
func (bl *BaseLine) offset() int64 {
	if bl.Type == LineType_Day {
		return 86400
	} else if bl.Type == LineType_Hour {
		return 1440
	} else if bl.Type == LineType_Minite {
		return 60
	}
	panic("not support")
}

func (bl *BaseLine) UnityTimeStamp(ts int64) int64 {
	offset := bl.offset()
	return utils.UnityTimeStamp(ts, offset)
}
func (bl *BaseLine) GetByTs(ts int64) (*LineNode, error) {
	ts = bl.UnityTimeStamp(ts)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	node, ok := bl.data[ts]
	if ok {
		return node, nil
	} else {
		return nil, fmt.Errorf("no data_crawler for_%v", ts)
	}
}

// get last one
func (bl *BaseLine) GetLastByTs(ts int64) (*LineNode, error) {
	nodes, err := bl.GetLastByTsAndCount(ts, 1)
	if err != nil {
		return nil, err
	}

	return nodes[0], nil
}

func (bl *BaseLine) GetLastByTsAndCount(ts int64, count int64) (resp []*LineNode, err error) {
	offset := bl.offset()
	ts = bl.UnityTimeStamp(ts)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	found := int64(0)
	for i := int64(0); found < count; i++ {
		timeK := ts - i*offset
		if timeK < bl.StartTs {
			break
		}

		node, ok := bl.data[timeK]
		if ok {
			resp = append(resp, node)
			found++
		}
	}
	if found < count {
		err = fmt.Errorf("no enough data_crawler for %v", count)
	}
	slices.Reverse(resp)
	return resp, err
}

func (bl *BaseLine) GetForwardByTsAndCount(ts int64, count int64) ([]*LineNode, error) {
	offset := bl.offset()
	ts = bl.UnityTimeStamp(ts)
	resp := make([]*LineNode, count)
	bl.Mu.Lock()
	defer bl.Mu.Unlock()
	found := int64(0)
	for i := int64(0); found < count; i++ {
		timeK := ts + i*offset
		if timeK > bl.EndTs {
			return nil, fmt.Errorf("no enough data_crawler for %v", timeK)
		}

		node, ok := bl.data[timeK]
		if ok {
			resp[found] = node
			found++
		}
	}
	return resp, nil
}

func (bl *BaseLine) AddData(ts int64, node *LineNode) {
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

func (bl *BaseLine) GetAllData() []*LineNode {
	if bl == nil {
		return nil
	}
	bl.Mu.Lock()

	res := make([]*LineNode, len(bl.data))
	i := 0
	for _, v := range bl.data {
		res[i] = v
		i++
	}
	bl.Mu.Unlock()
	return res
}

func (bl *BaseLine) GetAllSortedData() []*LineNode {
	ldRes := bl.GetAllData()
	sort.Slice(ldRes, func(i, j int) bool {
		return ldRes[i].TimeStamp < ldRes[j].TimeStamp
	})
	return ldRes
}
