package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // 注册信号类型[1,6](@ref)
	//SIGINT（Ctrl+C）和 SIGTERM（终止命令）是最常用的进程终止信号
	done := make(chan bool, 1) // 控制主程序退出的同步通道

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
