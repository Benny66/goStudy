package interview

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var i int64 = 1   //当前打印次数
var j int64 = 100 //打印次数
var mySync sync.WaitGroup

func TestChannel(t *testing.T) {
	mySync.Add(1)
	var dogChan = make(chan string, 0)
	var catChan = make(chan string, 0)
	var fishChan = make(chan string, 0)
	go dogChanF(dogChan, catChan)
	go catChanF(catChan, fishChan)
	go fishChanF(fishChan, dogChan)
	dogChan <- "dog"
	mySync.Wait()
}

func dogChanF(dogChan, catChan chan string) {
	for {
		if dog, ok := <-dogChan; ok {
			fmt.Println(dog)
			catChan <- "cat"
		} else {
			return
		}
	}
}

func catChanF(catChan, fishChan chan string) {
	for {
		if cat, ok := <-catChan; ok {
			fmt.Println(cat)
			fishChan <- "fish"
		} else {
			return
		}
	}
}

func fishChanF(fishChan, dogChan chan string) {
	for {
		if fish, ok := <-fishChan; ok {
			fmt.Println(fish)
			atomic.AddInt64(&i, 1)
			if i > j {
				mySync.Done()
			}
			dogChan <- "dog"
		} else {
			return
		}
	}
}
