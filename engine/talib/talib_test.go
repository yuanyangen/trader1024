package talib

import (
	"fmt"
	"testing"
)

func TestMA(t *testing.T) {
	in := []float64{1, 3, 4, 5, 6, 6, 6, 7, 3, 7, 8}
	fmt.Println(Sma(in, 5))
}
