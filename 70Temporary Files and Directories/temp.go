package main

//临时文件
import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	/**
		""：表示使用系统默认临时目录（如Linux的/tmp）
		"sample"：文件名模式，最终生成类似sample123456的随机文件名
	**/
	check(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name()) //删除临时文件

	_, err = f.Write([]byte{1, 2, 3, 4}) //写入二进制
	check(err)

	dname, err := os.MkdirTemp("", "sampledir") //临时文件夹
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")        //拼接
	err = os.WriteFile(fname, []byte{1, 2}, 0666) //0666表示所有用户可读写
	check(err)
}
