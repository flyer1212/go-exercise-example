package main

import "fmt"

/**

 */

type person struct {
	name string
	age int
}

func main() {
	s := person{name:"sean", age: 50}
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)
}