package main

import "fmt"

/**

 */
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg:= <- messages:
		fmt.Println("received" , msg)
	default:
		fmt.Println("no msg")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg:= <-messages:
		fmt.Println("received msg",msg)
	case sig:= <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}