package talib

import (
	"fmt"
	"testing"
)

func TestKama(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ou1 := Kama(in, 3, 2, 30)
	ou2 := CustomKama(in, 3, 2, 30)
	fmt.Println(in)
	fmt.Println(ou1, "\n", ou2)
}
