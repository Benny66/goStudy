package interview

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	range_1()
	range_2()
}

func range_1() {
	a := [3]int{1, 2, 3} // 数组，值类型
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a) //打印100，200，3
		}
		a[k] = 100 + v
		//因为a是数组，for循环时进行值复制，使用不改变循环v的值
		//当k=0，a[0] = 100 + 1
		//当k=1，a[1] = 100 + 2
		//当k=2，a[2] = 100 + 3
	}
	fmt.Println(a) //打印101，102，103
}

func range_2() {
	a := []int{1, 2, 3} // 切片，引用类型
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a) //打印100，200，3
		}
		a[k] = 100 + v
		//当k=0，a[0] = 100 + 1
		//当k=1，a[1] = 100 + 200
		//当k=2，a[2] = 100 + 3
	}
	fmt.Println(a) //1，2，3
}
