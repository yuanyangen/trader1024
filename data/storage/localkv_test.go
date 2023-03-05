package storage_test

import (
    "fmt"
    "github.com/yuanyangen/trader1024/data/storage"
    "github.com/yuanyangen/trader1024/engine/model"
    "testing"
)

func TestKvStorage(t *testing.T) {
    s := storage.InitHttpStorage("test")
    s.SaveData("test", model.LineType_Day,[]*model.KNode{&model.KNode{
        Open: 231,
        Close: 231,
        TimeStamp: 132131,
        },{
        Open: 23231,
        Close: 23341,
        TimeStamp: 13213211,
        }})
    res:=  s.GetAllData("test", model.LineType_Day)
    fmt.Println(res )
}