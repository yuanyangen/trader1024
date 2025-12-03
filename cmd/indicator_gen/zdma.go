package main

import (
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/plugins/indicator"
)

func getZDMAIndictor() []model.Indicator {
	return []model.Indicator{
		indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
		indicator.NewZDMAIndicator(2, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

		indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
		indicator.NewZDMAIndicator(5, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

		indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
		indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

		indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
		indicator.NewZDMAIndicator(20, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

		indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
		indicator.NewZDMAIndicator(40, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

		indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceClose)),
		indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceOpen)),
		indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceAVG)),
		indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceHigh)),
		indicator.NewZDMAIndicator(60, indicator.NewRawKlineIndicator(indicator.DataSourceLow)),

		indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceClose))),
		indicator.NewZDMAIndicator(10, indicator.NewZDMAIndicator(10, indicator.NewRawKlineIndicator(indicator.DataSourceOpen))),
	}
}

//func getZDKAMAIndictor() []model.Indicator {
//	return []model.Indicator{
//		indicator.NewZDKAMAIndicator(5, "close", indicator.DataFromClose),
//		indicator.NewZDKAMAIndicator(5, "open", indicator.DataFromOpen),
//		indicator.NewZDKAMAIndicator(5, "avg", indicator.DataFromAVG),
//		indicator.NewZDKAMAIndicator(10, "close", indicator.DataFromClose),
//		indicator.NewZDKAMAIndicator(10, "open", indicator.DataFromOpen),
//		indicator.NewZDKAMAIndicator(10, "avg", indicator.DataFromAVG),
//		indicator.NewZDKAMAIndicator(20, "close", indicator.DataFromClose),
//		indicator.NewZDKAMAIndicator(20, "open", indicator.DataFromOpen),
//		indicator.NewZDKAMAIndicator(20, "avg", indicator.DataFromAVG),
//		indicator.NewZDKAMAIndicator(40, "close", indicator.DataFromClose),
//		indicator.NewZDKAMAIndicator(40, "open", indicator.DataFromOpen),
//		indicator.NewZDKAMAIndicator(40, "avg", indicator.DataFromAVG),
//		indicator.NewZDKAMAIndicator(80, "close", indicator.DataFromClose),
//		indicator.NewZDKAMAIndicator(80, "open", indicator.DataFromOpen),
//		indicator.NewZDKAMAIndicator(80, "avg", indicator.DataFromAVG),
//	}
//}
//
//func getZDMA2Indictor() []model.Indicator {
//	return []model.Indicator{
//		//indicator.NewZDMAIndicator(5, "close", indicator.DataFromClose),
//		//indicator.NewZDMAIndicator(5, "open", indicator.DataFromOpen),
//		indicator.NewZDMAIndicator(10, "zdma2_close", indicator.NewZDMAIndicator(10, "close", indicator.DataFromClose).GetValueFromKNode),
//		indicator.NewZDMAIndicator(10, "zdma2_open", indicator.NewZDMAIndicator(10, "open", indicator.DataFromOpen).GetValueFromKNode),
//
//		//indicator.NewZDMAIndicator(20, "close", indicator.DataFromClose),
//		//indicator.NewZDMAIndicator(20, "open", indicator.DataFromOpen),
//		//indicator.NewZDMAIndicator(40, "close", indicator.DataFromClose),
//		//indicator.NewZDMAIndicator(40, "open", indicator.DataFromOpen),
//		//indicator.NewZDMAIndicator(80, "close", indicator.DataFromClose),
//		//indicator.NewZDMAIndicator(80, "open", indicator.DataFromOpen),
//	}
//}
