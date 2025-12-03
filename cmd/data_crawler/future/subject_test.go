package main

import (
	"context"
	"github.com/yuanyangen/trader1024/dal/mongo"
	"github.com/yuanyangen/trader1024/engine/model"
	"testing"
)

func TestSaveToMongo(t *testing.T) {
	//time.Local = time.FixedZone("CST", 8*3600)
	//
	//tn := time.Now()
	//expireT1 := tn.Format("2006-01-02 15:04:05")
	//expireT, _ := time.Parse("2006-01-02 15:04:05", expireT1)
	//logs.Info(tn.Unix())
	//logs.Info(expireT1)
	//logs.Info(expireT.Unix())
	return
	//Init()
	for _, v := range AllSubjects {
		mongo.SaveSubjectInfo(context.Background(), v)
	}
}

var AllSubjects = map[string]*model.SubjectDO{
	"不锈钢": {CNName: "不锈钢", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年9月25日", FirstContract: "20200323", OnlineTime: "1", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"橡胶":  {CNName: "橡胶", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年3月27日", FirstContract: "199405", ContractMonth: []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11}, ContractLastTradeDay: 15, RealContractLastTradeDay: 10},
	"沥青":  {CNName: "沥青", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年10月9日", FirstContract: "20140407"},
	"沪金":  {CNName: "沪金", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2008年1月9日", FirstContract: "20080707", ContractMonth: []int{3, 6, 9, 12}},
	"沪铅":  {CNName: "沪铅", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2011年3月24日", FirstContract: "20110920", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"沪铜":  {CNName: "沪铜", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年3月1日", FirstContract: "19930828", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"沪铝":  {CNName: "沪铝", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1992年5月28日", FirstContract: "19921124", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"沪银":  {CNName: "沪银", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年5月10日", FirstContract: "20121106", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"沪锌":  {CNName: "沪锌", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2007年3月26日", FirstContract: "20070922", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"沪锡":  {CNName: "沪锡", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2015年3月27日", FirstContract: "20150923", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"沪镍":  {CNName: "沪镍", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2015年3月27日", FirstContract: "20150923", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"热卷":  {CNName: "热卷", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年3月21日", FirstContract: "20140917", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"燃油":  {CNName: "燃油", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2018年3月26日", FirstContract: "20180922", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"纸浆":  {CNName: "纸浆", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2018年11月27日", FirstContract: "20190526", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"线材":  {CNName: "线材", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年3月1日", FirstContract: "19930828", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"螺纹钢": {CNName: "螺纹钢", Exchange: "上期所", Type: model.MarKetType_FUTURE, OnlineDay: "2009年3月27日", FirstContract: "20090923", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},

	"IC合约": {CNName: "IC合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "2015年4月16日", FirstContract: "20151013"},
	"IF合约": {CNName: "IF合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "2010年4月16日", FirstContract: "20101013"},
	"IH合约": {CNName: "IH合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"IM合约": {CNName: "IM合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"TF合约": {CNName: "TF合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"TS合约": {CNName: "TS合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},
	"T合约":  {CNName: "T合约", Exchange: "中金所", Type: model.MarKetType_FUTURE, OnlineDay: "", FirstContract: "00010630"},

	"LPG": {CNName: "LPG", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2020年3月30日", FirstContract: "20200926", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"PVC": {CNName: "PVC", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2009年5月25日", FirstContract: "20091121", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"乙二醇": {CNName: "乙二醇", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2020年3月30日", FirstContract: "20200926", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"塑料":  {CNName: "塑料", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2007年6月1日", FirstContract: "20071128", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"棕榈油": {CNName: "棕榈油", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2007年10月29日", FirstContract: "20080426", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"淀粉":  {CNName: "淀粉", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年12月19日", FirstContract: "20150617", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"焦炭":  {CNName: "焦炭", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2011年4月15日", FirstContract: "20111012"},
	"焦煤":  {CNName: "焦煤", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年3月22日", FirstContract: "20130918", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"玉米":  {CNName: "玉米", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2004年9月22日", FirstContract: "200711", ContractMonth: []int{1, 3, 5, 7, 9, 11}, ContractLastTradeDay: 10, RealContractLastTradeDay: 1},
	"生猪":  {CNName: "生猪", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2021年1月8日", FirstContract: "20210707"},
	"粳米":  {CNName: "粳米", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年8月16日", FirstContract: "20200212", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"纤维板": {CNName: "纤维板", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年12月6日", FirstContract: "20140604", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"聚丙烯": {CNName: "聚丙烯", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年2月28日", FirstContract: "20140827", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"胶合板": {CNName: "胶合板", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年12月6日", FirstContract: "20140604", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"苯乙烯": {CNName: "苯乙烯", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年9月26日", FirstContract: "20200324", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"豆一":  {CNName: "豆一", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2002年3月15日", FirstContract: "20020911", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"豆二":  {CNName: "豆二", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2004年12月22日", FirstContract: "20050620", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"豆油":  {CNName: "豆油", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2006年1月9日", FirstContract: "20060708", ContractMonth: []int{1, 3, 5, 7, 8, 9, 11, 12}},
	"豆粕":  {CNName: "豆粕", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2000年7月17日", FirstContract: "20010113", ContractMonth: []int{1, 3, 5, 7, 8, 9, 11, 12}},
	"铁矿石": {CNName: "铁矿石", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年10月18日", FirstContract: "20140416", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"鸡蛋":  {CNName: "鸡蛋", Exchange: "大商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年11月8日", FirstContract: "20140507", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},

	"PTA": {CNName: "PTA", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2006年12月18日", FirstContract: "20070616", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"动力煤": {CNName: "动力煤", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年9月26日", FirstContract: "20140325", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"尿素":  {CNName: "尿素", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年8月9日", FirstContract: "20200205", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"强麦":  {CNName: "强麦", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年1月17日", FirstContract: "20120715", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"晚籼稻": {CNName: "晚籼稻", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年7月8日", FirstContract: "20150104", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"普麦":  {CNName: "普麦", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "1993年7月27日", FirstContract: "19940123", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"棉纱":  {CNName: "棉纱", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2017年8月18日", FirstContract: "20180214", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"棉花":  {CNName: "棉花", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2004年10月28日", FirstContract: "20050426", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"玻璃":  {CNName: "玻璃", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年12月3日", FirstContract: "20130601"},
	"甲醇":  {CNName: "甲醇", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2011年10月28日", FirstContract: "20120425", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"白糖":  {CNName: "白糖", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2006年1月6日", FirstContract: "20060705", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"短纤":  {CNName: "短纤", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2020年10月12日", FirstContract: "20210410", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"硅铁":  {CNName: "硅铁", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年8月8日", FirstContract: "20150204", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"粳稻":  {CNName: "粳稻", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2013年11月18日", FirstContract: "20140517", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"红枣":  {CNName: "红枣", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年4月30日", FirstContract: "20191027", ContractMonth: []int{1, 3, 5, 7, 9, 12}},
	"纯碱":  {CNName: "纯碱", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2019年12月6日", FirstContract: "20200603", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	"花生":  {CNName: "花生", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2021年2月1日", FirstContract: "20210731", ContractMonth: []int{1, 3, 4, 10, 11, 12}},
	"苹果":  {CNName: "苹果", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2017年12月22日", FirstContract: "20180620", ContractMonth: []int{1, 3, 4, 5, 10, 11, 12}},
	"菜油":  {CNName: "菜油", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2022年8月26日", FirstContract: "20230222", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"菜籽":  {CNName: "菜籽", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年12月28日", FirstContract: "20130626", ContractMonth: []int{7, 8, 9, 11}},
	"菜粕":  {CNName: "菜粕", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2012年12月28日", FirstContract: "20130626", ContractMonth: []int{1, 3, 5, 7, 9, 11}},
	"锰硅":  {CNName: "锰硅", Exchange: "郑商所", Type: model.MarKetType_FUTURE, OnlineDay: "2014年8月8日", FirstContract: "20150204", ContractMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
}
