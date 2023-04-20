package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	concurrency     = 4                        // 并发下载的数量
	bufferSize      = 512 * 1024               // 下载缓冲区的大小
	chunkSize       = 1 * 1024 * 1024          // 每个goroutine下载的块大小
	totalBufferSize = concurrency * bufferSize // 总共使用的缓冲区大小
)

// 编写一个程序，使用goroutine实现并发下载多个文件，并将每个文件保存到本地磁盘上。512内存限制，10g大小文件
func main() {
	url := "https://www.bilibili.com/video/BV1es4y127eS/?spm_id_from=333.1007.tianma.1-1-1.click"

	res, err := http.Head(url)
	if err != nil {
		fmt.Printf("Error getting file size: %v\n", err)
		return
	}

	fileSize := res.ContentLength
	if fileSize <= 0 {
		fmt.Println("Invalid file size")
		return
	}

	fmt.Printf("File size: %d bytes\n", fileSize)

	// 创建一个文件
	file, err := os.Create("large_file.txt")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer file.Close()

	// 设置HTTP请求头，以便从指定的字节偏移量开始下载文件
	headers := make([]string, concurrency)
	for i := 0; i < concurrency; i++ {
		startByte := i * chunkSize
		endByte := startByte + bufferSize - 1
		if i == concurrency-1 {
			endByte = int(fileSize) - 1
		}
		headers[i] = fmt.Sprintf("bytes=%d-%d", startByte, endByte)
	}

	// 创建一个缓冲区，用于存储已下载的数据
	buffer := make([]byte, totalBufferSize)

	// 启动并发下载任务
	var downloaded int64
	for i := 0; i < concurrency; i++ {
		go func(index int) {
			client := &http.Client{}
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			req.Header.Add("Range", headers[index])
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error downloading chunk #%d: %v\n", index, err)
				return
			}
			defer resp.Body.Close()

			// 从响应体中读取数据，并将其写入到缓冲区和文件中
			for {
				n, err := resp.Body.Read(buffer[index*bufferSize : (index+1)*bufferSize])
				if n > 0 {
					file.WriteAt(buffer[index*bufferSize:index*bufferSize+n], int64(index*chunkSize)+downloaded)
					downloaded += int64(n)
				}
				if err != nil {
					if err != io.EOF {
						fmt.Printf("Error reading data from server for chunk #%d: %v\n", index, err)
					}
					break
				}
			}
		}(i)
	}

	// 等待所有goroutine执行完毕
	for downloaded < fileSize {
		fmt.Printf("\rDownloaded %d bytes...", downloaded)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("\nFile downloaded and saved: %s\n", file.Name())
}
