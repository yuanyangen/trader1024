package model

import (
	"fmt"
	"sort"
	"sync"

	"github.com/shopspring/decimal"
	"github.com/yuanyangen/trader1024/engine/logs"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type PositionType int

const PositionTypeLong PositionType = 1      //多头
const PositionTypeShort PositionType = 2     //空头
const PositionTypeClear PositionType = 3     //清仓
const PositionTypeClearLast PositionType = 4 // 中间状态，清除最近一次加仓位

var mu sync.Mutex

type ContractPositions struct {
	mu                    sync.Mutex
	Contract              *TradeObject
	Count                 decimal.Decimal //使用正表示多头， 使用负 表示空头，
	Details               []*PositionPair
	EndTimeToPositionPair map[int64]*PositionPair
}

func (p *ContractPositions) IsEmpty() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.Count.Equal(decimal.Zero)
}

// 默认直接使用最新的position进行操作
func (p *ContractPositions) ProcessOrder(order *Order) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if order.OrderType == OrderTypeSell {
		p.Count = p.Count.Sub(order.Count)
	} else if order.OrderType == OrderTypeBuy {
		p.Count = p.Count.Add(order.Count)
	} else {
		panic("should not reach here")
	}

	lastCount := order.Count

	tmpPositions := make([]*PositionPair, len(p.Details))
	for i, v := range p.Details {
		tmpPositions[i] = v
	}
	sort.Slice(tmpPositions, func(i, j int) bool {
		return tmpPositions[i].CreateTimeStamp > tmpPositions[j].CreateTimeStamp
	})

	for _, pp := range tmpPositions {
		if pp.Clear {
			continue
		}
		if order.OrderType == OrderTypeBuy && pp.Type == PositionTypeLong {
			continue
		}
		if order.OrderType == OrderTypeSell && pp.Type == PositionTypeShort {
			continue
		}
		lastCount = p.splitPosition(pp, lastCount, order)
		if lastCount.Equal(decimal.Zero) {
			break
		}
	}
	if lastCount.GreaterThan(decimal.Zero) {
		p.newPositionPair(lastCount, order)
	}
}
func (p *ContractPositions) newPositionPair(lastCount decimal.Decimal, order *Order) {
	newP1 := &Position{
		Count:     lastCount,
		Price:     order.Price,
		OrderInfo: order,
	}
	newPP := &PositionPair{
		CreateTimeStamp: order.CreateTimeStamp,
	}
	if order.OrderType == OrderTypeBuy {
		newPP.Type = PositionTypeLong
		newPP.Buy = newP1
	} else {
		newPP.Type = PositionTypeShort
		newPP.Sell = newP1
	}
	p.addPositionPair(newPP)
}

// 从一个确定的paire中，分割一个特定的头寸出来
func (p *ContractPositions) splitPosition(pp *PositionPair, lastCount decimal.Decimal, order *Order) decimal.Decimal {
	var p1 *Position
	var p2 = &Position{
		OrderInfo: order,
		Price:     order.Price,
	}
	if pp.Type == PositionTypeLong {
		p1 = pp.Buy
		pp.Sell = p2
	} else {
		p1 = pp.Sell
		pp.Buy = p2
	}

	if p1.Count.GreaterThan(lastCount) {
		p1LastCount := p1.Count.Sub(lastCount)
		p2.Count = lastCount
		pp.Clear = true
		pp.EndTimeStamp = order.CreateTimeStamp
		pp.genGain()
		newP1 := &Position{
			Count:     p1LastCount,
			Price:     p1.Price,
			OrderInfo: p1.OrderInfo,
		}
		newPP := &PositionPair{
			Type:            pp.Type,
			CreateTimeStamp: p1.OrderInfo.CreateTimeStamp,
		}
		if pp.Type == PositionTypeLong {
			newPP.Buy = newP1
		} else {
			newPP.Sell = newP1
		}
		p.addPositionPair(newPP)
		lastCount = decimal.Zero
	} else if pp.Buy.Count.Equal(lastCount) {
		p2.Count = lastCount
		p2.Price = order.Price
		pp.Clear = true
		pp.EndTimeStamp = order.CreateTimeStamp
		pp.genGain()
		lastCount = decimal.Zero
	} else {
		p2.Count = p1.Count
		p2.Price = order.Price
		pp.Clear = true
		pp.EndTimeStamp = order.CreateTimeStamp
		pp.genGain()
		lastCount = lastCount.Sub(p1.Count)
	}
	if pp.Clear {
		logs.Info("position_result: %v\n", pp.String())
	}
	return lastCount
}

func (p *ContractPositions) addPositionPair(pp *PositionPair) {
	p.Details = append(p.Details, pp)
	sort.Slice(p.Details, func(i, j int) bool {
		return p.Details[i].CreateTimeStamp < p.Details[j].CreateTimeStamp
	})
}

func (p *ContractPositions) GetLastPair() *PositionPair {
	pairs := []*PositionPair{}
	for _, pp := range p.Details {
		pairs = append(pairs, pp)
	}
	if len(pairs) == 0 {
		return nil
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].CreateTimeStamp < pairs[j].CreateTimeStamp
	})
	return pairs[len(pairs)-1]
}

func (p *ContractPositions) ReportToCmd() {
	pairs := []*PositionPair{}
	for _, pp := range p.Details {
		pairs = append(pairs, pp)
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Gain.LessThan(pairs[j].Gain)
		//return pairs[i].Gain.LessThan(pairs[j].Gain)
	})
	longCount := 0
	shortCount := 0
	longGain := 0.0
	shortGain := 0.0
	winCount := 0
	logs.Info("\n\n######################## start ########################\n")
	for _, positionPair := range pairs {
		logs.Info("%v:%v\n", p.Contract.Id(), positionPair.String())
		if positionPair.Clear {
			if positionPair.Type == PositionTypeLong {
				longCount++
				g, _ := positionPair.Gain.Float64()
				longGain += g
			}
			if positionPair.Type == PositionTypeShort {
				shortCount++
				g, _ := positionPair.Gain.Float64()
				shortGain += g
			}
			if positionPair.Gain.GreaterThan(decimal.Zero) {
				winCount++
			}
		}
	}

	logs.Info("total=%v win=%v win_ratio=%v long_gain=%v short_gain=%v gain_ratio=%v\n",
		longCount+shortCount,
		winCount,
		float64(winCount)/float64(longCount+shortCount),
		longGain,
		shortGain,
		longGain+shortGain,
	)

}

func (pp *PositionPair) genGain() {
	if !pp.Clear {
		return
	}
	pp.Gain = pp.Sell.Price.Sub(pp.Buy.Price).Mul(pp.Buy.Count)
}

type Position struct {
	Count     decimal.Decimal //使用正表示买， 使用负卖
	Price     decimal.Decimal
	OrderInfo *Order
}

type PositionPair struct {
	Type            PositionType
	Buy             *Position
	Sell            *Position
	Clear           bool // 是否买卖平衡
	Gain            decimal.Decimal
	CreateTimeStamp int64
	EndTimeStamp    int64
}

func (pt PositionType) String() string {
	if pt == 1 {
		return "多"
	}
	if pt == 2 {
		return "空"
	}
	return "unknown"
}

func (pt *PositionPair) String() string {
	buyPrice := 0.0
	buyTime := ""
	sellTime := ""
	sellPrice := 0.0
	if pt.Buy != nil {
		buyTime = utils.TsToDateString(pt.Buy.OrderInfo.CreateTimeStamp)
		//buyReason = pt.Buy.OrderInfo.Reason
		buyPrice = pt.Buy.Price.InexactFloat64()
	}
	if pt.Sell != nil {
		sellTime = utils.TsToDateString(pt.Sell.OrderInfo.CreateTimeStamp)
		//sellReason = pt.Sell.OrderInfo.Reason
		sellPrice = pt.Sell.Price.InexactFloat64()
	}
	win := "unknown"
	count := "Empty"
	Gain := "Empty"
	if pt.Clear {
		if pt.Gain.LessThan(decimal.Zero) {
			win = "false"
		} else {
			win = "true"
		}
		count = pt.Buy.Count.String()
		Gain = pt.Gain.String()
	}
	return fmt.Sprintf("win:%v gain:%v detail: %v:%v buy:%v %.3f:  sell:%v %.3f",
		Gain, win,
		pt.Type.String(), count,
		buyTime, buyPrice,
		sellTime, sellPrice,
	)
}
