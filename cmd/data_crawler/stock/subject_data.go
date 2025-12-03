package main

import "github.com/yuanyangen/trader1024/engine/model"

var AllSubjects = map[string]*model.SubjectDO{
	"比亚迪": {CNName: "比亚迪", UniqueCode: "002594.SZ", Exchange: "深圳交易所", Type: model.MarKetType_STOCK, OnlineDay: "2011年6月30日", TypeLevel1: []string{"制造业"}, TypeLevel2: []string{"新能源汽车", "电池制造", "​​汽车整车"}},
}
