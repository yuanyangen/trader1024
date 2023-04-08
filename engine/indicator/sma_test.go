package indicator

import (
	"fmt"
	"github.com/markcheno/go-talib"
	"testing"
)

func TestSma(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ou1 := talib.Sma(in, 3)
	//ou2 := talib.Ema(in, 3)
	ou2 := talib.Kama(in, 3)
	fmt.Println(ou1, "\n", ou2)
}
