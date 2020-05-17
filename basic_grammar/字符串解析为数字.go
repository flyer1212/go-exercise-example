package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
λ 字符串解析为数字 a b c d
1.234
[字符串解析为数字 a b c d]
[a b c d]
c

 */
func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	argsWithProg := os.Args
	fmt.Println(argsWithProg)
	arg := os.Args[1:]
	fmt.Println(arg)
	args := os.Args[3]
	fmt.Println(args)
}
