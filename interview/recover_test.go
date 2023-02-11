package interview

import (
	"errors"
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	fmt.Println(test())
}
func test() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%s", r))
		}
	}()
	err = raisePanic()
	return err
}
func raisePanic() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("raisePanic函数异常：%s", r))
		}
	}()
	panic("发生了错误")
}
