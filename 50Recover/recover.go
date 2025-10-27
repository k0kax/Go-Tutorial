package main

//recover 阻止panic终止运行
import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		//捕获panic
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic() //触发panic,立即中断函数运行，执行defer函数

	fmt.Println("After mayPanic()")
}
