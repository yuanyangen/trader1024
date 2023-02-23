package markets

import "github.com/yuanyangen/trader1024/engine/model"

var allMarkets = map[string]*model.Market{
	"snm":    {Type: model.MarKetType_FUTURE, Code: "snm", Name: "沪锡主力", SecId: "113.snm", Exchange: "上期所"},
	"APM":    {Type: model.MarKetType_FUTURE, Code: "APM", Name: "苹果主力", SecId: "115.APM", Exchange: "郑商所"},
	"070130": {Type: model.MarKetType_FUTURE, Code: "070130", Name: "IH主力合约", SecId: "8.070130", Exchange: "中金所"},
	"040130": {Type: model.MarKetType_FUTURE, Code: "040130", Name: "IF主力合约", SecId: "8.040130", Exchange: "中金所"},
	"URM":    {Type: model.MarKetType_FUTURE, Code: "URM", Name: "尿素主力", SecId: "115.URM", Exchange: "郑商所"},
	"vm":     {Type: model.MarKetType_FUTURE, Code: "vm", Name: "PVC主力", SecId: "114.vm", Exchange: "大商所"},
	"alm":    {Type: model.MarKetType_FUTURE, Code: "alm", Name: "沪铝主力", SecId: "113.alm", Exchange: "上期所"},
	"znm":    {Type: model.MarKetType_FUTURE, Code: "znm", Name: "沪锌主力", SecId: "113.znm", Exchange: "上期所"},
	"fum":    {Type: model.MarKetType_FUTURE, Code: "fum", Name: "燃油主力", SecId: "113.fum", Exchange: "上期所"},
	"TAM":    {Type: model.MarKetType_FUTURE, Code: "TAM", Name: "PTA主力", SecId: "115.TAM", Exchange: "郑商所"},
	"ppm":    {Type: model.MarKetType_FUTURE, Code: "ppm", Name: "聚丙烯主力", SecId: "114.ppm", Exchange: "大商所"},
	"PFM":    {Type: model.MarKetType_FUTURE, Code: "PFM", Name: "短纤主力", SecId: "115.PFM", Exchange: "郑商所"},
	"lm":     {Type: model.MarKetType_FUTURE, Code: "lm", Name: "塑料主力", SecId: "114.lm", Exchange: "大商所"},
	"agm":    {Type: model.MarKetType_FUTURE, Code: "agm", Name: "沪银主力", SecId: "113.agm", Exchange: "上期所"},
	"spm":    {Type: model.MarKetType_FUTURE, Code: "spm", Name: "纸浆主力", SecId: "113.spm", Exchange: "上期所"},
	"SFM":    {Type: model.MarKetType_FUTURE, Code: "SFM", Name: "硅铁主力", SecId: "115.SFM", Exchange: "郑商所"},
	"060130": {Type: model.MarKetType_FUTURE, Code: "060130", Name: "IC主力合约", SecId: "8.060130", Exchange: "中金所"},
	"pbm":    {Type: model.MarKetType_FUTURE, Code: "pbm", Name: "沪铅主力", SecId: "113.pbm", Exchange: "上期所"},
	"jdm":    {Type: model.MarKetType_FUTURE, Code: "jdm", Name: "鸡蛋主力", SecId: "114.jdm", Exchange: "大商所"},
	"egm":    {Type: model.MarKetType_FUTURE, Code: "egm", Name: "乙二醇主力", SecId: "114.egm", Exchange: "大商所"},
	"aum":    {Type: model.MarKetType_FUTURE, Code: "aum", Name: "沪金主力", SecId: "113.aum", Exchange: "上期所"},
	"150130": {Type: model.MarKetType_FUTURE, Code: "150130", Name: "IM主力合约", SecId: "8.150130", Exchange: "中金所"},
	"jmm":    {Type: model.MarKetType_FUTURE, Code: "jmm", Name: "焦煤主力", SecId: "114.jmm", Exchange: "大商所"},
	"PKM":    {Type: model.MarKetType_FUTURE, Code: "PKM", Name: "花生主力", SecId: "115.PKM", Exchange: "郑商所"},
	"SMM":    {Type: model.MarKetType_FUTURE, Code: "SMM", Name: "锰硅主力", SecId: "115.SMM", Exchange: "郑商所"},
	"jm":     {Type: model.MarKetType_FUTURE, Code: "jm", Name: "焦炭主力", SecId: "114.jm", Exchange: "大商所"},
	"ssm":    {Type: model.MarKetType_FUTURE, Code: "ssm", Name: "不锈钢主力", SecId: "113.ssm", Exchange: "上期所"},
	"wrm":    {Type: model.MarKetType_FUTURE, Code: "wrm", Name: "线材主力", SecId: "113.wrm", Exchange: "上期所"},
	"mm":     {Type: model.MarKetType_FUTURE, Code: "mm", Name: "豆粕主力", SecId: "114.mm", Exchange: "大商所"},
	"CJM":    {Type: model.MarKetType_FUTURE, Code: "CJM", Name: "红枣主力", SecId: "115.CJM", Exchange: "郑商所"},
	"SAM":    {Type: model.MarKetType_FUTURE, Code: "SAM", Name: "纯碱主力", SecId: "115.SAM", Exchange: "郑商所"},
	"csm":    {Type: model.MarKetType_FUTURE, Code: "csm", Name: "淀粉主力", SecId: "114.csm", Exchange: "大商所"},
	"MAM":    {Type: model.MarKetType_FUTURE, Code: "MAM", Name: "甲醇主力", SecId: "115.MAM", Exchange: "郑商所"},
	"cm":     {Type: model.MarKetType_FUTURE, Code: "cm", Name: "玉米主力", SecId: "114.cm", Exchange: "大商所"},
	"ebm":    {Type: model.MarKetType_FUTURE, Code: "ebm", Name: "苯乙烯主力", SecId: "114.ebm", Exchange: "大商所"},
	"rrm":    {Type: model.MarKetType_FUTURE, Code: "rrm", Name: "粳米主力", SecId: "114.rrm", Exchange: "大商所"},
	"am":     {Type: model.MarKetType_FUTURE, Code: "am", Name: "豆一主力", SecId: "114.am", Exchange: "大商所"},
	"RMM":    {Type: model.MarKetType_FUTURE, Code: "RMM", Name: "菜粕主力", SecId: "115.RMM", Exchange: "郑商所"},
	"CYM":    {Type: model.MarKetType_FUTURE, Code: "CYM", Name: "棉纱主力", SecId: "115.CYM", Exchange: "郑商所"},
	"bbm":    {Type: model.MarKetType_FUTURE, Code: "bbm", Name: "胶合板主力", SecId: "114.bbm", Exchange: "大商所"},
	"ZCM":    {Type: model.MarKetType_FUTURE, Code: "ZCM", Name: "动力煤主力", SecId: "115.ZCM", Exchange: "郑商所"},
	"WHM":    {Type: model.MarKetType_FUTURE, Code: "WHM", Name: "强麦主力", SecId: "115.WHM", Exchange: "郑商所"},
	"RIM":    {Type: model.MarKetType_FUTURE, Code: "RIM", Name: "早籼稻主力", SecId: "115.RIM", Exchange: "郑商所"},
	"PMM":    {Type: model.MarKetType_FUTURE, Code: "PMM", Name: "普麦主力", SecId: "115.PMM", Exchange: "郑商所"},
	"LRM":    {Type: model.MarKetType_FUTURE, Code: "LRM", Name: "晚籼稻主力", SecId: "115.LRM", Exchange: "郑商所"},
	"JRM":    {Type: model.MarKetType_FUTURE, Code: "JRM", Name: "粳稻主力", SecId: "115.JRM", Exchange: "郑商所"},
	"bm":     {Type: model.MarKetType_FUTURE, Code: "bm", Name: "豆二主力", SecId: "114.bm", Exchange: "大商所"},
	"130130": {Type: model.MarKetType_FUTURE, Code: "130130", Name: "TS主力合约", SecId: "8.130130", Exchange: "中金所"},
	"SRM":    {Type: model.MarKetType_FUTURE, Code: "SRM", Name: "白糖主力", SecId: "115.SRM", Exchange: "郑商所"},
	"050130": {Type: model.MarKetType_FUTURE, Code: "050130", Name: "TF主力合约", SecId: "8.050130", Exchange: "中金所"},
	"cum":    {Type: model.MarKetType_FUTURE, Code: "cum", Name: "沪铜主力", SecId: "113.cum", Exchange: "上期所"},
	"lhm":    {Type: model.MarKetType_FUTURE, Code: "lhm", Name: "生猪主力", SecId: "114.lhm", Exchange: "大商所"},
	"bum":    {Type: model.MarKetType_FUTURE, Code: "bum", Name: "沥青主力", SecId: "113.bum", Exchange: "上期所"},
	"110130": {Type: model.MarKetType_FUTURE, Code: "110130", Name: "T主力合约", SecId: "8.110130", Exchange: "中金所"},
	"rbm":    {Type: model.MarKetType_FUTURE, Code: "rbm", Name: "螺纹钢主力", SecId: "113.rbm", Exchange: "上期所"},
	"CFM":    {Type: model.MarKetType_FUTURE, Code: "CFM", Name: "棉花主力", SecId: "115.CFM", Exchange: "郑商所"},
	"RSM":    {Type: model.MarKetType_FUTURE, Code: "RSM", Name: "菜籽主力", SecId: "115.RSM", Exchange: "郑商所"},
	"hcm":    {Type: model.MarKetType_FUTURE, Code: "hcm", Name: "热卷主力", SecId: "113.hcm", Exchange: "上期所"},
	"fbm":    {Type: model.MarKetType_FUTURE, Code: "fbm", Name: "纤维板主力", SecId: "114.fbm", Exchange: "大商所"},
	"OIM":    {Type: model.MarKetType_FUTURE, Code: "OIM", Name: "菜油主力", SecId: "115.OIM", Exchange: "郑商所"},
	"ym":     {Type: model.MarKetType_FUTURE, Code: "ym", Name: "豆油主力", SecId: "114.ym", Exchange: "大商所"},
	"FGM":    {Type: model.MarKetType_FUTURE, Code: "FGM", Name: "玻璃主力", SecId: "115.FGM", Exchange: "郑商所"},
	"rum":    {Type: model.MarKetType_FUTURE, Code: "rum", Name: "橡胶主力", SecId: "113.rum", Exchange: "上期所"},
	"pm":     {Type: model.MarKetType_FUTURE, Code: "pm", Name: "棕榈油主力", SecId: "114.pm", Exchange: "大商所"},
	"im":     {Type: model.MarKetType_FUTURE, Code: "im", Name: "铁矿石主力", SecId: "114.im", Exchange: "大商所"},
	"pgm":    {Type: model.MarKetType_FUTURE, Code: "pgm", Name: "LPG主力", SecId: "114.pgm", Exchange: "大商所"},
	"nim":    {Type: model.MarKetType_FUTURE, Code: "nim", Name: "沪镍主力", SecId: "113.nim", Exchange: "上期所"},
}

func GetMarketById(id string) *model.Market {
	v, _ := allMarkets[id]
	return v
}

func GetAllFutureMarketIds() []string {
	res := []string{}
	for _, v := range allMarkets {
		if v.Type == model.MarKetType_FUTURE {
			res = append(res, v.Code)
		}
	}
	return res
}
