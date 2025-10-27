package main

//"//go:embed" 是一个编译器指令，允许程序在构建时将任意文件及文件夹嵌入到 Go 二进制文件中
import "embed"

//把这个文件的内容嵌入进fileString里了
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS //实现了一个简易的​​虚拟文件系统

func main() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
