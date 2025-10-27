package main

//结构体/接口 嵌入继承
import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base //匿名嵌入
	str  string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v,str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe()) //调用提升方法

	type describer interface {
		describe() string
	}
	//通过嵌入带有方法的结构体，可以使其他结构体自动实现特定接口
	var d describer = co //通过接口调用 co有describe（）
	fmt.Println("describer: ", d.describe())
}
