package butter

/*	Copyright (c) 2020, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Second order filter
type filter2 struct {
	s2, s1, a1, a2, b0, b1, b2 float64
}

func (f *filter2) Next(u float64) float64 {
	t := u - f.a1*f.s1 - f.a2*f.s2 // Direct Form 2
	y := f.b0*t + f.b1*f.s1 + f.b2*f.s2
	f.s2, f.s1 = f.s1, t // shift memory
	return y
}

func (f *filter2) NextS(u, y []float64) {
	n := len(u)
	if n > len(y) {
		n = len(y)
	}
	for i := 0; i < n; i++ {
		y[i] = f.Next(u[i])
	}
}

func (f *filter2) Reset(u, y float64) {
	// Internal states will be in arithmetic progression except HP
	// All divisions are safe
	hp := f.b1 == -2*f.b0 // HP?

	if hp {
		f.s1 = -y / (2 * f.b0)
		f.s2 = u - f.a1*f.s1
	} else {
		f.s1 = y / (f.b1 + 2*f.b0)
		f.s2 = u - (f.a1+2)*f.s1
	}

	if f.b2 < 0 || hp { // BP or HP?
		f.s2 /= 1 + f.a1 + f.a2
		f.s1 += f.s2
	} else {
		f.s2 /= f.a2 - 1
	}
}

// NewLowPass2 creates second order Low-Pass filter
func NewLowPass2(wc float64) Filter {
	if !valid(wc) {
		return nil
	}
	var f filter2
	wat := prewarp(wc) // wa * dt
	wat2 := wat * wat
	wat *= 2.82842712474619 // sqrt(8)

	a0 := wat2 + 4 + wat
	f.a1 = 2 * (wat2 - 4) / a0
	f.a2 = (wat2 + 4 - wat) / a0
	f.b0 = wat2 / a0
	f.b2 = f.b0
	f.b1 = 2 * f.b0
	f.s2, f.s1 = 0, 0
	return &f
}

// Filter represents a generic filter
type Filter interface {
	// Next runs filter one step with input u and returns output
	Next(float64) float64

	// NextS runs filter on a slice of input u and writes result in output y.
	// Input and output can be the same buffer.
	NextS(u, y []float64)

	// Reset sets internals of filter such that y will be the next output assuming
	// u is the next input approximately.
	Reset(u, y float64)
}

func valid(wc float64) bool {
	return .0001 < wc && wc < 3.1415 // reasonable bound
}

func valid2(wc, wd float64) bool {
	return .0001 < wc && wc+.0001 < wd && wd < 3.1415
}

// returns prewarped analog cut-off * dt, given desired cut-off * dt
func prewarp(wc float64) float64 {
	x := wc / 2
	x *= x
	return wc * (15 - x) / (15 - 6*x) // wa * dt
}
