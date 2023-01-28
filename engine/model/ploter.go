package model

import "github.com/go-echarts/go-echarts/charts"

type Plotter interface {
	Overlap(plotter Plotter)
	GetReactChart() *charts.RectChart
}
