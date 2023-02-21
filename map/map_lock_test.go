package map_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMapLock(t *testing.T) {
	var mapLock = mapLock{
		Data: make(map[int]int, 0),
	}
	go func() {
		for i := 1; i < 11; i++ {
			mapLock.RLock()
			mapLock.Data[i] = i * 100
			mapLock.RUnlock()
		}
	}()
	go func() {
		for i := 1; i < 11; i++ {
			mapLock.RLock()
			mapLock.Data[i] = i * 1001
			mapLock.RUnlock()
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println(mapLock.Data)
}

type mapLock struct {
	sync.RWMutex
	Data map[int]int
}
