package main

import (
	"fmt"
	"time"
)

/**

 */
func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	p(now.Hour())
	p(now.Year())
	p(now.Second())
	p(now.Unix())//秒
	p(now.UnixNano()) //纳秒
	p(now.UnixNano()/1000000) // 毫秒
}