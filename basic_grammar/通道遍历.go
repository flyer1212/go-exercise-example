package main

import "fmt"

/**
   通道关闭后，通道里面的元素也是可以得到的

   通道不关闭，range会阻塞
 */
func main() {

	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}