package main

//WaitGroup 等待组？
import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) //增加计数

		go func() {
			defer wg.Done() //减少计数，等价于 Add(-1)
			worker(i)
		}()
	}
	wg.Wait() //阻塞等待
	//wg.Wait() 通过信号量阻塞主协程，直到所有 Done() 调用使计数器归零,才解除阻塞

}

/*
执行流程时序图

主协程                              Worker协程（G1-G5）
├─1. 初始化 wg (sync.WaitGroup)───────→|
├─2. 循环启动5个协程                   |→ 创建G1-G5，状态：_Grunnable[3](@ref)
│   ├─wg.Add(1)（计数器从0→1→2→3→4→5） |
│   └─go闭包捕获变量i（存在竞态风险）     |→ G1-G5进入P本地队列[9](@ref)
├─3. 执行wg.Wait()───────────────────→|→ 主协程阻塞（状态：_Gwaiting）
│                                     |→ 调度器将CPU时间片分配给G1-G5[4](@ref)
│                                     |
├─4. G1-G5并发执行阶段─────────────────→|
│   │                                 |→ G1: 打印"Worker 6 starting"（闭包陷阱）
│   │                                 |→ G2: 打印"Worker 6 starting"（i值已递增）
│   │                                 |→ G3-G5同理（均可能输出i的最终值6）
│   ├─time.Sleep(1秒)────────────────→|→ G1-G5进入_Gwaiting状态[5](@ref)
│   └─defer wg.Done()───────────────→ |→ 计数器递减（5→4→3→2→1→0）
│                                     |→ 每次Done触发runtime_Semrelease[8](@ref)
├─5. 所有G完成，主协程唤醒───────────────→|← 当计数器归零，runtime_Semacquire解除阻塞
└─6. 主协程继续执行─────────────────────→| 输出程序结束

*/
