package main

import (
	"fmt"
	"time"
)

/**

 */
func workers(id int, jobs <-chan int, results chan<- int) {
	// 没有传递任务之前，这里是阻塞的
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go workers(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		fmt.Println(<-results)
	}
}
