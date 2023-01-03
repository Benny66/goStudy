package main


import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
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
