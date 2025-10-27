package main

//随机数
import (
	"fmt"
	"math/rand/v2"
)

func main() {

	fmt.Print(rand.IntN(100), ",") //0-100
	fmt.Print(rand.IntN(100))
	fmt.Println()

	fmt.Print((rand.Float64())) //0.0-1.0

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64())*5 + 5)
	fmt.Println()

	/**
	基于 PCG 算法的随机数生成。42 和 1024 是两个 uint64 类型的种子值。
	PCG 算法要求提供两个独立的种子，这使得生成的随机序列具有更高的随机性。
	**/
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
