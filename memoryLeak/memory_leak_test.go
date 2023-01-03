package memoryleak

/*
 * @Author: wenzhicong wenzhicong@jasonanime.com
 * @Date: 2023-01-03 16:32:32
 * @LastEditors: wenzhicong wenzhicong@jasonanime.com
 * @LastEditTime: 2023-01-03 16:35:40
 * @FilePath: /goStudy/memoryLeak/memory_leak_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func TestMemoryLeak(t *testing.T) {
	//内存泄露
	Slice1()
	//Slice2 开辟一段新的内存地址，Slice1 和 Slice2不共享同一块内存地址
	// Slice2()
	http.ListenAndServe(":9999", nil)
}
func Slice1() {
	slice1 := []int{3, 4, 5, 6, 7}
	slice2 := slice1[1:3]

	fmt.Printf("slice1 addr: %p", &slice1)
	fmt.Println()
	fmt.Printf("slice2 addr: %p", &slice2)
	fmt.Println()

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%v:[%v]  ", slice1[i], &slice1[i])
	}
	fmt.Println()
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("%v:[%v]  ", slice2[i], &slice2[i])
	}
	fmt.Println()
}

func Slice2() {
	initSlice := []int{3, 4, 5, 6, 7}
	//partSlice := initSlice[1:3]

	var partSlice []int
	partSlice = append(partSlice, initSlice[1:3]...) // append

	fmt.Printf("initSlice addr: %p", &initSlice)
	fmt.Println()
	fmt.Printf("partSlice addr: %p", &partSlice)
	fmt.Println()

	for i := 0; i < len(initSlice); i++ {
		fmt.Printf("%v:[%v]  ", initSlice[i], &initSlice[i])
	}
	fmt.Println()
	for i := 0; i < len(partSlice); i++ {
		fmt.Printf("%v:[%v]  ", partSlice[i], &partSlice[i])
	}
	fmt.Println()
}
