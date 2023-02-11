package indicator

import (
	"fmt"
	"github.com/markcheno/go-talib"
	"testing"
)

func TestSma(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6}
	out := talib.Sma(in, 6)
	fmt.Println(out)
}
