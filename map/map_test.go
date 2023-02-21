package map_test

import (
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	var data = make(map[int]int, 0)
	go func() {
		for i := 1; i < 11; i++ {
			data[i] = i * 100
		}
	}()
	go func() {
		for i := 1; i < 11; i++ {
			data[i] = i * 1000
		}
	}()
	time.Sleep(10 * time.Second)

}
