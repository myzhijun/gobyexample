package main

import "fmt"

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Print(i, "-")
	}
}
func main() {
	//f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
