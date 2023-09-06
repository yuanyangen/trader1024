package main

func Init() {
	const storageDataPath = "/home/yuanyangen/HomeData/go/trader1024/data/datas"
	InitAllStorage(storageDataPath, []string{"eastmoney", "main", "sina", "test"})
}

func InitAllStorage(dirPath string, names []string) {
	var dbSplitCount = 128
	for _, name := range names {
		InitStorage(dirPath, name, dbSplitCount)
	}
}
