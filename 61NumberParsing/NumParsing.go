package main

//数字解析
import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64) //string->float64
	fmt.Println(f)

	//当 ParseInt 的 base 参数设置为 0 时，函数会根据字符串的 ​​前缀​​ 自动推断进制
	//bitSize 参数指定目标整数的 ​​位数​​
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	//Atoi（ASCII to Integer）主要用于将字符串转换为十进制（base-10）整数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
