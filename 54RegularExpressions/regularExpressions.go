package main

//正则表达式
import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	//匹配p开头ch结尾的一个或多个字母
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("MatchString:", match)
	//预编译正则表达式
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("MatchString:", r.MatchString("peach"))

	//查找首个匹配
	fmt.Println("FindString:", r.FindString("peach punch"))

	//匹配索引
	fmt.Println("FindString idx:", r.FindStringIndex("peach punch"))

	//提取子匹配
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//提取子匹配的索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//所有匹配 -1返回所有匹配项
	fmt.Println("FindAllString1:", r.FindAllString("peach punch pinch", -1))

	//所有匹配的索引
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	//返回所有匹配的两个
	fmt.Println("FindAllString2:", r.FindAllString("peach punch pnch", 2))

	fmt.Println(r.Match([]byte("peach"))) //字节切片

	r = regexp.MustCompile("p([a-z]+)ch") //用panic代替error
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	//动态替换函数
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
