package indicator

import (
	"fmt"
	"testing"
)

func TestLr(t *testing.T) {

	out := LinearReg([]float64{1}, 2)
	fmt.Println(out)
}
