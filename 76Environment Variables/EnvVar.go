package main

//环境变量
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() { //所有的环境变量KEY=value
		pair := strings.SplitN(e, "=", 2) //将字符串 e 按分隔符 = 分割，最多分割成 ​​2 个子字符串​​（键和值）
		fmt.Println(pair[0])
	}
}
