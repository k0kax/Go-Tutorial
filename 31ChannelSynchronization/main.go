package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
	//主协程阻塞等待接收信号，通道接收的核心同步手段
	//其作用是通过阻塞当前协程（如主协程）,直到从其他协程（如工作协程）接收到信号
	//阻塞条件​​：若通道无数据，主协程进入阻塞状态（G 状态转为 Waiting）
	//唤醒条件​​：当 done 通道接收到数据时，调度器将主协程重新标记为可运行（Runnable）

	/*
			主协程启动
		│
		├─ 创建缓冲通道 done
		├─ 启动 worker 协程（异步）
		│   │
		│   ├─ 打印 "working..."（立即输出）
		│   ├─ 休眠 1 秒（协程挂起）
		│   ├─ 恢复后打印 "done"
		│   └─ 发送 true 到 done（缓冲区写入）
		│
		├─ 主协程执行 <-done（缓冲区有数据，立即读取）
		│
		└─ 程序退出
	*/
}
