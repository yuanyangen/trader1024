package talib

import "math"

func CustomKama(values []float64, period, fastEma, slowEma int, lastValue float64) []float64 {
	if len(values) <= period {
		panic("data error")
	}
	fastAlpha := 2.0 / (float64(fastEma) + 1.0)
	slowAlpha := 2 / (float64(slowEma) + 1.0)
	kamaV := float64(0.0)
	result := make([]float64, len(values))
	change := []float64{}
	for key := range values {
		if key == 0 {
			change = append(change, 0)
			continue

		}
		change = append(change, math.Abs(values[key]-values[key-1]))
		if key-period < 0 {
			continue
		}
		mon := math.Abs(values[key] - values[key-period])
		vol := 0.0
		for j := len(change) - period; j < len(change); j++ {
			vol += change[j]
		}
		er := 0.0
		if vol != 0 {
			er = mon / vol
		}
		alpha := math.Pow(er*(fastAlpha-slowAlpha)+slowAlpha, 2)
		if lastValue == 0 {
			sum := 0.0
			for j := key - period; j < key; j++ {
				sum += values[j]
			}
			lastValue = sum / float64(period)
		}
		kamaV = alpha*values[key] + (1-alpha)*lastValue

		result[key] = kamaV
	}
	return result
}
