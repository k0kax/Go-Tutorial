package main

//互斥锁
import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex     // 互斥锁保护共享数据
	counters map[string]int // 需保护的共享字典
}

func (c *Container) inc(name string) {

	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 确保锁释放（防御式编程）
	c.counters[name]++  // 临界区操作
}

func main() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup // 同步协程执行

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3) // 注册3个待完成任务

	go doIncrement("a", 10000) // 启动两个协程操作"a"
	go doIncrement("a", 10000)
	go doIncrement("b", 10000) // 一个协程操作"b"

	wg.Wait()               // 等待所有协程完成
	fmt.Println(c.counters) // 输出结果
}
