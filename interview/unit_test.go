package interview

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnit(t *testing.T) {
	unit_test()
}

func unit_test() {
	var a uint = 0
	var b uint = 1
	c := a - b
	fmt.Println(reflect.TypeOf(c)) //unit
	fmt.Println(c)                 //2^64 - 1，如果是32位的系统难受2^32-1
}
