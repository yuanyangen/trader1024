package account

import "github.com/shopspring/decimal"

type PositionType int

const PositionTypeLong = 1 //多头
const PositionShort = 2    //空头

type Position struct {
	Count decimal.Decimal //使用正表示多头， 使用负 表示空头，
}

func (p *Position) IsEmpty() bool {
	return p.Count.Equal(decimal.NewFromInt(0))
}
