package main

//泛型（Generics），又叫类型参数（Type Parameters）
import "fmt"

/*
​​泛型参数​​：
S ~[]E：约束 S 必须是元素类型为 E 的切片（支持自定义切片类型，如 type MySlice []int）。
E comparable：约束 E 必须是可比较类型（支持 == 和 != 操作符），包括：布尔值、数字、字符串、指针、通道等。
​​参数与返回值​​：
s S：待搜索的切片。
v E：目标元素。
返回值：首个匹配元素的索引，未找到返回 -1。
*/
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

// 数据结构 链表
// T any 表示链表可存储任意类型元素
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 入链
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// 遍历链表生成切片
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
