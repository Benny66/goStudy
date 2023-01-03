package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
1、runtime.GOMAXPROCS(1) 设置了 P 只有一个，即程序只能运行一个 协程（Goroutine）；
2、主协程 main 被挂载在这个唯一 P 中；
3、两个 for 循环创建了 20 个协程（Goroutine），因为 P 正在被占用，所有的 Goroutine 依次挂在这个 P 的待执行的队列中
	这里需要关注一下协程的创建与挂载在 P 下的逻辑。
	因为执行结果可见，如果忽视第一行结果可见，这个 Goroutine 的执行是 FIFO （先进先出）的关系
	而第一行结果出现的原因是，P 有两个队列，一个是这个 FIFO 的队列，还有一个叫做 runnext 的队列。新创建的 Goroutine 会先被挂在这个 runnext 中；
*/
func TestGoroutine(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i1: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i2: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
