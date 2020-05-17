package main

import "fmt"

/**
  可缓冲的通道，可设置缓存数量
*/
func main() {
	message := make(chan string, 2)
	message <- "bufffer"
	message <- "channel"
	go func() { message <- "dfdf" }()
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
