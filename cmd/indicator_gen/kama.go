package main

//
//func getKAMAIndictor() []model.Indicator {
//	closeZDMA10 := indicator.NewZDMAIndicator(10, "close", indicator.DataFromClose)
//	openZDMA10 := indicator.NewZDMAIndicator(10, "open", indicator.DataFromOpen)
//	avgZDMA10 := indicator.NewZDMAIndicator(10, "avg", indicator.DataFromAVG)
//	return []model.Indicator{
//		//indicator.NewKAMAIndicator(2, indicator.DataSourceClose),
//		//indicator.NewKAMAIndicator(5, indicator.DataSourceClose),
//		//indicator.NewKAMAIndicator(10, indicator.DataSourceClose),
//		//indicator.NewKAMAIndicator(20, indicator.DataSourceClose),
//		//indicator.NewKAMAIndicator(30, indicator.DataSourceClose),
//		//indicator.NewKAMAIndicator(40, indicator.DataSourceClose),
//		//
//		//indicator.NewKAMAIndicator(2, indicator.DataSourceOpen),
//		//indicator.NewKAMAIndicator(5, indicator.DataSourceOpen),
//		//indicator.NewKAMAIndicator(10, indicator.DataSourceOpen),
//		//indicator.NewKAMAIndicator(20, indicator.DataSourceOpen),
//		//indicator.NewKAMAIndicator(30, indicator.DataSourceOpen),
//		//indicator.NewKAMAIndicator(40, indicator.DataSourceOpen),
//		//
//		//indicator.NewKAMAIndicator(2, indicator.DataSourceLow),
//		//indicator.NewKAMAIndicator(5, indicator.DataSourceLow),
//		//indicator.NewKAMAIndicator(10, indicator.DataSourceLow),
//		//indicator.NewKAMAIndicator(20, indicator.DataSourceLow),
//		//indicator.NewKAMAIndicator(30, indicator.DataSourceLow),
//		//indicator.NewKAMAIndicator(40, indicator.DataSourceLow),
//		//
//		//indicator.NewKAMAIndicator(2, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(5, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(10, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(20, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(30, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(40, indicator.DataSourceHigh),
//
//		indicator.NewKAMAIndicator(2, closeZDMA10.Name(), closeZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(2, openZDMA10.Name(), openZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(2, avgZDMA10.Name(), avgZDMA10.GetValueFromKNode),
//
//		indicator.NewKAMAIndicator(5, closeZDMA10.Name(), closeZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(5, openZDMA10.Name(), openZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(5, avgZDMA10.Name(), avgZDMA10.GetValueFromKNode),
//
//		//indicator.NewZDMAIndicator(10, "open", indicator.DataFromOpen),
//		//indicator.NewZDMAIndicator(10, "avg", indicator.DataFromAVG),
//		//
//		//indicator.NewKAMAIndicator(5, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(10, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(20, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(30, indicator.DataSourceHigh),
//		//indicator.NewKAMAIndicator(40, indicator.DataSourceHigh),
//	}
//}
//
//func getKAMAForZDMAIndictor() []model.Indicator {
//	closeZDMA10 := indicator.NewZDMAIndicator(20, "close", indicator.DataFromClose)
//	openZDMA10 := indicator.NewZDMAIndicator(20, "open", indicator.DataFromOpen)
//	avgZDMA10 := indicator.NewZDMAIndicator(20, "avg", indicator.DataFromAVG)
//	//closeZDMA10 := indicator.NewZDMAIndicator(20, "close", indicator.DataFromClose)
//	//openZDMA10 := indicator.NewZDMAIndicator(20, "open", indicator.DataFromOpen)
//	//avgZDMA10 := indicator.NewZDMAIndicator(20, "avg", indicator.DataFromAVG)
//	return []model.Indicator{
//		indicator.NewKAMAIndicator(2, closeZDMA10.Name(), closeZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(2, openZDMA10.Name(), openZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(2, avgZDMA10.Name(), avgZDMA10.GetValueFromKNode),
//
//		indicator.NewKAMAIndicator(5, closeZDMA10.Name(), closeZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(5, openZDMA10.Name(), openZDMA10.GetValueFromKNode),
//		indicator.NewKAMAIndicator(5, avgZDMA10.Name(), avgZDMA10.GetValueFromKNode),
//	}
//}
