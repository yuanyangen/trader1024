package config

import (
	"fmt"
)

const MasterStorageAddr = "127.0.0.1"

func GetMasterMongoUri() string {
	return fmt.Sprintf("mongodb://%v:%v/?retryWrites=true&w=majority", MasterStorageAddr, 27017)
}

func GetSlaveMongoUri() string {
	return fmt.Sprintf("mongodb://%v:%v/?retryWrites=true&w=majority", MasterStorageAddr, 27017)
}

func GetArkBaseURL() string {
	return "https://ark.cn-beijing.volces.com/api/v3"
}
func GetArkAPIKey() string {
	return "78087bb4-d2a8-4d44-b50c-a074d3227564"
}

func GetArkModelName() string {
	return "doubao-1-5-thinking-pro-250415"
}
