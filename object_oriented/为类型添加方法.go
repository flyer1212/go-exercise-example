package main

import "fmt"

/**
  给任意类型添加方法
  go语言都是基于值传递的, 如果需要修改原值,只能用指针


*/
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) add(b Integer){
	*a += b
}
func (a Integer) add2(b Integer){
	a += b
}

type People struct {
	name string
	age  int
}

func (people People) getAge() int {
	return people.age
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}
	people := People{"we", 21}
	fmt.Println(people.getAge())

	 

	var a2 Integer =1
	a2.add(2)
	fmt.Println("a=", a2)
	a2.add2(2)
	fmt.Println("a=", a2)
}
