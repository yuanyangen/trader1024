package utils

import "os"

func CreateDirIfNotExist(path string) {
    _,err:=os.Stat(path)
    if err != nil {
        err = os.MkdirAll(path, 0755)
        if err != nil {
            panic("create dir error")
        }
    }
}