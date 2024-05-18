package stroage_server

func InitAllStorage(dirPath string, names []string) {
	var dbSplitCount = 128
	for _, name := range names {
		InitStorage(dirPath, name, dbSplitCount)
	}
}
