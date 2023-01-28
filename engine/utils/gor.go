package utils

func AsyncRun(f func()) {
	go func() {
		//defer func() {
		//	if err := recover(); err != nil {
		//		logs.
		//	}
		//}()

		f()
	}()
}
