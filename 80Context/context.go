package main

/**
	Context 能够跨越 API 边界和 goroutine，
	传递截止时间（deadlines）、取消信号（cancellation signals）
	以及其他请求范围（request-scoped）的值。
**/
import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context() // 获取请求上下文
	fmt.Println("server:hello handler started")
	defer fmt.Println("server:hello handler ended")

	select {
	case <-time.After(10 * time.Second): //等待十秒
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done(): //中断执行其他操作
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
