package main

//
/*
	关于通道的方向问题
	chan<- string 数据 发送 到chan
	<-chan string 从chan 接受 数据

*/
import "fmt"

//单向发送通道：只能发送数据到 pings
func ping(pings chan<- string, msg string) {
	pings <- msg

}

//单向接收通道：pings 接收数据 ，pongs发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
