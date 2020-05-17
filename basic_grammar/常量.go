package main

import (
	"fmt"
	"math"
)
/*
   任意var可以出现的地方，都可以写常量
 */
const s string = "constant"

func main()  {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20/n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}