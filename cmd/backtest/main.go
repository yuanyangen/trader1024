package main

import (
	"github.com/yuanyangen/trader1024/data/storage_client"
	"github.com/yuanyangen/trader1024/engine/account"
	"github.com/yuanyangen/trader1024/engine/engine"
	"github.com/yuanyangen/trader1024/engine/event"
	"github.com/yuanyangen/trader1024/strategy/portfolio"
	"github.com/yuanyangen/trader1024/strategy/strategy"
	"time"
)

func main() {

	e := engine.NewLiveExecuteEngine(
		event.NewBackTestDailyEventTrigger(1577841635, 1760927420),
		[]engine.PortfolioStrategy{
			portfolio.MultiStepPortfolio,
			portfolio.Evacuation,
		})
	e.RegisterStrategy(strategy.NewSingleSMAlineStrategyFactory)
	//e.RegisterContract("不锈钢", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沥青", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪金", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪铅", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪铜", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪铝", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪银", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪锌", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪锡", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("沪镍", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("热卷", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("燃油", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("纸浆", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("线材", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("螺纹钢", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("IC合约", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("IF合约", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("IH合约", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("IM合约", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("TF合约", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("TS合约", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("T合约", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("LPG", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("PVC", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("乙二醇", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("塑料", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("棕榈油", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("淀粉", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("焦炭", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("焦煤", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("玉米", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("生猪", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("粳米", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("纤维板", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("聚丙烯", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("胶合板", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("苯乙烯", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("豆一", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("豆二", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("豆油", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("豆粕", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("铁矿石", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("鸡蛋", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("PTA", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("动力煤", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("尿素", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("强麦", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("晚籼稻", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("普麦", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("棉纱", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("棉花", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("玻璃", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("甲醇", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("白糖", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("短纤", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("硅铁", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("粳稻", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("红枣", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("纯碱", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("花生", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("苹果", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("菜油", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("菜籽", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("菜粕", "", storage_client.SinaHttpStorage())
	//e.RegisterContract("锰硅", "", storage_client.SinaHttpStorage())

	e.RegisterContract("橡胶", []string{"202303"}, storage_client.SinaHttpStorage())
	//e.RegisterContract("玉米", []string{"202303", "202305"}, storage_client.SinaHttpStorage())

	account.RegisterAccount(account.NewAccount(1000000))
	e.Start()
	time.Sleep(time.Hour)
}
