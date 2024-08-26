package config

import (
	"fmt"
	"os"
	"path"
)

const MasterStorageAddr = "192.168.1.106"
const DefaultHttpServerPort = 8888
const pathAddr = "HomeData/go/trader1024/data_crawler/datas"

var StorageDataPath string

func init() {
	homePath := os.Getenv("HOME")
	if homePath == "" {
		panic("HOME not in env")
	}
	StorageDataPath = path.Join(homePath, pathAddr)
}

func GetMasterHttpServerAddr() string {
	return fmt.Sprintf("http://%v:%v", MasterStorageAddr, DefaultHttpServerPort)
}

func GetLocalHttpServerAddr() string {
	return fmt.Sprintf("http://%v:%v", "127.0.0.1", DefaultHttpServerPort)
}

func GetMasterMongoUri() string {
	return fmt.Sprintf("mongodb://%v:%v/?retryWrites=true&w=majority", MasterStorageAddr, 27017)
}

func GetSlaveMongoUri() string {
	return fmt.Sprintf("mongodb://%v:%v/?retryWrites=true&w=majority", MasterStorageAddr, 27017)
}
