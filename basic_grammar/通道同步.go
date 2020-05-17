package main

import (
	"fmt"
	"time"
)

/**

 */

func worker(done chan bool){
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool,1)
	go worker(done)


	// 顺序执行的
	if <-done {
		println("2323")
	}
	println("dsds")
}