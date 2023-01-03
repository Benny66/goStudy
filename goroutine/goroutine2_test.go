package goroutine

/*
 * @Author: wenzhicong wenzhicong@jasonanime.com
 * @Date: 2022-12-28 14:24:16
 * @LastEditors: wenzhicong wenzhicong@jasonanime.com
 * @LastEditTime: 2023-01-03 16:24:53
 * @FilePath: /goStudy/goroutine/goroutine2_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGoroutine2(t *testing.T) {
	runtime.GOMAXPROCS(1)
	var sync sync.WaitGroup
	sync.Add(6)
	go func() {
		defer sync.Done()
		fmt.Println("goroutine 1")
		go func() {
			defer sync.Done()
			fmt.Println("goroutine 1-1")
			go func() {
				defer sync.Done()
				fmt.Println("goroutine 1-1-1")
			}()
		}()
		go func() {
			defer sync.Done()
			fmt.Println("goroutine 1-2")
			go func() {
				defer sync.Done()
				fmt.Println("goroutine 1-2-1")
			}()
		}()
	}()
	go func() {
		defer sync.Done()
		fmt.Println("goroutine 2")
	}()
	sync.Wait()
	return
}
