package main

import "fmt"

type rect struct {
	width, height int
}

//指针接受
func (r *rect) area() int {
	r.width = 20
	return r.width * r.height
}

//值接受
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("perim: ", r.perim())
	fmt.Println("area: ", r.area())

	rp := &r
	fmt.Println("perim:", rp.perim())
	fmt.Println("area: ", rp.area())

}
