package main

import "fmt"

func main() {
	//2个值 通道带缓冲
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
