package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// 并发程序，可以从一个URL列表中同时下载多个文件，并统计总共下载的字节数
func download(url string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: %v\n", url, err)
		return
	}

	ch <- len(data)
	fmt.Printf("%s downloaded (%d bytes)\n", url, len(data))
}

func main() {
	urls := []string{
		"https://www.example.com/",
		"https://golang.org/",
		"https://www.google.com/",
	}

	var wg sync.WaitGroup
	ch := make(chan int)

	for _, url := range urls {
		wg.Add(1)
		go download(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalBytes := 0
	for bytes := range ch {
		totalBytes += bytes
	}

	fmt.Printf("Total bytes downloaded: %d\n", totalBytes)
}
