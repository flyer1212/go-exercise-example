package main

import (
	"fmt"
	"time"
)

/**
   定时器：未来某一段时间
   可以在失效之前，取消定时
 */
func main() {



	timer1 := time.NewTimer(time.Second*2)
	<-timer1.C // 会阻塞在这里
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println("wewe")
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("stop")
	}
	go func() {
		fmt.Println("wwe")
	}()
}