package main

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

func getAtrIndictor() []model.Indicator {
	return []model.Indicator{
		indicator.NewATRIndicator(2),
		indicator.NewATRIndicator(5),
		indicator.NewATRIndicator(10),
		indicator.NewATRIndicator(20),
		indicator.NewATRIndicator(30),
		indicator.NewATRIndicator(40),
	}
}
