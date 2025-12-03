package kalman

import (
	"github.com/konimarti/kalman"
	"github.com/konimarti/lti"
	"gonum.org/v1/gonum/mat"
)

func Filter(in []float64) []float64 {
	// define LTI system
	lti := lti.Discrete{
		Ad: mat.NewDense(1, 1, []float64{1}),
		Bd: mat.NewDense(1, 1, nil),
		C:  mat.NewDense(1, 1, []float64{1}),
		D:  mat.NewDense(1, 1, nil),
	}

	// system noise / process model covariance matrix ("Systemrauschen")
	Gd := mat.NewDense(1, 1, []float64{1})

	ctx := kalman.Context{
		// initial state
		X: mat.NewVecDense(1, []float64{in[0]}),
		// initial covariance matrix
		P: mat.NewDense(1, 1, []float64{0}),
	}

	// create ROSE filter
	gammaR := 9.0
	alphaR := 0.5
	alphaM := 0.9
	filter := kalman.NewRoseFilter(lti, Gd, gammaR, alphaR, alphaM)

	// no control
	u := mat.NewVecDense(1, nil)
	out := []float64{}
	for _, row := range in {
		// new measurement
		y := mat.NewVecDense(1, []float64{row})

		// apply filter
		filter.Apply(&ctx, y, u)
		// get corrected state vector
		state := filter.State()
		out = append(out, state.AtVec(0))
	}
	return out
}
