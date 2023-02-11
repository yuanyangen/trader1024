package markets

import "github.com/yuanyangen/trader1024/engine/model"

var allMarkets = map[string]*model.Market{
	"ICM": &model.Market{
		Type:     model.MarKetType_FUTURE,
		Name:     "中证500股指主力",
		MarketId: "ICM",
	},
}

func GetMarketById(id string) *model.Market {
	v, _ := allMarkets[id]
	return v
}
