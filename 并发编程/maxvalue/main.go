package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 生成随机整数数组
	n := 1000000
	arr := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000000000)
	}

	// 设定并发数为10
	concurrency := 10
	chunkSize := (n + concurrency - 1) / concurrency

	// 启动并发协程
	var wg sync.WaitGroup
	maxChan := make(chan int, concurrency)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}
		go func(start, end int) {
			defer wg.Done()
			max := arr[start]
			for j := start + 1; j < end; j++ {
				if arr[j] > max {
					max = arr[j]
				}
			}
			maxChan <- max
		}(start, end)
	}

	// 等待所有协程结束并收集子任务的结果
	go func() {
		wg.Wait()
		close(maxChan)
	}()

	// 找到所有子任务中的最大值
	max := arr[0]
	for val := range maxChan {
		if val > max {
			max = val
		}
	}

	fmt.Println("The maximum value in the array is:", max)
}
