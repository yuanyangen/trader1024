package markets

import (
	"github.com/yuanyangen/trader1024/engine/model"
)

// 合约是根据上市时间+180天计算出来的， 可能会错+漏， 但是影下你个很小， 忽略不计。
var AllSubjects = map[string]*model.Subject{
	"不锈钢": {CNName: "不锈钢", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年9月25日", FirstContract: "20200323"},
	"橡胶":  {CNName: "橡胶", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年3月27日", FirstContract: "19930923"},
	"沥青":  {CNName: "沥青", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年10月9日", FirstContract: "20140407"},
	"沪金":  {CNName: "沪金", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2008年1月9日", FirstContract: "20080707"},
	"沪铅":  {CNName: "沪铅", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2011年3月24日", FirstContract: "20110920"},
	"沪铜":  {CNName: "沪铜", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年3月1日", FirstContract: "19930828"},
	"沪铝":  {CNName: "沪铝", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1992年5月28日", FirstContract: "19921124"},
	"沪银":  {CNName: "沪银", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年5月10日", FirstContract: "20121106"},
	"沪锌":  {CNName: "沪锌", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2007年3月26日", FirstContract: "20070922"},
	"沪锡":  {CNName: "沪锡", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2015年3月27日", FirstContract: "20150923"},
	"沪镍":  {CNName: "沪镍", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2015年3月27日", FirstContract: "20150923"},
	"热卷":  {CNName: "热卷", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年3月21日", FirstContract: "20140917"},
	"燃油":  {CNName: "燃油", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2018年3月26日", FirstContract: "20180922"},
	"纸浆":  {CNName: "纸浆", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2018年11月27日", FirstContract: "20190526"},
	"线材":  {CNName: "线材", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年3月1日", FirstContract: "19930828"},
	"螺纹钢": {CNName: "螺纹钢", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2009年3月27日", FirstContract: "20090923"},

	"IC合约": {CNName: "IC合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "2015年4月16日", FirstContract: "20151013"},
	"IF合约": {CNName: "IF合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "2010年4月16日", FirstContract: "20101013"},
	"IH合约": {CNName: "IH合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"IM合约": {CNName: "IM合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"TF合约": {CNName: "TF合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"TS合约": {CNName: "TS合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"T合约":  {CNName: "T合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},

	"LPG": {CNName: "LPG", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2020年3月30日", FirstContract: "20200926"},
	"PVC": {CNName: "PVC", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2009年5月25日", FirstContract: "20091121"},
	"乙二醇": {CNName: "乙二醇", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2020年3月30日", FirstContract: "20200926"},
	"塑料":  {CNName: "塑料", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2007年6月1日", FirstContract: "20071128"},
	"棕榈油": {CNName: "棕榈油", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2007年10月29日", FirstContract: "20080426"},
	"淀粉":  {CNName: "淀粉", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年12月19日", FirstContract: "20150617"},
	"焦炭":  {CNName: "焦炭", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2011年4月15日", FirstContract: "20111012"},
	"焦煤":  {CNName: "焦煤", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年3月22日", FirstContract: "20130918"},
	"玉米":  {CNName: "玉米", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2004年9月22日", FirstContract: "20050321"},
	"生猪":  {CNName: "生猪", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2021年1月8日", FirstContract: "20210707"},
	"粳米":  {CNName: "粳米", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年8月16日", FirstContract: "20200212"},
	"纤维板": {CNName: "纤维板", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年12月6日", FirstContract: "20140604"},
	"聚丙烯": {CNName: "聚丙烯", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年2月28日", FirstContract: "20140827"},
	"胶合板": {CNName: "胶合板", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年12月6日", FirstContract: "20140604"},
	"苯乙烯": {CNName: "苯乙烯", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年9月26日", FirstContract: "20200324"},
	"豆一":  {CNName: "豆一", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2002年3月15日", FirstContract: "20020911"},
	"豆二":  {CNName: "豆二", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2004年12月22日", FirstContract: "20050620"},
	"豆油":  {CNName: "豆油", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2006年1月9日", FirstContract: "20060708"},
	"豆粕":  {CNName: "豆粕", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2000年7月17日", FirstContract: "20010113"},
	"铁矿石": {CNName: "铁矿石", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年10月18日", FirstContract: "20140416"},
	"鸡蛋":  {CNName: "鸡蛋", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年11月8日", FirstContract: "20140507"},

	"PTA": {CNName: "PTA", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2006年12月18日", FirstContract: "20070616"},
	"动力煤": {CNName: "动力煤", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年9月26日", FirstContract: "20140325"},
	"尿素":  {CNName: "尿素", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年8月9日", FirstContract: "20200205"},
	"强麦":  {CNName: "强麦", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年1月17日", FirstContract: "20120715"},
	"晚籼稻": {CNName: "晚籼稻", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年7月8日", FirstContract: "20150104"},
	"普麦":  {CNName: "普麦", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年7月27日", FirstContract: "19940123"},
	"棉纱":  {CNName: "棉纱", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2017年8月18日", FirstContract: "20180214"},
	"棉花":  {CNName: "棉花", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2004年10月28日", FirstContract: "20050426"},
	"玻璃":  {CNName: "玻璃", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年12月3日", FirstContract: "20130601"},
	"甲醇":  {CNName: "甲醇", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2011年10月28日", FirstContract: "20120425"},
	"白糖":  {CNName: "白糖", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2006年1月6日", FirstContract: "20060705"},
	"短纤":  {CNName: "短纤", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2020年10月12日", FirstContract: "20210410"},
	"硅铁":  {CNName: "硅铁", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年8月8日", FirstContract: "20150204"},
	"粳稻":  {CNName: "粳稻", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年11月18日", FirstContract: "20140517"},
	"红枣":  {CNName: "红枣", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年4月30日", FirstContract: "20191027"},
	"纯碱":  {CNName: "纯碱", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年12月6日", FirstContract: "20200603"},
	"花生":  {CNName: "花生", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2021年2月1日", FirstContract: "20210731"},
	"苹果":  {CNName: "苹果", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2017年12月22日", FirstContract: "20180620"},
	"菜油":  {CNName: "菜油", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2022年8月26日", FirstContract: "20230222"},
	"菜籽":  {CNName: "菜籽", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年12月28日", FirstContract: "20130626"},
	"菜粕":  {CNName: "菜粕", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年12月28日", FirstContract: "20130626"},
	"锰硅":  {CNName: "锰硅", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年8月8日", FirstContract: "20150204"},
}

type VendorMarket interface {
	GetMarketFromVendor(CNName string, date string) *model.Contract
}

func GetSubjectByCnNam(name string) *model.Subject {
	v, _ := AllSubjects[name]
	return v
}

func GetAllFutureMarketIds() []string {
	res := []string{}
	for _, v := range AllSubjects {
		if v.Type == model.MarKetType_FUTURE {
			res = append(res, v.CNName)
		}
	}
	return res
}

func GetAllFutureSubjects() []*model.Subject {
	res := []*model.Subject{}
	for _, v := range AllSubjects {
		if v.Type == model.MarKetType_FUTURE {
			res = append(res, v)
		}
	}
	return res
}
