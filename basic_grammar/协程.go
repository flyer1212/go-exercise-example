package main

import "fmt"

/**

 */

func f(from string) {
	for i := 0; i< 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	 // 必须下面的，才能看到协程的执行
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}