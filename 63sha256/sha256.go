package main

//sha256 hash
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s)) //得到byte的slice

	bs := h.Sum(nil) //直接获取哈希结果​，创建新切片

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
