package storage_test

import (
	"fmt"
	"github.com/yuanyangen/trader1024/data/storage"
	"github.com/yuanyangen/trader1024/engine/model"
	"testing"
)

func TestHttpStorage(t *testing.T) {
	s := storage.EastMoneyHttpStorage()
	res := s.GetAllData("FG403", model.LineType_Day)
	fmt.Println(res)
}
