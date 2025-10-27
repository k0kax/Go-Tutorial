package main

//限流
import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//固定速率限流
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter //阻塞200ms间隔到达
		fmt.Println("request1", req, time.Now())
	}

	//突发限流器
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ { //预补充3个令牌
		burstyLimiter <- time.Now() //填充令牌
	}

	go func() { //定时补充令牌
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	//突发请求 5个
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter //取出令牌
		fmt.Println("request2", req, time.Now())
	}
}
