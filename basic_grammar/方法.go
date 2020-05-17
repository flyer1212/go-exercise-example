package main

import "fmt"

/**
  Go 自动处理方法调用时1的值和指针之间的变化。
  1. 可以用指针来调用方法，避免方法调用时产生一个拷贝

  为结构体类型定义方法
 */

type rect struct {
	width, height int
}

func (r *rect) area() int{
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height;
}
func main() {
	r := rect{width:10, height:5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ",rp.area())
	fmt.Println("perim: ", rp.perim())
}