package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}
func add(a ...int) int {
	var sum int = 0
	for _, n := range a {
		sum += n
	}
	return sum
}
func main() {
	fmt.Println("ans:", plus(1, 2))
	fmt.Println("ans:", add())
}
