package main

import (
	"fmt"
)

func main() {
	reqCh := make(chan request)
	respCh := make(chan response)

	// 启动 10 个 goroutine 并发写入 map
	for i := 0; i < 10; i++ {
		go writeMap(reqCh, respCh)
	}

	// 启动 10 个 goroutine 并发读取 map
	for i := 0; i < 10; i++ {
		go readMap(reqCh, respCh)
	}

	// 向请求通道中发送读取和写入请求
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			reqCh <- request{operation: "write", key: fmt.Sprintf("key%d", i/2), value: i / 2}
		} else {
			reqCh <- request{operation: "read", key: fmt.Sprintf("key%d", i/2)}
		}
	}

	// 打印接收到的响应
	for i := 0; i < 20; i++ {
		resp := <-respCh
		if resp.err != nil {
			fmt.Printf("error: %v\n", resp.err)
		} else {
			fmt.Printf("key: %s, value: %d\n", resp.key, resp.value)
		}
	}
}

type request struct {
	operation string
	key       string
	value     int
}

type response struct {
	key   string
	value int
	err   error
}

func writeMap(reqCh chan request, respCh chan response) {
	m := make(map[string]int)
	for req := range reqCh {
		if req.operation == "write" {
			m[req.key] = req.value
			respCh <- response{key: req.key, value: req.value}
		} else {
			respCh <- response{key: req.key, err: fmt.Errorf("invalid operation: %s", req.operation)}
		}
	}
}

func readMap(reqCh chan request, respCh chan response) {
	m := make(map[string]int)
	for req := range reqCh {
		if req.operation == "read" {
			value, ok := m[req.key]
			if ok {
				respCh <- response{key: req.key, value: value}
			} else {
				respCh <- response{key: req.key, err: fmt.Errorf("key not found: %s", req.key)}
			}
		} else {
			respCh <- response{key: req.key, err: fmt.Errorf("invalid operation: %s", req.operation)}
		}
	}
}
