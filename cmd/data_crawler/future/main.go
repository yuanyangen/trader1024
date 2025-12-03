package main

import "sync"

func main() {
	asyncStartDataCrawler()
	wait()
}

func wait() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
