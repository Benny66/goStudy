package main

/*
 * @Author: wenzhicong wenzhicong@jasonanime.com
 * @Date: 2023-01-03 11:39:45
 * @LastEditors: wenzhicong wenzhicong@jasonanime.com
 * @LastEditTime: 2023-01-03 16:22:48
 * @FilePath: /defer/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"fmt"
	"testing"
)

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
