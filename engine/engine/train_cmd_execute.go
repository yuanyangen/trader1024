package engine

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/yuanyangen/trader1024/engine/model"
	"github.com/yuanyangen/trader1024/engine/utils"
	"github.com/yuanyangen/trader1024/strategy/indicator"
	"os"
	"sync"
	"time"
)

type Train struct {
	Contract     *model.Contract
	kline        model.ContractIndicator
	trainResults map[int64]*TrainResult
	mu           sync.Mutex
	allDone      bool
}
type Report struct {
	MarketName string
	AllCount   int64

	ShortCount             int64
	ShortWinCountAfter1Day WinPercentCount
	ShortWinCountAfter5Day WinPercentCount
	//ShortWinCountAfter10Day WinPercentCount
	//ShortWinCountAfter15Day WinPercentCount
	ShortWinCountAfter20Day WinPercentCount

	LongCount             int64
	LongWinCountAfter1Day WinPercentCount
	LongWinCountAfter5Day WinPercentCount
	//LongWinCountAfter15Day WinPercentCount
	//LongWinCountAfter10Day WinPercentCount
	LongWinCountAfter20Day WinPercentCount
}

type WinPercentCount struct {
	LoseCount         int64
	Win0PercentCount  int64
	Win1PercentCount  int64
	Win3PercentCount  int64
	Win5PercentCount  int64
	Win7PercentCount  int64
	Win10PercentCount int64
}

type StrategyResult struct {
	RiseFallAfter1Day  float64
	RiseFallAfter5Day  float64
	RiseFallAfter20Day float64
	allDone            bool
}

type TrainResult struct {
	strategyReq         *ContractPortfolioReq
	strategyTrainResult *StrategyResult
}

func newTrain(contract *model.Contract, kline model.ContractIndicator, portfolioStrategy PortfolioStrategy) CmdExecutor {
	t := &Train{
		Contract:     contract,
		kline:        kline,
		trainResults: map[int64]*TrainResult{},
	}
	t.calcResultAndReportDaemon()
	return t
}

func (t *Train) ExecuteCmd(req *ContractPortfolioReq) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.trainResults[req.Ts] = &TrainResult{strategyReq: req, strategyTrainResult: &StrategyResult{}}
	t.calcResult()
	//t.Report()
}

func (t *Train) calcResultAndReportDaemon() {
	utils.AsyncRun(func() {
		for {
			//if !t.allDone {
			//	t.calcResult()
			//}
			t.Report()
			time.Sleep(time.Second)
		}
	})
}

func (t *Train) calcResult() {

	var err1, err5, err20 error
	var allDone = true
	for ts, result := range t.trainResults {
		if result.strategyTrainResult.allDone {
			continue
		}
		result.strategyTrainResult.RiseFallAfter1Day, err1 = t.calcDayResult(ts, 1)
		result.strategyTrainResult.RiseFallAfter5Day, err5 = t.calcDayResult(ts, 5)
		result.strategyTrainResult.RiseFallAfter20Day, err20 = t.calcDayResult(ts, 20)
		if err1 == nil && err5 == nil && err20 == nil {
			result.strategyTrainResult.allDone = true
		}
		if result.strategyTrainResult.allDone == false {
			allDone = false
		}
	}
	t.allDone = allDone
	return
}

// 计算 xx day 之后， 价格相对当前值的变化
func (t *Train) calcDayResult(ts int64, day int64) (float64, error) {
	kline, ok := t.kline.(*indicator.KLineIndicator)
	if !ok {
		panic("should not reach here")
	}
	nodes, err := kline.GetForwardByTsAndCount(ts, day+1)
	if err != nil {
		return 0, err
	}
	kNodes := model.NewKnodesFromAny(nodes)
	current := kNodes[0]
	xDayAfter := kNodes[len(kNodes)-1]
	rise := xDayAfter.GetValue() - current.GetValue()
	return rise, nil
}

func (t *Train) genReport(result []*TrainResult) *Report {
	r := &Report{
		//MarketName: t,
	}
	for _, v := range result {
		r.AllCount++
		if v.strategyReq == nil || v.strategyReq.StrategyResult == nil {
			fmt.Println("error")
			//panic("fasdfas")
			continue
		}
		if v.strategyReq.StrategyResult.Cmd == model.StrategyOutLong {
			r.LongCount++
			t.genOneDayReport(v.strategyReq.StrategyResult.Cmd, &r.LongWinCountAfter1Day, v.strategyTrainResult.RiseFallAfter1Day, v.strategyReq.StrategyResult.Price.InexactFloat64())
			t.genOneDayReport(v.strategyReq.StrategyResult.Cmd, &r.LongWinCountAfter5Day, v.strategyTrainResult.RiseFallAfter5Day, v.strategyReq.StrategyResult.Price.InexactFloat64())
			t.genOneDayReport(v.strategyReq.StrategyResult.Cmd, &r.LongWinCountAfter20Day, v.strategyTrainResult.RiseFallAfter20Day, v.strategyReq.StrategyResult.Price.InexactFloat64())

		} else if v.strategyReq.StrategyResult.Cmd == model.StrategyOutShort {
			r.ShortCount++
			t.genOneDayReport(v.strategyReq.StrategyResult.Cmd, &r.ShortWinCountAfter1Day, v.strategyTrainResult.RiseFallAfter1Day, v.strategyReq.StrategyResult.Price.InexactFloat64())
			t.genOneDayReport(v.strategyReq.StrategyResult.Cmd, &r.ShortWinCountAfter5Day, v.strategyTrainResult.RiseFallAfter5Day, v.strategyReq.StrategyResult.Price.InexactFloat64())
			t.genOneDayReport(v.strategyReq.StrategyResult.Cmd, &r.ShortWinCountAfter20Day, v.strategyTrainResult.RiseFallAfter20Day, v.strategyReq.StrategyResult.Price.InexactFloat64())
		} else {
			continue
			//panic("should not reach here")
		}
	}
	return r
}

func (t *Train) genOneDayReport(longShort model.StrategyOut, wpc *WinPercentCount, changePrice, buyPrice float64) {
	changePercent := changePrice / buyPrice
	if longShort == model.StrategyOutShort {
		changePercent *= -1
	}
	if changePercent <= 0 {
		wpc.LoseCount++
	}
	if changePercent > 0 {
		wpc.Win0PercentCount++
	}
	if changePercent > 0.01 {
		wpc.Win1PercentCount++
	}
	if changePercent > 0.03 {
		wpc.Win3PercentCount++
	}
	if changePercent > 0.05 {
		wpc.Win5PercentCount++
	}
	if changePercent > 0.07 {
		wpc.Win7PercentCount++
	}
	if changePercent > 0.10 {
		wpc.Win10PercentCount++
	}
}

func (t *Train) genReportAll() *Report {
	var result []*TrainResult
	for _, v := range t.trainResults {
		result = append(result, v)
	}
	return t.genReport(result)
}

func (t *Train) Report() {
	t.mu.Lock()
	defer t.mu.Unlock()
	report := t.genReportAll()
	fmt.Printf("=============================================================================================\n")
	ta := table.NewWriter()
	ta.SetOutputMirror(os.Stdout)
	ta.AppendHeader(table.Row{"指标名字", "lose", "0%", "1%", "3%", "5%", "7%", "10%"})
	ta.AppendRow(table.Row{
		"整体1天之后胜率",
		toString(report.LongWinCountAfter1Day.LoseCount+report.ShortWinCountAfter1Day.LoseCount, report.AllCount),
		toString(report.LongWinCountAfter1Day.Win0PercentCount+report.ShortWinCountAfter1Day.Win0PercentCount, report.AllCount),

		toString(report.LongWinCountAfter1Day.Win1PercentCount+report.ShortWinCountAfter1Day.Win1PercentCount, report.AllCount),
		toString(report.LongWinCountAfter1Day.Win3PercentCount+report.ShortWinCountAfter1Day.Win3PercentCount, report.AllCount),
		toString(report.LongWinCountAfter1Day.Win5PercentCount+report.ShortWinCountAfter1Day.Win5PercentCount, report.AllCount),
		toString(report.LongWinCountAfter1Day.Win7PercentCount+report.ShortWinCountAfter1Day.Win7PercentCount, report.AllCount),
		toString(report.LongWinCountAfter1Day.Win10PercentCount+report.ShortWinCountAfter1Day.Win10PercentCount, report.AllCount),
	})
	ta.AppendRow(table.Row{
		"整体5天之后胜率",
		toString(report.LongWinCountAfter5Day.LoseCount+report.ShortWinCountAfter5Day.LoseCount, report.AllCount),
		toString(report.LongWinCountAfter5Day.Win0PercentCount+report.ShortWinCountAfter5Day.Win0PercentCount, report.AllCount),
		toString(report.LongWinCountAfter5Day.Win1PercentCount+report.ShortWinCountAfter5Day.Win1PercentCount, report.AllCount),
		toString(report.LongWinCountAfter5Day.Win3PercentCount+report.ShortWinCountAfter5Day.Win3PercentCount, report.AllCount),
		toString(report.LongWinCountAfter5Day.Win5PercentCount+report.ShortWinCountAfter5Day.Win5PercentCount, report.AllCount),
		toString(report.LongWinCountAfter5Day.Win7PercentCount+report.ShortWinCountAfter5Day.Win7PercentCount, report.AllCount),
		toString(report.LongWinCountAfter5Day.Win10PercentCount+report.ShortWinCountAfter5Day.Win10PercentCount, report.AllCount),
	})
	ta.AppendRow(table.Row{
		"整体20天之后胜率",
		toString(report.LongWinCountAfter20Day.LoseCount+report.ShortWinCountAfter20Day.LoseCount, report.AllCount),
		toString(report.LongWinCountAfter20Day.Win0PercentCount+report.ShortWinCountAfter20Day.Win0PercentCount, report.AllCount),
		toString(report.LongWinCountAfter20Day.Win1PercentCount+report.ShortWinCountAfter20Day.Win1PercentCount, report.AllCount),
		toString(report.LongWinCountAfter20Day.Win3PercentCount+report.ShortWinCountAfter20Day.Win3PercentCount, report.AllCount),
		toString(report.LongWinCountAfter20Day.Win5PercentCount+report.ShortWinCountAfter20Day.Win5PercentCount, report.AllCount),
		toString(report.LongWinCountAfter20Day.Win7PercentCount+report.ShortWinCountAfter20Day.Win7PercentCount, report.AllCount),
		toString(report.LongWinCountAfter20Day.Win10PercentCount+report.ShortWinCountAfter20Day.Win10PercentCount, report.AllCount),
	})

	ta.AppendRow(table.Row{
		"做多1天之后胜率",
		toString(report.LongWinCountAfter1Day.LoseCount, report.LongCount),
		toString(report.LongWinCountAfter1Day.Win0PercentCount, report.LongCount),
		toString(report.LongWinCountAfter1Day.Win1PercentCount, report.LongCount),
		toString(report.LongWinCountAfter1Day.Win3PercentCount, report.LongCount),
		toString(report.LongWinCountAfter1Day.Win5PercentCount, report.LongCount),
		toString(report.LongWinCountAfter1Day.Win7PercentCount, report.LongCount),
		toString(report.LongWinCountAfter1Day.Win10PercentCount, report.LongCount),
	})
	ta.AppendRow(table.Row{
		"做多5天之后胜率",
		toString(report.LongWinCountAfter5Day.LoseCount, report.LongCount),
		toString(report.LongWinCountAfter5Day.Win0PercentCount, report.LongCount),
		toString(report.LongWinCountAfter5Day.Win1PercentCount, report.LongCount),
		toString(report.LongWinCountAfter5Day.Win3PercentCount, report.LongCount),
		toString(report.LongWinCountAfter5Day.Win5PercentCount, report.LongCount),
		toString(report.LongWinCountAfter5Day.Win7PercentCount, report.LongCount),
		toString(report.LongWinCountAfter5Day.Win10PercentCount, report.LongCount),
	})
	ta.AppendRow(table.Row{
		"做多20天之后胜率",
		toString(report.LongWinCountAfter20Day.LoseCount, report.LongCount),
		toString(report.LongWinCountAfter20Day.Win0PercentCount, report.LongCount),
		toString(report.LongWinCountAfter20Day.Win1PercentCount, report.LongCount),
		toString(report.LongWinCountAfter20Day.Win3PercentCount, report.LongCount),
		toString(report.LongWinCountAfter20Day.Win5PercentCount, report.LongCount),
		toString(report.LongWinCountAfter20Day.Win7PercentCount, report.LongCount),
		toString(report.LongWinCountAfter20Day.Win10PercentCount, report.LongCount),
	})

	ta.AppendRow(table.Row{
		"做空1天之后胜率",
		toString(report.ShortWinCountAfter1Day.LoseCount, report.ShortCount),
		toString(report.ShortWinCountAfter1Day.Win0PercentCount, report.ShortCount),

		toString(report.ShortWinCountAfter1Day.Win1PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter1Day.Win3PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter1Day.Win5PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter1Day.Win7PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter1Day.Win10PercentCount, report.ShortCount),
	})
	ta.AppendRow(table.Row{
		"做空5天之后胜率",
		toString(report.ShortWinCountAfter5Day.LoseCount, report.ShortCount),
		toString(report.ShortWinCountAfter5Day.Win0PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter5Day.Win1PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter5Day.Win3PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter5Day.Win5PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter5Day.Win7PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter5Day.Win10PercentCount, report.ShortCount),
	})
	ta.AppendRow(table.Row{
		"做空20天之后胜率",
		toString(report.ShortWinCountAfter20Day.LoseCount, report.ShortCount),
		toString(report.ShortWinCountAfter20Day.Win0PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter20Day.Win1PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter20Day.Win3PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter20Day.Win5PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter20Day.Win7PercentCount, report.ShortCount),
		toString(report.ShortWinCountAfter20Day.Win10PercentCount, report.ShortCount),
	})
	ta.Render()
}

func toString(win, total int64) string {
	return fmt.Sprintf("%v(%v/%v)", float64(win)/float64(total), win, total)
}
