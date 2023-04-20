package main

import "fmt"

// 编写一个程序，使用goroutine并发计算斐波那契数列，并输出结果。
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(100, c)
	for i := range c {
		fmt.Println(i)
	}
}
