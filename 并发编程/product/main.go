package main

import (
	"fmt"
	"time"
)

// 实现一个生产者-消费者模型，其中包含一个生产者goroutine和多个消费者goroutine，生产者会不断生成数据放入通道中，而消费者则不断从通道中取出数据进行处理。
func main() {
	var mChan = make(chan int, 4)
	go func(mChan chan int) {
		production(mChan)
		defer close(mChan)
	}(mChan)

	for i := range mChan {
		go func(i int) {
			fmt.Println(i)
		}(i)
		time.Sleep(500 * time.Millisecond) // 消费者休眠500毫秒
	}
}

func production(mChan chan int) {
	for i := 0; i < 100; i++ {
		mChan <- i
		time.Sleep(100 * time.Millisecond) // 生产者休眠100毫秒
	}
}
