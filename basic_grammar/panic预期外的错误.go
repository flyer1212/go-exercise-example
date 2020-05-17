package main

import "os"

/**

 */
func main() {
	//可以单独运行，终止程序
	// panic("a problem")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

}