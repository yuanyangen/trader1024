package utils

import "math"

func SliceFloatGt(in []float64, i float64) bool {
	for _, v := range in {
		if v == 1 { //
			return false
		}
		if math.IsNaN(v) {
			return false
		}
		if v < i {
			return false
		}
	}
	return true
}

func SliceFloatLt(in []float64, i float64) bool {
	for _, v := range in {
		if v > i {
			return false
		}
	}
	return true
}
