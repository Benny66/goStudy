package closure

import (
	"context"
	"testing"
	"time"
)

/*
	闭包
*/
func TestClosure(t *testing.T) {
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithTimeout(ctx, time.Second*3)
	go func() {
		x()()
	}()
	for {
		select {
		case <-cancelCtx.Done():
			cancelFunc()
			return
		case <-time.After(time.Second * 100):
			cancelFunc()
			return
		}
	}
}

//x->y fn -> func (1、打印z，2、执行y) -> func (1、打印z，2、执行y)....无限嵌套,直到上下文3s超时关闭，程序return
func x() (y func()) {
	y = func() {
		println("y")
	}
	return func() {
		println("z")
		y()
	}
}
