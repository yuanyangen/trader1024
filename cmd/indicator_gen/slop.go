package main

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

func getSlopIndictor() []model.Indicator {
	return []model.Indicator{
		indicator.NewSlopIndicator(2, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		indicator.NewSlopIndicator(2, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
		indicator.NewSlopIndicator(2, indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceAVG))),
	}
}
