package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	m := make(map[string]string)
	m["name"] = "world"

	wg.Add(1) // 增加 WaitGroup 计数
	go func() {
		defer wg.Done() // 在函数结束时减少 WaitGroup 计数
		mu.Lock()
		defer mu.Unlock()
		m["name"] = "data race"
	}()

	wg.Wait()         // 等待 goroutine 完成
	mu.Lock()         // 添加锁以保护对 m 的访问
	defer mu.Unlock() // 确保在函数结束时解锁
	fmt.Println("Hello,", m["name"])
}
