package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//消费者协程
	go func() {
		for {
			j, more := <-jobs //检测通道是否关闭
			if more {         // 通道未关闭且有数据
				fmt.Println("received job", j)
			} else { // 通道已关闭且无数据
				fmt.Println("received all jobs")
				done <- true // 通知主协程任务完成
				return
			}
		}
	}()

	//主协程
	for j := 1; j <= 3; j++ {
		jobs <- j // 发送任务到 jobs,通知消费者协程数据已发送完毕
		fmt.Println("sent job", j)
	}
	close(jobs) // 发送完毕后关闭通道,chan没法无法接受数据
	fmt.Println("sent all jobs")

	<-done // 等待消费者协程完成 取出值
	//协程阻塞等待接收信号，通道接收的核心同步手段
	//其作用是通过阻塞当前协程（如主协程）,直到从其他协程（如工作协程）接收到信号
	//阻塞条件​​：若通道无数据，主协程进入阻塞状态（G 状态转为 Waiting）
	//唤醒条件​​：当 done 通道接收到数据时，调度器将主协程重新标记为可运行（Runnable）

	_, ok := <-jobs // 检测通道是否已关闭
	fmt.Println("received more jobs:", ok)

}

/*
执行流程时序图​

主协程                            消费者协程
|                                     |
| 创建 jobs（缓冲）和 done（无缓冲）    |
|-------------------------------------|
| 启动消费者协程                       |→ 进入循环监听 jobs
|                                     |
| 发送任务1 → jobs                    |→ 接收任务1，打印 "received job 1"
| 发送任务2 → jobs                    |→ 接收任务2，打印 "received job 2"
| 发送任务3 → jobs                    |→ 接收任务3，打印 "received job 3"
| 关闭 jobs                          |→ 检测到关闭，打印 "received all jobs"，发送 done → 主协程
| 等待 done                          |← 接收 done 信号，继续执行
| 检测 jobs 状态 → ok=false            |
| 输出 "received more jobs: false"    |

*/
