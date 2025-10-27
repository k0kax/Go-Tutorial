package main

//工作池
import (
	"fmt"
	"time"
)

// jobs 接受数据，results 发送数据
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//3 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5 job
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//（从results 收集结果）/​​协程同步
	/*
		循环次数等于任务总数 numJobs，确保主协程在所有任务完成后才继续执行后续逻辑。
		若未接收完所有结果，主协程会永久阻塞，防止程序提前退出导致协程未完成
	*/
	for a := 1; a <= numJobs; a++ {
		<-results // 阻塞等待触发
		//唤醒条件​​：当 results通道接收到数据时，调度器将主协程重新标记为可运行（Runnable）
	}
}
