package main

import (
	"os"
	"os/exec"
)

func main() {

	// 设置控制台编码为 UTF-8（解决中文乱码）
	exec.Command("cmd", "/C", "chcp 65001").Run()

	// 查找 PowerShell 路径
	binary, lookErr := exec.LookPath("powershell.exe")
	if lookErr != nil {
		panic(lookErr)
	}

	// 定义 PowerShell 命令参数
	args := []string{
		"-NoProfile", // 不加载用户配置文件
		"-Command",   // 执行命令模式
		"Get-ChildItem -Force | Format-Table Name, Mode, Length, LastWriteTime", // 等效于 ls -alh
	}

	env := os.Environ()

	// 创建命令对象
	cmd := exec.Command(binary, args...)
	cmd.Env = env
	cmd.Stdout = os.Stdout // 输出到标准终端
	cmd.Stderr = os.Stderr // 错误输出到终端

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	// 	execErr := syscall.Exec(binary, args, env)
	// 	if execErr != nil {
	// 		panic(execErr)
	// 	}
	/**
	syscall.Exec 是 Go 语言中直接调用操作系统底层 ​
	​进程替换系统调用​​ 的函数。
	其行为与 Unix/Linux 的 execve() 类似，
	会 ​​完全替换当前进程的代码段、数据段和堆栈​​，
	加载并执行新的程序映像。执行成功后，
	原进程代码不再继续运行。
	**/
}
