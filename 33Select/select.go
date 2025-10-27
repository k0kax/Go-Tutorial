package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1, "i=", i)
		case msg2 := <-c2:
			fmt.Println("received:", msg2, "i=", i)
		}
	}

	/*
	 时间轴		 Goroutine 1	 		Goroutine 2	 			主程序 select 行为
	  0秒	 	 启动，开始休眠1秒		启动，开始休眠2秒		 进入第一次 select，阻塞等待
	  1秒	 	 发送 "one" 到 c1		继续休眠				触发 c1 的 case，输出结果
	  1秒~2秒	       结束				继续休眠				进入第二次 select，阻塞等待
	  2秒	-						   发送 "two" 到 c2		   触发 c2 的 case，输出结果
	*/

}
