package talib

import (
	"fmt"
)

func main() {
	filter := NewSecondOrderGaussianLowPassFilter(10.0, 100.0)
	input := 10.5
	filteredValue := filter.Apply(input)
	fmt.Println(filteredValue)
}

// 二阶高斯低通滤波器结构
type SecondOrderGaussianLowPassFilter struct {
	prevOutput1 float64
	prevOutput2 float64
	prevInput1  float64
	prevInput2  float64
	a1          float64
	a2          float64
	b0          float64
	b1          float64
	b2          float64
}

// 初始化滤波器
func NewSecondOrderGaussianLowPassFilter(cutoffFrequency float64, samplingRate float64) *SecondOrderGaussianLowPassFilter {
	q := 0.707
	w0 := 2 * 3.141592653589793 * cutoffFrequency / samplingRate
	alpha := w0 / (2 * q)
	cw0 := w0 * w0
	a1 := -2 * cw0 / (alpha*alpha + 2*alpha*w0 + cw0)
	a2 := (alpha*alpha - 2*alpha*w0 + cw0) / (alpha*alpha + 2*alpha*w0 + cw0)
	b0 := alpha * alpha / (alpha*alpha + 2*alpha*w0 + cw0)
	b1 := 2 * b0
	b2 := b0
	return &SecondOrderGaussianLowPassFilter{
		prevOutput1: 0,
		prevOutput2: 0,
		prevInput1:  0,
		prevInput2:  0,
		a1:          a1,
		a2:          a2,
		b0:          b0,
		b1:          b1,
		b2:          b2,
	}
}

// 应用滤波器
func (f *SecondOrderGaussianLowPassFilter) Apply(input float64) float64 {
	output := f.b0*input + f.b1*f.prevInput1 + f.b2*f.prevInput2 + f.a1*f.prevOutput1 + f.a2*f.prevOutput2
	f.prevInput2 = f.prevInput1
	f.prevInput1 = input
	f.prevOutput2 = f.prevOutput1
	f.prevOutput1 = output
	return output
}
