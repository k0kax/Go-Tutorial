package main

//在 Go 语言中，获取自 ​​Unix 纪元​​（1970 年 1 月 1 日 UTC 时间）以来的时间戳（秒、毫秒、微秒、纳秒）是一项常见需求。
import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
