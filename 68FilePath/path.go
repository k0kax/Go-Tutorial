package main

//文件路径构建
import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	//路径拼接
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	/**
	冗余斜杠处理​​：dir1// 被规范化为 dir1/
	​​路径回溯解析​​：dir1/../dir1 通过 Clean 优化为等效最短路径 dir1
	**/
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	//路径分解
	fmt.Println("Dir(p):", filepath.Dir(p))   //目录部分（去掉最后的元素）
	fmt.Println("Base(p):", filepath.Base(p)) //最终元素（文件名）

	fmt.Println(filepath.IsAbs("Dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename) //获取后缀
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext)) //去除后缀获取文件名

	//相对路径计算
	/**
		基础路径​​：a/b
		​​目标路径​​：a/b/t/file
		​​共同祖先​​：a/b 是目标路径的直接父目录
		​​相对路径生成​​：
		目标路径比基础路径多出 t/file 部分
		无需回溯上级目录，直接拼接子路径
		输出结果为 t/file
	**/
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	/**
		逻辑分解​​：

		​​基础路径​​：a/b
		​​目标路径​​：a/c/t/file
		​​共同祖先​​：a（需从 a/b 上溯到 a）
		​​路径回溯与拼接​​：
		从 a/b 到 a：需上溯一层（..）
		从 a 到 a/c/t/file：需拼接 c/t/file
		合并路径为 ../../c/t/file
		第一个 .. 从 a/b 到 a
		第二个 .. 用于抵消基础路径和目标路径层数差异（此处因目标路径层级更深，可能需要调整）
	**/
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
