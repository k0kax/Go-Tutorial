package main

//命令行参数
import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a atring")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string //也可以声明使用已经存在的变量，需传入变量的指针
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse() //解析

	fmt.Println("word:", *wordPtr) //使用指针
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args()) //
}
