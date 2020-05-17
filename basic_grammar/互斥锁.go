package main

import (
	"fmt"
	"sync"
	"time"
)

/**

 */
func main() {

	var mutex = &sync.Mutex{}
	var ops int64 = 0

	for r := 0; r < 10000; r++ {
		go func() {
			mutex.Lock()
			ops++
			mutex.Unlock()
		}()
	}
	// 一定要sleep, 否则可能没有执行完

	// 加互斥锁刚好10000次
	// 不加互斥锁 9596多次
	time.Sleep(time.Second)
	fmt.Println(ops)

}
