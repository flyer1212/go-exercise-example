package main

/**
   只进不出的通道
 */
func ping(pings chan <- string, msg string){
	pings <- msg
}
/**
  可进可出的通道
 */
func pong(pings <-chan string, pongs chan <-string){
	msg := <- pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pass msg")
	pong(pings, pongs)
	println(<-pongs)
}