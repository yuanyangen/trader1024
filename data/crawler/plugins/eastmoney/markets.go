package eastmoney

import (
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/model"
)

type EastMoneyMarket struct {
	CNName  string
	Code    string
	Handler func(code string, date string) string
}

var allSubjectsMap = map[string]*EastMoneyMarket{}

var allSubjects = []*EastMoneyMarket{
	{CNName: "不锈钢", Code: "113.ss", Handler: getVendorId1},
	{CNName: "橡胶", Code: "113.ru", Handler: getVendorId1},
	{CNName: "沥青", Code: "113.bu", Handler: getVendorId1},
	{CNName: "沪金", Code: "113.au", Handler: getVendorId1},
	{CNName: "沪铅", Code: "113.pb", Handler: getVendorId1},
	{CNName: "沪铜", Code: "113.cu", Handler: getVendorId1},
	{CNName: "沪铝", Code: "113.al", Handler: getVendorId1},
	{CNName: "沪银", Code: "113.ag", Handler: getVendorId1},
	{CNName: "沪锌", Code: "113.zn", Handler: getVendorId1},
	{CNName: "沪锡", Code: "113.sn", Handler: getVendorId1},
	{CNName: "沪镍", Code: "113.ni", Handler: getVendorId1},
	{CNName: "热卷", Code: "113.hc", Handler: getVendorId1},
	{CNName: "燃油", Code: "113.fu", Handler: getVendorId1},
	{CNName: "纸浆", Code: "113.sp", Handler: getVendorId1},
	{CNName: "线材", Code: "113.wr", Handler: getVendorId1},
	{CNName: "螺纹钢", Code: "113.rb", Handler: getVendorId1},
	{CNName: "IC合约", Code: "8.060130", Handler: getVendorId2},
	{CNName: "IF合约", Code: "8.040130", Handler: getVendorId2},
	{CNName: "IH合约", Code: "8.070130", Handler: getVendorId2},
	{CNName: "IM合约", Code: "8.150130", Handler: getVendorId2},
	{CNName: "TF合约", Code: "8.050130", Handler: getVendorId2},
	{CNName: "TS合约", Code: "8.130130", Handler: getVendorId2},
	{CNName: "T合约", Code: "8.110130", Handler: getVendorId2},
	{CNName: "LPG", Code: "114.pg", Handler: getVendorId3},
	{CNName: "PVC", Code: "114.v", Handler: getVendorId3},
	{CNName: "乙二醇", Code: "114.eg", Handler: getVendorId3},
	{CNName: "塑料", Code: "114.l", Handler: getVendorId3},
	{CNName: "棕榈油", Code: "114.p", Handler: getVendorId3},
	{CNName: "淀粉", Code: "114.cs", Handler: getVendorId3},
	{CNName: "焦炭", Code: "114.j", Handler: getVendorId3},
	{CNName: "焦煤", Code: "114.jm", Handler: getVendorId3},
	{CNName: "玉米", Code: "114.c", Handler: getVendorId3},
	{CNName: "生猪", Code: "114.lh", Handler: getVendorId3},
	{CNName: "粳米", Code: "114.rr", Handler: getVendorId3},
	{CNName: "纤维板", Code: "114.fb", Handler: getVendorId3},
	{CNName: "聚丙烯", Code: "114.pp", Handler: getVendorId3},
	{CNName: "胶合板", Code: "114.bb", Handler: getVendorId3},
	{CNName: "苯乙烯", Code: "114.eb", Handler: getVendorId3},
	{CNName: "豆一", Code: "114.a", Handler: getVendorId3},
	{CNName: "豆二", Code: "114.b", Handler: getVendorId3},
	{CNName: "豆油", Code: "114.y", Handler: getVendorId3},
	{CNName: "豆粕", Code: "114.m", Handler: getVendorId3},
	{CNName: "铁矿石", Code: "114.i", Handler: getVendorId3},
	{CNName: "鸡蛋", Code: "114.jd", Handler: getVendorId3},
	{CNName: "PTA", Code: "115.TA", Handler: getVendorId4},
	{CNName: "动力煤", Code: "115.ZC", Handler: getVendorId4},
	{CNName: "尿素", Code: "115.UR", Handler: getVendorId4},
	{CNName: "强麦", Code: "115.WH", Handler: getVendorId4},
	{CNName: "早籼稻", Code: "115.RI", Handler: getVendorId4},
	{CNName: "晚籼稻", Code: "115.LR", Handler: getVendorId4},
	{CNName: "普麦", Code: "115.PM", Handler: getVendorId4},
	{CNName: "棉纱", Code: "115.CY", Handler: getVendorId4},
	{CNName: "棉花", Code: "115.CF", Handler: getVendorId4},
	{CNName: "玻璃", Code: "115.FG", Handler: getVendorId4},
	{CNName: "甲醇", Code: "115.MA", Handler: getVendorId4},
	{CNName: "白糖", Code: "115.SR", Handler: getVendorId4},
	{CNName: "短纤", Code: "115.PF", Handler: getVendorId4},
	{CNName: "硅铁", Code: "115.SF", Handler: getVendorId4},
	{CNName: "粳稻", Code: "115.JR", Handler: getVendorId4},
	{CNName: "红枣", Code: "115.CJ", Handler: getVendorId4},
	{CNName: "纯碱", Code: "115.SA", Handler: getVendorId4},
	{CNName: "花生", Code: "115.PK", Handler: getVendorId4},
	{CNName: "苹果", Code: "115.AP", Handler: getVendorId4},
	{CNName: "菜油", Code: "115.OI", Handler: getVendorId4},
	{CNName: "菜籽", Code: "115.RS", Handler: getVendorId4},
	{CNName: "菜粕", Code: "115.RM", Handler: getVendorId4},
	{CNName: "锰硅", Code: "115.SM", Handler: getVendorId4},
}

func init() {
	for _, v := range allSubjects {
		allSubjectsMap[v.CNName] = v
	}
}

func (em *EastMoneyMarket) GetVendorCode(date string) string {
	return em.Handler(em.Code, date)
}

func getVendorId1(code string, date string) string {
	if date == "" {
		return code + "m"
	}
	return code + date
}
func getVendorId2(code string, date string) string {
	return ""
}
func getVendorId3(code string, date string) string {
	if date == "" {
		return code + "m"
	}
	return code + date
}
func getVendorId4(code string, date string) string {
	if date == "" {
		return code + "M"
	}
	return code + date
}

// date的格式是"230809"
func GetMarketByCnName(cnName string, date string) *model.Market {
	m := markets.GetMarketByCnNam(cnName)
	if m == nil {
		panic("markte not support")
	}

	v, _ := allSubjectsMap[cnName]
	if v == nil {
		panic("markte not support by east money")
	}
	if v.Code == "" || v.Handler == nil {
		panic("markte not support by east money")
	}

	return &model.Market{
		Subject:  m,
		VendorId: v.GetVendorCode(date),
	}
}
