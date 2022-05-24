package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	nexInt := intSeq()
	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())
	newInts := intSeq()
	fmt.Println(newInts())
}
