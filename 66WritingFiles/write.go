package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	//WriteFile覆写内容
	err := os.WriteFile("./test.txt", d1, 0644) // 0644 表示属主可读写
	check(err)

	f, err := os.Create("./test2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	n3, err := f.WriteString("writes\n") //追加写入
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() //强制将文件系统缓存中的数据（包括元数据）刷新到物理存储设备（如磁盘/SSD）

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() //清空缓存
}
