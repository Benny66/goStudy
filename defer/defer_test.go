package main

import (
	"fmt"
	"testing"
)

/*
defer 在程序返回时执行，后入先出，
*/
func TestDefer(t *testing.T) {
	// fmt.Println("return:", b())
	fmt.Println("return:", *(c()))
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}
