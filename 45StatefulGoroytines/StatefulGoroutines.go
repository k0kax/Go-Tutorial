package main

//有状态的Goroutine/具备状态管理的协程
/*
该程序通过​​通道通信​​实现共享状态的并发安全访问，包含以下核心组件：

1.​​状态管理协程​​：唯一拥有共享map[int]int的协程，通过select处理读写请求
​​2.读写操作结构体​​：readOp/writeOp封装请求与响应通道
​​3.并发读写协程​​：100个读协程和10个写协程，通过通道发送请求
​​4.原子计数器​​：atomic.AddUint64统计操作次数，避免竞争

*/
import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64 //原子计数器
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	//状态管理
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	//读通道
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read                 //发送请求
				<-read.resp                   //等待响应
				atomic.AddUint64(&readOps, 1) // 原子递增读操作计数
				time.Sleep(time.Millisecond)
			}

		}()
	}
	//写通道
	for W := 0; W < 100; W++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(1000),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps) //读原子锁
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}
