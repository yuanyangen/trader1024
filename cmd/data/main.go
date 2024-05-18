package main

import (
	"github.com/yuanyangen/trader1024/config"
	stroage_server "github.com/yuanyangen/trader1024/data/storage_server"
)

func main() {
	stroage_server.InitAllStorage(config.StorageDataPath, []string{"eastmoney", "main", "sina", "test"})
	asyncStartDataCrawler()
	startDataServer()
}
