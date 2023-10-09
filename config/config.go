package config

import (
	"os"
	"path"
)

const HttpStorageAddr = "http://127.0.0.1:8888"
const pathAddr = "HomeData/go/trader1024/data/datas"

var StorageDataPath string

func init() {
	homePath := os.Getenv("HOME")
	if homePath == "" {
		panic("HOME not in env")
	}
	StorageDataPath = path.Join(homePath, pathAddr)
}
