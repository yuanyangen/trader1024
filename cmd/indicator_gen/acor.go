package main

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

func getAOCRIndictor() []model.Indicator {
	return []model.Indicator{
		indicator.NewAOCRIndicator(2),
		indicator.NewAOCRIndicator(5),
		indicator.NewAOCRIndicator(10),
		indicator.NewAOCRIndicator(20),
		indicator.NewAOCRIndicator(30),
		indicator.NewAOCRIndicator(40),
	}
}
