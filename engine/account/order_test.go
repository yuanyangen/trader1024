package account

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestDecimal(t *testing.T) {
	a := decimal.NewFromFloat(2.0)
	b := a.Mul(decimal.NewFromFloat(-1))
	fmt.Println(b.Float64())
}
