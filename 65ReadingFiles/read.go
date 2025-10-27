package main

//文件读取
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//读取整个文件
	dat, err := os.ReadFile("./test.txt") //同目录下
	check(err)
	fmt.Print(string(dat))
	//分块读取
	f, err := os.Open("./test.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1) //n1实际读取字节数
	check(err)
	fmt.Printf("\n%d bytes:%s\n", n1, string(b1[:n1]))

	//文件指针定位
	o2, err := f.Seek(6, io.SeekStart) //从起始位置偏移6字节，o2 表示新的偏移量（固定为 6）
	check(err)
	b2 := make([]byte, 10)
	n2, err := f.Read(b2) //n2：第二次读取的字节数​
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(2, io.SeekCurrent) //当前位置
	check(err)

	_, err = f.Seek(-4, io.SeekEnd) //文件末尾
	check(err)

	o3, err := f.Seek(3, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2) //强制最少读取2个字节
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	//缓冲读取与预读
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) //预读 5 字节但不移动文件指针
	check(err)
	fmt.Printf("5 bytes:%s\n", string(b4))

	f.Close()
}
