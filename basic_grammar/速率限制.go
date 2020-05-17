package main

import (
	"fmt"
	"time"
)

/**
  每隔固定时间执行一次
*/
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter // 每200ms 限制一次
		fmt.Println("request", req, time.Now())
	}

	// 先放入3个进入通道
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 每隔两百毫秒，继续往里面添加
	go func() {
		for t := range time.Tick(time.Millisecond * 1000) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter // 速率限制
		fmt.Println("request ", req, time.Now())
	}
}
