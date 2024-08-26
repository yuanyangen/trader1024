package utils

import "github.com/yuanyangen/trader1024/engine/logs"

func AsyncRun(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.Info("panic %v", err)

			}
		}()
		f()
	}()
}
