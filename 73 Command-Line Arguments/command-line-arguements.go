package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args        //提供原始命令行参数的访问，第一个元素是程序的路径
	argsWithoutProg := os.Args[1:] //这个才是实际的参数

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
