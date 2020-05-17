package main

import "fmt"

/**
   通过指针， 通过引用传递值 和 数据结构

 */

func zeroval(ival int){
	ival = 0
}

func zeroptr(iptr *int){
	*iptr =0
}
func main() {
	i := 1
	fmt.Println("init", i)
	zeroval(i)
	fmt.Println("zeroval", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer:", &i)
}