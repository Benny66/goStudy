package main

/*
 * @Author: wenzhicong wenzhicong@jasonanime.com
 * @Date: 2022-12-26 11:16:22
 * @LastEditors: wenzhicong wenzhicong@jasonanime.com
 * @LastEditTime: 2022-12-27 10:08:58
 * @FilePath: /channel/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	num := 1000000
	startTime := time.Now().UnixNano() / int64(time.Millisecond)
	channelAdd(num, 4)
	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(endTime - startTime)

	startTime = time.Now().UnixNano() / int64(time.Millisecond)
	forAdd(num)
	endTime = time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(endTime - startTime)
}

func channelAdd(n int, c int) int64 {
	var sync sync.WaitGroup
	var chans = make(chan int, 4)
	sync.Add(c)
	var addResult int64
	for i := 0; i < c; i++ {
		go func(work int) {
			defer sync.Done()
			for {
				data, ok := <-chans
				if !ok {
					break
				}
				if isPrime(data) {
					// 数据争用,使用atomic
					atomic.AddInt64(&addResult, 1)
				}
			}
		}(i)
	}
	for i := 1; i < n; i++ {
		chans <- i
	}
	close(chans)
	sync.Wait()
	fmt.Println(addResult)
	return addResult
}

func forAdd(n int) int {
	var addResult int
	for i := 1; i < n; i++ {
		if isPrime(i) {
			addResult += 1
		}
	}
	fmt.Println(addResult)
	return addResult
}

// 判断n是不是质数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
