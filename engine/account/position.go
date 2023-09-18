package account

import (
	"fmt"
	"github.com/shopspring/decimal"
	"sort"
	"sync"
	"time"
)

type PositionType int

const PositionTypeLong PositionType = 1  //多头
const PositionTypeShort PositionType = 2 //空头

func (pt PositionType) String() string {
	if pt == 1 {
		return "多"
	}
	if pt == 2 {
		return "空"
	}
	return "unknown"
}

type ContractPosition struct {
	mu       sync.Mutex
	MarketId string
	Count    decimal.Decimal //使用正表示多头， 使用负 表示空头，
	Details  []*PositionPair
}

func (p *ContractPosition) IsEmpty() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.Count.Equal(decimal.Zero)
}

// 默认直接使用最老的position进行操作
func (p *ContractPosition) ProcessOrder(order *Order) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.processTotalCount(order)

	lastCount := order.Count
	for _, pp := range p.Details {
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
func (p *ContractPosition) processTotalCount(order *Order) {
	if order.OrderType == OrderTypeSell {
		p.Count = p.Count.Sub(order.Count)
	} else if order.OrderType == OrderTypeBuy {
		p.Count = p.Count.Add(order.Count)
	} else {
		panic("should not reach here")
	}
}
func (p *ContractPosition) newPositionPair(lastCount decimal.Decimal, order *Order) {
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
func (p *ContractPosition) splitPosition(pp *PositionPair, lastCount decimal.Decimal, order *Order) decimal.Decimal {
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
		pp.genGain()
		lastCount = decimal.Zero
	} else {
		p2.Count = p1.Count
		p2.Price = order.Price
		pp.Clear = true
		pp.genGain()
		lastCount = lastCount.Sub(p1.Count)

	}
	return lastCount
}

func (p *ContractPosition) addPositionPair(pp *PositionPair) {
	p.Details = append(p.Details, pp)
	sort.Slice(p.Details, func(i, j int) bool {
		return p.Details[i].CreateTimeStamp < p.Details[j].CreateTimeStamp
	})
}

func (p *ContractPosition) Report() {

	pairs := []*PositionPair{}
	for _, pp := range p.Details {
		pairs = append(pairs, pp)
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Gain.LessThan(pairs[j].Gain)
	})
	longCount := 0
	shortCount := 0
	longGain := 0.0
	shortGain := 0.0
	winCount := 0
	for _, positionPair := range pairs {
		positionPair.Report()
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

	fmt.Printf("total=%v win=%v win_ratio=%v long_gain=%v short_gain=%v gain_ratio=%v\n",
		longCount+shortCount,
		winCount,
		float64(winCount)/float64(longCount+shortCount),
		longGain,
		shortGain,
		longGain+shortGain,
	)

}

type PositionPair struct {
	Type            PositionType
	Buy             *Position
	Sell            *Position
	Clear           bool // 是否买卖平衡
	Gain            decimal.Decimal
	CreateTimeStamp int64
}

func (pp *PositionPair) genGain() {
	if !pp.Clear {
		return
	}
	pp.Gain = pp.Sell.Price.Sub(pp.Buy.Price).Mul(pp.Buy.Count)
}

func (pp *PositionPair) Report() {
	fmt.Printf("order_type=%v ", pp.Type.String())
	if pp.Buy != nil {
		fmt.Printf("buy_time=%v buy_price=%v buy_count=%v ",
			time.Unix(pp.Buy.OrderInfo.CreateTimeStamp, 0).Format("2006-01-02 15:04:05"),
			pp.Buy.Price.String(),
			pp.Buy.Count.String(),
		)
	}
	if pp.Sell != nil {
		fmt.Printf("sell_time=%v sell_price=%v sell_count=%v ",
			time.Unix(pp.Sell.OrderInfo.CreateTimeStamp, 0).Format("2006-01-02 15:04:05"),
			pp.Sell.Price.String(),
			pp.Sell.Count.String(),
		)
	}
	if pp.Sell != nil && pp.Buy != nil {
		fmt.Printf("gain=%v", pp.Gain.String())
	}
	fmt.Printf("\n")
}

type Position struct {
	Count     decimal.Decimal //使用正表示买， 使用负卖
	Price     decimal.Decimal
	OrderInfo *Order
}
