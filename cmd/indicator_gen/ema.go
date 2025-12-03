package main

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

func getEMAIndictor() []model.Indicator {
	return []model.Indicator{
		indicator.NewEMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewEMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewEMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewEMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewEMAIndicator(30, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewEMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
	}
}
