package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages: //message 无数据可读
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	/*
		default 分支用于在通道操作无法立即完成时提供非阻塞行为。
		无缓冲通道的发送和接收操作需要同步的发送方和接收方，否则会阻塞

	*/

	msg := "hi"
	select {
	case messages <- msg: //发送字符，无缓冲通道，无接收者
		//“无接收者”指​​向通道（Channel）发送数据时，没有其他协程（Goroutine）在监听或接收该通道的数据​​。
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	/*
		发送操作需要接收方就绪才能成功，否则会触发 default 分支。
		若将 messages 改为缓冲通道（如 make(chan string, 1)），发送会成功，输出 sent message hi。
	*/

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
