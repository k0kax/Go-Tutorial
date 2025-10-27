package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!")

	os.Exit(3)
}

/**
os.Exit 是 Go 语言中唯一能​​立即终止程序​​的入口函数，其核心行为包括：

​​跳过所有延迟调用​​：直接绕过 defer 注册的函数（如网页1指出"调用os.Exit后不会执行后续代码，包括defer语句"）
​​中断执行流​​：终止时不会将控制权交还给调用栈（如网页5对比指出"os.Exit会立即退出程序，且不执行任何延迟函数"）

**/
