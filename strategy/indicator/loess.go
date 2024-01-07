package indicator

import (
	"fmt"
	"github.com/chewxy/stl/loess"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/yuanyangen/trader1024/engine/indicator_base"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
)

type LoessIndicator struct {
	*indicator_base.IndicatorCommon
	kline     model.ContractIndicator
	LoessLine *indicator_base.Line
	period    int64
}

func NewLoessIndicator(kline model.ContractIndicator, period int64) *LoessIndicator {
	Loess := &LoessIndicator{
		IndicatorCommon: indicator_base.NewIndicatorCommon(),
		period:          period,
		LoessLine:       indicator_base.NewLine(model.LineType_Day, fmt.Sprintf("Loess_%v", period)),
		kline:           kline,
	}
	kline.AddChildrenIndicator(Loess)
	return Loess
}

func (Loess *LoessIndicator) Name() string {
	return fmt.Sprintf("Loess_%v", Loess.period)
}

func (Loess *LoessIndicator) AddData(ts int64, node model.DataNode) {
	datas := Loess.kline.GetAllSortedData()
	priceData := []float64{}
	for _, v := range datas {
		priceData = append(priceData, v.GetValue())
	}
	if len(datas) < 12 {
		Loess.LoessLine.AddData(ts, priceData[len(priceData)-1])
		return
	}
	result, _ := loess.Smooth(priceData, 3, 5, loess.Linear)
	if len(result) > 0 {
		avg := result[len(result)-1]
		Loess.LoessLine.AddData(ts, avg)
	} else {
		Loess.LoessLine.AddData(ts, priceData[len(priceData)-1])
		return
	}
}
func (Loess *LoessIndicator) GetAllSortedData() []model.DataNode {
	return nil
}

func (Loess *LoessIndicator) GetByTs(ts int64) (model.DataNode, error) {
	if Loess.LoessLine == nil {
		panic("LoessLine error")
	}
	if Loess.period == 0 {
		panic("erPeriod empty")
	}
	return Loess.LoessLine.GetByTs(ts)
}
func (Loess *LoessIndicator) GetLastByTsAndCount(ts, period int64) ([]model.DataNode, error) {
	return nil, nil
}

func (Loess *LoessIndicator) DoPlot(kline *charts.Kline, ratioLine *charts.Line) {
	allData := Loess.LoessLine.GetAllSortedData()
	x := make([]string, len(allData))
	y := make([]float64, len(allData))
	for i, v := range allData {
		x[i] = utils.TsToDateString(v.GetTs())
		y[i] = v.GetValue()
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: Loess.Name()})
	line.AddXAxis(x).AddYAxis(Loess.Name(), y, charts.LineOpts{ConnectNulls: false})
	kline.Overlap(line)
}
