package main

/*
iterators 迭代器的本质​​是一个​​接收 yield 回调函数的函数
可分为三类：
 1.无参数迭代器（较少使用）
	func(yield func() bool)
 2.单值迭代器（值类型）
	func(yield func(V) bool)
 3.键值对迭代器（键值类型）
	func(yield func(K, V) bool)

推迭代器
核心机制：通过调用 yield 函数逐个推送元素。
若 yield 返回 false，迭代器会立即终止遍历。


具体见相关文章
*/
import (
	"fmt"
	"iter"
	"slices"
)

// 同generics.go 泛型列表
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

/*
迭代器实现核心-遍历
单值迭代器

迭代器协议​​：符合 iter.Seq[T] 类型（函数签名 func(yield func(T) bool)）
​
​延迟执行特性​​：迭代逻辑在闭包中，实际遍历由 for...range 触发

​​控制流​​：调用 yield(e.val) 时：

	若返回 true：继续迭代下一个节点
	若返回 false：立即终止迭代（如 break 语句触发）
*/
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) { //将节点值传递给yield
				return //当yield返回false时终止迭代
			}
		}
	}
}

// 斐波那契数列-单值迭代
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1 //状态量

		for { // 无限循环生成序列
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	//将迭代器结果收集到切片（类似 Python 的 list(iterator)）
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
