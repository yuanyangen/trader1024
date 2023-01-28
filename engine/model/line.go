package model

import (
	"sort"
)

type Line struct {
	*BaseLine
	Name string
}

type LineNode struct {
	Value     float64
	TimeStamp int64
}

func NewLine(t LineType, name string) *Line {
	return &Line{
		BaseLine: &BaseLine{
			Type: t,
		},
		Name: name,
	}
}

func (k *Line) GetByTsAndCount(ts int64, count int64) ([]*LineNode, error) {
	res, err := k.BaseLine.GetByTsAndCount(ts, count)
	if err != nil {
		return nil, err
	}
	newRes := make([]*LineNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*LineNode)
	}

	return newRes, nil
}

func (k *Line) AddData(ts int64, value float64) {
	node := &LineNode{
		TimeStamp: ts, Value: value,
	}
	k.BaseLine.AddData(ts, node)
}

func (k *Line) GetAllSortedData() []*LineNode {
	res := k.BaseLine.GetAllSortedData()
	newRes := make([]*LineNode, len(res))
	for i, v := range res {
		newRes[i] = v.(*LineNode)
	}

	sort.Slice(newRes, func(i, j int) bool {
		return newRes[i].TimeStamp < newRes[j].TimeStamp
	})
	return newRes
}
