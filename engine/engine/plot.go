package engine

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/model"
)

func DoPlot(p *charts.Page, bl *model.KLine) {
	kline := charts.NewKLine()
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: bl.Name},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	x, y := convertData(bl)
	kline.AddXAxis(x).AddYAxis(bl.Name, y)
	p.Add(kline)
	addDataToKline(kline, "avg", bl, func(dn *model.KLineNode) float64 {
		return (dn.High + dn.Low) / 2
	})
	addDataToKline(kline, "sma2", bl, func(dn *model.KLineNode) float64 {
		return dn.SmaData[2]
	})
	addDataToKline(kline, "sma5", bl, func(dn *model.KLineNode) float64 {
		return dn.SmaData[5]
	})
	addDataToKline(kline, "sma10", bl, func(dn *model.KLineNode) float64 {
		return dn.SmaData[10]
	})
	addDataToKline(kline, "sma20", bl, func(dn *model.KLineNode) float64 {
		return dn.SmaData[20]
	})
	addDataToKline(kline, "sma40", bl, func(dn *model.KLineNode) float64 {
		return dn.SmaData[40]
	})
}

func overlapLineToKline(kline *charts.Kline, name string, x []string, values []float64) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: name})
	line.AddXAxis(x).AddYAxis(name, values, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}

func convertData(bl *model.KLine) ([]string, [][4]float32) {
	kDatas := bl.GetAllSortedData()
	x := make([]string, len(kDatas))
	y := make([][4]float32, len(kDatas))
	for i, kn := range kDatas {
		x[i] = kn.TimeStampDesc
		y[i] = [4]float32{
			float32(kn.Open),
			float32(kn.Close),
			float32(kn.Low),
			float32(kn.High),
		}
	}
	return x, y
}

func addDataToKline(kline *charts.Kline, name string, bl *model.KLine, dataGetter func(dn *model.KLineNode) float64) {
	kDatas := bl.GetAllSortedData()
	x := make([]string, len(kDatas))
	y := make([]float64, len(kDatas))
	for i, kn := range kDatas {
		v := dataGetter(kn)
		if v > 0 {
			x[i] = kn.TimeStampDesc
			y[i] = dataGetter(kn)
		}

	}
	overlapLineToKline(kline, name, x, y)
}
