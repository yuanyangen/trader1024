package eastmoney

import (
	"github.com/yuanyangen/trader1024/data/markets"
	"github.com/yuanyangen/trader1024/engine/model"
)

type EastMoneyContract struct {
	CNName          string
	Code            string
	VendorIdHandler func(code string, date string) string
}

var allSubjectsMap = map[string]*EastMoneyContract{}

var allSubjects = []*EastMoneyContract{
	{CNName: "不锈钢", Code: "113.ss", VendorIdHandler: getVendorId1},
	{CNName: "橡胶", Code: "113.ru", VendorIdHandler: getVendorId1},
	{CNName: "沥青", Code: "113.bu", VendorIdHandler: getVendorId1},
	{CNName: "沪金", Code: "113.au", VendorIdHandler: getVendorId1},
	{CNName: "沪铅", Code: "113.pb", VendorIdHandler: getVendorId1},
	{CNName: "沪铜", Code: "113.cu", VendorIdHandler: getVendorId1},
	{CNName: "沪铝", Code: "113.al", VendorIdHandler: getVendorId1},
	{CNName: "沪银", Code: "113.ag", VendorIdHandler: getVendorId1},
	{CNName: "沪锌", Code: "113.zn", VendorIdHandler: getVendorId1},
	{CNName: "沪锡", Code: "113.sn", VendorIdHandler: getVendorId1},
	{CNName: "沪镍", Code: "113.ni", VendorIdHandler: getVendorId1},
	{CNName: "热卷", Code: "113.hc", VendorIdHandler: getVendorId1},
	{CNName: "燃油", Code: "113.fu", VendorIdHandler: getVendorId1},
	{CNName: "纸浆", Code: "113.sp", VendorIdHandler: getVendorId1},
	{CNName: "线材", Code: "113.wr", VendorIdHandler: getVendorId1},
	{CNName: "螺纹钢", Code: "113.rb", VendorIdHandler: getVendorId1},
	{CNName: "IC合约", Code: "8.060130", VendorIdHandler: getVendorId2},
	{CNName: "IF合约", Code: "8.040130", VendorIdHandler: getVendorId2},
	{CNName: "IH合约", Code: "8.070130", VendorIdHandler: getVendorId2},
	{CNName: "IM合约", Code: "8.150130", VendorIdHandler: getVendorId2},
	{CNName: "TF合约", Code: "8.050130", VendorIdHandler: getVendorId2},
	{CNName: "TS合约", Code: "8.130130", VendorIdHandler: getVendorId2},
	{CNName: "T合约", Code: "8.110130", VendorIdHandler: getVendorId2},
	{CNName: "LPG", Code: "114.pg", VendorIdHandler: getVendorId3},
	{CNName: "PVC", Code: "114.v", VendorIdHandler: getVendorId3},
	{CNName: "乙二醇", Code: "114.eg", VendorIdHandler: getVendorId3},
	{CNName: "塑料", Code: "114.l", VendorIdHandler: getVendorId3},
	{CNName: "棕榈油", Code: "114.p", VendorIdHandler: getVendorId3},
	{CNName: "淀粉", Code: "114.cs", VendorIdHandler: getVendorId3},
	{CNName: "焦炭", Code: "114.j", VendorIdHandler: getVendorId3},
	{CNName: "焦煤", Code: "114.jm", VendorIdHandler: getVendorId3},
	{CNName: "玉米", Code: "114.c", VendorIdHandler: getVendorId3},
	{CNName: "生猪", Code: "114.lh", VendorIdHandler: getVendorId3},
	{CNName: "粳米", Code: "114.rr", VendorIdHandler: getVendorId3},
	{CNName: "纤维板", Code: "114.fb", VendorIdHandler: getVendorId3},
	{CNName: "聚丙烯", Code: "114.pp", VendorIdHandler: getVendorId3},
	{CNName: "胶合板", Code: "114.bb", VendorIdHandler: getVendorId3},
	{CNName: "苯乙烯", Code: "114.eb", VendorIdHandler: getVendorId3},
	{CNName: "豆一", Code: "114.a", VendorIdHandler: getVendorId3},
	{CNName: "豆二", Code: "114.b", VendorIdHandler: getVendorId3},
	{CNName: "豆油", Code: "114.y", VendorIdHandler: getVendorId3},
	{CNName: "豆粕", Code: "114.m", VendorIdHandler: getVendorId3},
	{CNName: "铁矿石", Code: "114.i", VendorIdHandler: getVendorId3},
	{CNName: "鸡蛋", Code: "114.jd", VendorIdHandler: getVendorId3},
	{CNName: "PTA", Code: "115.TA", VendorIdHandler: getVendorId4},
	{CNName: "动力煤", Code: "115.ZC", VendorIdHandler: getVendorId4},
	{CNName: "尿素", Code: "115.UR", VendorIdHandler: getVendorId4},
	{CNName: "强麦", Code: "115.WH", VendorIdHandler: getVendorId4},
	{CNName: "早籼稻", Code: "115.RI", VendorIdHandler: getVendorId4},
	{CNName: "晚籼稻", Code: "115.LR", VendorIdHandler: getVendorId4},
	{CNName: "普麦", Code: "115.PM", VendorIdHandler: getVendorId4},
	{CNName: "棉纱", Code: "115.CY", VendorIdHandler: getVendorId4},
	{CNName: "棉花", Code: "115.CF", VendorIdHandler: getVendorId4},
	{CNName: "玻璃", Code: "115.FG", VendorIdHandler: getVendorId4},
	{CNName: "甲醇", Code: "115.MA", VendorIdHandler: getVendorId4},
	{CNName: "白糖", Code: "115.SR", VendorIdHandler: getVendorId4},
	{CNName: "短纤", Code: "115.PF", VendorIdHandler: getVendorId4},
	{CNName: "硅铁", Code: "115.SF", VendorIdHandler: getVendorId4},
	{CNName: "粳稻", Code: "115.JR", VendorIdHandler: getVendorId4},
	{CNName: "红枣", Code: "115.CJ", VendorIdHandler: getVendorId4},
	{CNName: "纯碱", Code: "115.SA", VendorIdHandler: getVendorId4},
	{CNName: "花生", Code: "115.PK", VendorIdHandler: getVendorId4},
	{CNName: "苹果", Code: "115.AP", VendorIdHandler: getVendorId4},
	{CNName: "菜油", Code: "115.OI", VendorIdHandler: getVendorId4},
	{CNName: "菜籽", Code: "115.RS", VendorIdHandler: getVendorId4},
	{CNName: "菜粕", Code: "115.RM", VendorIdHandler: getVendorId4},
	{CNName: "锰硅", Code: "115.SM", VendorIdHandler: getVendorId4},
}

func init() {
	for _, v := range allSubjects {
		allSubjectsMap[v.CNName] = v
	}
}

func GetVendorCode(cnName, date string) string {
	for _, v := range allSubjects {
		if v.CNName == cnName {
			return v.VendorIdHandler(v.Code, date)
		}
	}
	return ""
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
func GetContractByCnName(cnName string, date string) *model.Contract {
	m := markets.GetSubjectByCnNam(cnName)
	if m == nil {
		panic("markte not support")
	}

	v, _ := allSubjectsMap[cnName]
	if v == nil {
		panic("markte not support by east money")
	}
	if v.Code == "" || v.VendorIdHandler == nil {
		panic("markte not support by east money")
	}

	return &model.Contract{
		Subject:      m,
		ContractTime: date,
	}
}
