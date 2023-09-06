package main

import (
	"fmt"
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/model"
	"sort"
	"time"
)

func main() {
	out := []*model.Subject{}
	for _, s := range markets.AllSubjects {
		if s.OnlineDay != "" {
			s.OnlineTime, _ = time.Parse("2006年1月2日", s.OnlineDay)
		}
		out = append(out, s)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Exchange+out[i].Name < out[j].Exchange+out[j].Name
	})
	for _, v := range out {
		fmt.Printf(`"%v": {Name: "%v", Exchange: "%v", Type: model.MarKetType_FUTURE, OnlineDay: "%v", FirstContract: "%v"},`+"\n", v.Name, v.Name, v.Exchange, v.OnlineDay, v.OnlineTime.Add(time.Hour*24*180).Format("20060102"))
	}
}
