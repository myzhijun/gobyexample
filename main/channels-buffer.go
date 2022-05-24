package main

import "fmt"

func main() {
	msg := make(chan string, 2)
	msg <- "dujuan"
	msg <- "zhijun"
	message := <-msg
	fmt.Println(message)
	message = <-msg
	fmt.Println(message)
}
