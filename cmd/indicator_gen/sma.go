package main

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

func getSmaIndictor() []model.Indicator {
	return []model.Indicator{
		indicator.NewSMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewSMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewSMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewSMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewSMAIndicator(30, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewSMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
	}
}
