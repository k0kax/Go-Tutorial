package main

import "os"

func main() {
	panic("a problem") //立即中断程序

	_, err := os.Create("/tmp/file") //不会执行
	if err != nil {
		panic(err)
	}
}
