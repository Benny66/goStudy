package main

import (
	"fmt"
	"testing"
)

/*
defer 在程序返回时执行，后入先出，
*/
func TestDefer2(t *testing.T) {
	// fmt.Println("return:", b())
	fmt.Println(f7())
}

func f1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func f2() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func f3() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func f4() int {
	t := 5
	defer func() {
		t += 1
	}()
	return t
}

func f5() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f6() (r int) {
	t := 5
	defer func() {
		r = t + 5
	}()
	return r
}

func f7() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return r
}
