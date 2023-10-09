package main

import "github.com/yuanyangen/trader1024/config"

func Init() {
	InitAllStorage(config.StorageDataPath, []string{"eastmoney", "main", "sina", "test"})
}

func InitAllStorage(dirPath string, names []string) {
	var dbSplitCount = 128
	for _, name := range names {
		InitStorage(dirPath, name, dbSplitCount)
	}
}
