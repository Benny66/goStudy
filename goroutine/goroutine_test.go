package goroutine

/*
 * @Author: wenzhicong wenzhicong@jasonanime.com
 * @Date: 2022-12-27 10:45:44
 * @LastEditors: wenzhicong wenzhicong@jasonanime.com
 * @LastEditTime: 2023-01-03 16:25:24
 * @FilePath: /goStudy/goroutine/goroutine_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

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
