package main

//启动进程
import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	//临时改下cmd的gbk编码格式为utf-8
	codes := exec.Command("cmd", "/c", "chcp 65001")

	dateOut, e := codes.Output()
	if e != nil {
		panic(e)
	}
	//命令执行窗口，需输入调用程序的完整参数信息
	dateCmd := exec.Command("cmd", "/c", "echo %DATE% %TIME%")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(" > date")
	fmt.Println(string(dateOut))
	//cmd的"/c"表示执行命令后关闭窗口
	_, err = exec.Command("cmd", "/C", "date", "/invalidflag").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error: //命令执行失败（如路径错误）
			fmt.Println("failed executing:", err)
		case *exec.ExitError: //命令执行但返回非零状态码（如参数错误）
			fmt.Println("command exit rc = ", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("findstr", "hello")

	grepIn, _ := grepCmd.StdinPipe()   //管道输入
	grepOut, _ := grepCmd.StdoutPipe() //管道输出
	grepCmd.Start()
	grepIn.Write([]byte("hello gr123423ep\ngoobye findstr")) //写入管道
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut) //读取输出
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// /A：显示所有文件（含隐藏文件）
	// /W：宽列表格式（类似 Unix 的 -h 可读格式）
	lsCmd := exec.Command("cmd", "/C", "dir /A /W")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> dir /A /W")
	fmt.Println(string(lsOut))
}
