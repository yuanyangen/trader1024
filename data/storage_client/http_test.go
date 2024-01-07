package storage_client_test

import (
	"fmt"
	"github.com/yuanyangen/trader1024/data/storage_client"
	"github.com/yuanyangen/trader1024/engine/model"
	"testing"
)

func TestHttpStorage(t *testing.T) {
	s := storage_client.SinaHttpStorage()
	res := s.GetAllData("玉米201001", model.LineType_Day)
	fmt.Println(res)
}

func TestKvStorage(t *testing.T) {
	s := &storage_client.HttpStorageClient{
		Name: "test",
	}
	s.SaveData("测试test1", model.LineType_Day, []*model.KNode{&model.KNode{
		Open:      231111,
		Close:     231111,
		TimeStamp: 13213111,
	}, {
		Open:      22532311,
		Close:     23353454341,
		TimeStamp: 13435324213211,
	}})
	res := s.GetAllData("测试test1", model.LineType_Day)
	fmt.Println(res)
}
