package model

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/bytedance/sonic"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/logs"
)

type Crawler interface {
	CrawlDaily(ctx context.Context, market *TradeObject) ([]*KLineNode, error)
	CrawlMinute(ctx context.Context, market *TradeObject) ([]*KLineNode, error)
	CrawlWeekly(ctx context.Context, market *TradeObject) ([]*KLineNode, error)
}

type KLine struct {
	Indicators    []Indicator
	allCachedData []*KLineNode
	*BaseLine
}

type KLineNode struct {
	UniqueCode    string
	Type          LineType
	TimeStamp     int64
	TimeStampDesc string
	High          float64
	Low           float64
	Open          float64
	Close         float64
	Volume        float64
	Turnover      float64 // 成交额
	Swing         float64 // 振幅
	Increase      float64 // 涨跌幅 ??
	IncreaseMount float64 // 涨跌额 ??
	TurnoverRate  float64 //换手率

	AtrData  map[int64]float64
	AOCRData map[int64]float64

	JMAData      map[string]float64
	EMAData      map[string]float64
	SMAData      map[string]float64
	RawKlineData map[string]float64
	KAMAData     map[string]float64
	SlopData     map[string]float64
	ZDMAData     map[string]float64 //https://mp.weixin.qq.com/s?__biz=MzIxNzUyNTI4MA==&mid=2247484438&idx=1&sn=3f2c8d47efcc30ed734dbd6c62e2efe9&chksm=97f93959a08eb04ff37ed759372ab10d6f24931ab05aeb125397e10e7c92036fac6b6829ff4f&scene=21#wechat_redirect
	ZDKAMAData   map[string]float64 //https://mp.weixin.qq.com/s?__biz=MzIxNzUyNTI4MA==&mid=2247484438&idx=1&sn=3f2c8d47efcc30ed734dbd6c62e2efe9&chksm=97f93959a08eb04ff37ed759372ab10d6f24931ab05aeb125397e10e7c92036fac6b6829ff4f&scene=21#wechat_redirect
	LPFData      map[string]float64 //https://mp.weixin.qq.com/s?__biz=MzIxNzUyNTI4MA==&mid=2247484438&idx=1&sn=3f2c8d47efcc30ed734dbd6c62e2efe9&chksm=97f93959a08eb04ff37ed759372ab10d6f24931ab05aeb125397e10e7c92036fac6b6829ff4f&scene=21#wechat_redirect
	LRData       map[string]float64
	MomentumData map[string]float64
}

func NewDataNodeNewFromAny(val any) *KLineNode {
	if val == nil {
		return nil
	}
	r, _ := val.(*KLineNode)
	return r
}

func NewDataNodeFromJson(val []byte) *KLineNode {
	knode := &KLineNode{}
	err := json.Unmarshal(val, &knode)
	if err != nil {
		panic(fmt.Sprintf("data_crawler in db unmarshal error %v", err))
	}
	return knode
}

func (dn *KLineNode) GetTs() int64 {
	return dn.TimeStamp
}
func (dn *KLineNode) String() string {
	r, _ := sonic.MarshalString(dn)
	return r
}

func NewKLine(name string, t LineType) *KLine {
	bl := &KLine{
		BaseLine: NewBaseLine(name, t),
	}
	return bl
}

func (bl *KLine) GetNodeByTs(ts int64) (*KLineNode, error) {
	return bl.convertAnyToLineNode(bl.GetByTs(ts))
}

// get last one
func (bl *KLine) GetLastNodeByTs(ts int64) (*KLineNode, error) {
	return bl.convertAnyToLineNode(bl.GetLastByTs(ts))
}

func (bl *KLine) GetAllLastNodeByTs(ts int64) []*KLineNode {
	allData := bl.GetAllSortedData()
	out := []*KLineNode{}
	for _, v := range allData {
		if v.TimeStamp <= ts {
			out = append(out, v)
		}
	}
	return out
}

func (bl *KLine) GetLastNodeByTsAndCount(ts int64, count int64) ([]*KLineNode, error) {
	return bl.convertAnyToLineNodes(bl.GetLastByTsAndCount(ts, count))
}

func (bl *KLine) GetForwardNodeByTsAndCount(ts int64, count int64) ([]*KLineNode, error) {
	return bl.convertAnyToLineNodes(bl.GetForwardByTsAndCount(ts, count))
}

func (bl *KLine) AddNodeData(ts int64, node *KLineNode) {
	bl.AddData(ts, &LineNode{DataNode: node, TimeStamp: ts})
}

func (bl *KLine) GetAllNodeData() []*KLineNode {
	allData := bl.GetAllData()
	out, _ := bl.convertAnyToLineNodes(allData, nil)
	return out
}

func (bl *KLine) GetAllSortedData() []*KLineNode {
	ldRes := bl.GetAllNodeData()
	sort.Slice(ldRes, func(i, j int) bool {
		return ldRes[i].TimeStamp < ldRes[j].TimeStamp
	})
	return ldRes
}

func (bl *KLine) convertAnyToLineNode(in *LineNode, err error) (*KLineNode, error) {
	if err != nil {
		return nil, err
	}
	if in != nil && in.DataNode != nil {
		node, ok := in.DataNode.(*KLineNode)
		if !ok {
			return nil, fmt.Errorf("data not lineNode")
		}
		return node, nil
	}
	return nil, fmt.Errorf("no line node")
}

func (bl *KLine) convertAnyToLineNodes(in []*LineNode, err error) ([]*KLineNode, error) {
	out := []*KLineNode{}
	for _, vv := range in {
		if vv != nil && vv.DataNode != nil {
			node, ok := vv.DataNode.(*KLineNode)
			if !ok {
				return nil, fmt.Errorf("data not lineNode")
			}
			out = append(out, node)
		}
	}
	return out, err
}

func (bl *KLine) DoPlot(p *charts.Page) {
	kline := charts.NewKLine()
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: bl.Name},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	x, y := convertData(bl)
	kline.AddXAxis(x).AddYAxis(bl.Name, y)
	p.Add(kline)
	for _, indi := range bl.Indicators {
		addDataToKline(kline, bl, indi)
	}
}

func overlapLineToKline(kline *charts.Kline, name string, x []string, values []float64) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: name}, charts.YAxisOpts{Scale: true, GridIndex: -100})
	line.AddXAxis(x).AddYAxis(name, values, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}

func convertData(bl *KLine) ([]string, [][4]float32) {
	kDatas := bl.GetAllSortedData()
	x := make([]string, len(kDatas))
	y := make([][4]float32, len(kDatas))
	allOpen := make([]float32, 0)
	for i, kn := range kDatas {
		x[i] = kn.TimeStampDesc
		y[i] = [4]float32{
			float32(kn.Open),
			float32(kn.Close),
			float32(kn.Low),
			float32(kn.High),
		}
		allOpen = append(allOpen, float32(kn.Open))
	}
	r, _ := sonic.MarshalString(allOpen)
	logs.Info("%v", r)
	return x, y
}

func addDataToKline(kline *charts.Kline, bl *KLine, ind Indicator) {
	kDatas := bl.GetAllSortedData()
	x := make([]string, len(kDatas))
	y := make([]float64, len(kDatas))
	for i, kn := range kDatas {
		v := ind.GetValueFromKNode(kn)
		//if v > 0 {
		x[i] = kn.TimeStampDesc
		y[i] = v
		//}

	}
	overlapLineToKline(kline, ind.Name(), x, y)
}
