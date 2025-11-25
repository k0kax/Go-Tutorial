package main

import (
	"fmt"
	"slices"
)

// 基础
func base() {
	var s []string
	fmt.Println("unint:", s, s == nil, len(s) == 0)

	s = make([]string, 3, 5) //切片长度3，底层数组容量为5
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}

// 扩容
func extend() {
	// 创建一个长度为 2，容量为 4 的切片
	// 底层数组：[0, 0, _, _]
	// 切片视图：[0, 0]
	s := make([]int, 2, 4)
	fmt.Printf("初始状态: len=%d, cap=%d, 切片=%v\n", len(s), cap(s), s)

	// 1. append 一个元素
	// len=2 < cap=4，不需要扩容
	s = append(s, 10)
	fmt.Printf("append 10 后: len=%d, cap=%d, 切片=%v\n", len(s), cap(s), s)

	// 2. 再 append 一个元素
	// len=3 < cap=4，仍然不需要扩容
	s = append(s, 20)
	fmt.Printf("append 20 后: len=%d, cap=%d, 切片=%v\n", len(s), cap(s), s)

	// 3. 再 append 一个元素
	// len=4 == cap=4，触发扩容！
	// 新容量通常是原容量的 2 倍，即 8
	// 创建新数组，复制元素，然后 append 30
	s = append(s, 30)
	fmt.Printf("append 30 后: len=%d, cap=%d, 切片=%v\n", len(s), cap(s), s)
}

// 覆盖操作
func override() {
	// 1. 切片 s 最初指向数组 arr1
	arr1 := [2]int{10, 20}
	s := make([]int, 5)
	fmt.Println("s未被覆盖时:", s, ",len=", len(s), ",cap=", cap(s))
	s = arr1[:]
	fmt.Println("s被arr1覆盖后指向:", s, ",len=", len(s), ",cap=", cap(s)) // 输出: [10 20]

	// 2. 关键：将 s 重新赋值，让它指向新的数组 arr2
	arr2 := [3]int{100, 200, 300}
	s = arr2[:]                                                      // 这就是“覆盖”操作
	fmt.Println("s被arr2覆盖后指向:", s, ",len=", len(s), ",cap=", cap(s)) // 输出: [100 200]
}
func main() {
	//base()
	extend()
	//override()
}
