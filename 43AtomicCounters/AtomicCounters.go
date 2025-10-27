// 原子计数器
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1) // 注册一个待完成的协程

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1) // 原子递增计数器
			}
			wg.Done() // 标记协程完成
		}()
	}
	wg.Wait() // 阻塞等待所有协程完成

	fmt.Println("ops:", ops.Load()) // 输出最终值
}
