package config

import (
	"fmt"
)

const MasterStorageAddr = "192.168.1.106"

func GetMasterMongoUri() string {
	return fmt.Sprintf("mongodb://%v:%v/?retryWrites=true&w=majority", MasterStorageAddr, 27017)
}

func GetSlaveMongoUri() string {
	return fmt.Sprintf("mongodb://%v:%v/?retryWrites=true&w=majority", MasterStorageAddr, 27017)
}
