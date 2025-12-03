package indicator

import "github.com/yuanyangen/trader1024/engine/model"

func DataFromClose(in *model.KLineNode) float64  { return in.Close }
func DataFromHigh(in *model.KLineNode) float64   { return in.High }
func DataFromLow(in *model.KLineNode) float64    { return in.Low }
func DataFromOpen(in *model.KLineNode) float64   { return in.Open }
func DataFromAVG(in *model.KLineNode) float64    { return (in.Open + in.Close) / 2 }
func DataFromChange(in *model.KLineNode) float64 { return in.Close - in.Open }
