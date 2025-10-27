package main

//运行cat .\test.txt | go run .\line.go
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 创建缓冲扫描器
	scanner := bufio.NewScanner(os.Stdin) // 绑定到标准输入 [2,5](@ref)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)

		if ucl == "QUIT" {
			break
		}
	}

	//scanner.Err() 用于检查 Scanner 在读取输入流（如文件、网络连接等）过程中是否发生了错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		/**
		错误输出：fmt.Fprintln(os.Stderr, ...)
		​​输出目标​​：将错误信息写入 os.Stderr（标准错误流），而非 os.Stdout（标准输出流）。这种设计可将正常输出与错误日志分离，便于后续日志分析或重定向。
		​格式化输出​​：fmt.Fprintln 会在输出内容后自动添加换行符，并支持多参数拼接，例如 "error:" 和 err 变量
		**/
		os.Exit(1) //调用 os.Exit(1) 会立即终止当前进程，跳过所有 defer 延迟函数，直接向操作系统返回状态码 1（非零状态码通常表示异常退出）
	}
}
